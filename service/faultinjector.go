package service

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	faultinjectorpb "github.com/Netflix-skunkworks/grpc_fault/faultinjector"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var (
	defaultInterceptor FaultInjectorInterceptor
)

type FaultInjectorInterceptor interface {
	UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error
	// Only the method is passed to the fault injection inspection
	StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error)
	RegisterService(service grpc.ServiceDesc)
}

// GenerateInterceptor can be used to register an interceptor for a given grpc service. If a service is not registered
// the interceptor methods will error out
func RegisterInterceptor(service grpc.ServiceDesc) {
	if defaultInterceptor == nil {
		return
	}
	defaultInterceptor.RegisterService(service)
}

func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	if defaultInterceptor == nil {
		return invoker(ctx, method, req, reply, cc, opts...)
	}

	return defaultInterceptor.UnaryClientInterceptor(ctx, method, req, reply, cc, invoker, opts...)
}

func StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	if defaultInterceptor == nil {
		return streamer(ctx, desc, cc, method, opts...)
	}

	return defaultInterceptor.StreamClientInterceptor(ctx, desc, cc, method, streamer, opts...)
}

// NewInterceptor can be used to generate a single instance of the interceptor
func NewInterceptor(ctx context.Context, server *grpc.Server) FaultInjectorInterceptor {
	fii := &faultInjectorInterceptor{}
	fis := &faultInjectorServer{
		fii: fii,
	}

	faultinjectorpb.RegisterFaultInjectorServer(server, fis)

	return fii
}

type faultInjectorInterceptor struct {
	// Map of services -> *serviceFaultInjectorInterceptor
	services sync.Map
}

func (f *faultInjectorInterceptor) RegisterService(service grpc.ServiceDesc) {
	f.services.LoadOrStore(service.ServiceName, &serviceFaultInjectorInterceptor{
		service: service,
	})
}

func (f *faultInjectorInterceptor) UnaryClientInterceptor(ctx context.Context, methodAndService string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// Can either be:
	// * /helloworld.Greeter/SayHello
	// Or:
	// * helloworld.Greeter/SayHello
	splitMethod := strings.Split(strings.TrimLeft(methodAndService, "/"), "/")
	serviceName := splitMethod[0]
	methodName := splitMethod[1]
	val, ok := f.services.Load(serviceName)
	if ok {
		sfii := val.(*serviceFaultInjectorInterceptor)
		val, ok = sfii.methods.Load(methodName)
		if ok {
			mfii := val.(*methodFaultInjectorInterceptor)
			val, err := vm.Run(mfii.program, makeEnv(req))
			if err != nil {
				log.Ctx(ctx).WithLevel(zerolog.WarnLevel).Err(err).Msg("Could not run program")
			} else if val != nil {
				err, ok := val.(error)
				if !ok {
					log.Ctx(ctx).WithLevel(zerolog.WarnLevel).Str("type", fmt.Sprintf("%T", err)).Msg("Received unexpected type from program")
				}
				if status.Code(err) != codes.OK {
					return err
				}
			}
		}
	}
	return invoker(ctx, methodAndService, req, reply, cc, opts...)
}

func (f *faultInjectorInterceptor) StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Ctx(ctx).WithLevel(zerolog.InfoLevel).Str("method", method).Msg("Invoked stream client interceptor")
	return streamer(ctx, desc, cc, method, opts...)
}

type serviceFaultInjectorInterceptor struct {
	service grpc.ServiceDesc
	// Map of method name -> methodFaultInjectorInterceptor
	methods sync.Map
}

func (s *serviceFaultInjectorInterceptor) getMethod(name string) *grpc.MethodDesc {
	for idx := range s.service.Methods {
		method := s.service.Methods[idx]
		if method.MethodName == name {
			return &method
		}
	}
	return nil
}

type methodFaultInjectorInterceptor struct {
	expression string
	program    *vm.Program
}

type faultInjectorServer struct {
	faultinjectorpb.UnimplementedFaultInjectorServer
	fii *faultInjectorInterceptor
}

func (f *faultInjectorServer) EnumerateServices(ctx context.Context, request *faultinjectorpb.EnumerateServicesRequest) (*faultinjectorpb.EnumerateServicesResponse, error) {
	resp := faultinjectorpb.EnumerateServicesResponse{}
	f.fii.services.Range(func(key, value interface{}) bool {
		servicename := key.(string)
		fi := value.(*serviceFaultInjectorInterceptor)
		methods := []*faultinjectorpb.Method{}
		for idx := range fi.service.Methods {
			methodDesc := fi.service.Methods[idx]
			method := &faultinjectorpb.Method{
				Name: methodDesc.MethodName,
			}
			val, ok := fi.methods.Load(methodDesc.MethodName)
			if ok {
				mfii := val.(*methodFaultInjectorInterceptor)
				method.Expression = mfii.expression
			}

			methods = append(methods, method)
		}

		svc := &faultinjectorpb.Service{
			Name:    servicename,
			Methods: methods,
		}
		resp.Services = append(resp.Services, svc)

		return true
	})

	return &resp, nil
}
func (f *faultInjectorServer) RegisterFault(ctx context.Context, req *faultinjectorpb.RegisterFaultRequest) (*faultinjectorpb.RegisterFaultResponse, error) {
	val, ok := f.fii.services.Load(req.Service)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "service %s not found", req.Service)
	}

	service := val.(*serviceFaultInjectorInterceptor)
	method := service.getMethod(req.Method)
	if method == nil {
		return nil, status.Errorf(codes.NotFound, "method %s not found", req.Method)
	}

	// ht is the interface of the service.
	ht := reflect.TypeOf(service.service.HandlerType).Elem()
	reflectedMethod, ok := ht.MethodByName(method.MethodName)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "method %s not found on reflection", method.MethodName)
	}

	fakeReq := reflect.New(reflectedMethod.Type.In(1)).Interface()
	env := makeEnv(fakeReq)
	pgrm, err := expr.Compile(req.Expression, expr.Env(env), expr.Optimize(true))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Could not compile the program: %s", err.Error())
	}

	service.methods.Store(method.MethodName, &methodFaultInjectorInterceptor{
		expression: req.Expression,
		program:    pgrm,
	})

	return &faultinjectorpb.RegisterFaultResponse{}, nil
}

func makeCodeFunction(c codes.Code) func(msg string) error {
	return func(msg string) error {
		return status.Error(c, msg)
	}
}

func makeEnv(req interface{}) map[string]interface{} {
	env := map[string]interface{}{
		"req": req,
		"OK": func() error {
			return nil
		},
		"Canceled":           makeCodeFunction(codes.Canceled),
		"Unknown":            makeCodeFunction(codes.Unknown),
		"InvalidArgument":    makeCodeFunction(codes.InvalidArgument),
		"DeadlineExceeded":   makeCodeFunction(codes.DeadlineExceeded),
		"NotFound":           makeCodeFunction(codes.NotFound),
		"AlreadyExists":      makeCodeFunction(codes.AlreadyExists),
		"PermissionDenied":   makeCodeFunction(codes.PermissionDenied),
		"ResourceExhausted":  makeCodeFunction(codes.ResourceExhausted),
		"FailedPrecondition": makeCodeFunction(codes.FailedPrecondition),
		"Aborted":            makeCodeFunction(codes.Aborted),
		"OutOfRange":         makeCodeFunction(codes.OutOfRange),
		"Unimplemented":      makeCodeFunction(codes.Unimplemented),
		"Internal":           makeCodeFunction(codes.Internal),
		"Unavailable":        makeCodeFunction(codes.Unavailable),
		"DataLoss":           makeCodeFunction(codes.DataLoss),
		"Unauthenticated":    makeCodeFunction(codes.Unauthenticated),
	}

	return env
}

func (f *faultInjectorServer) RemoveFault(ctx context.Context, req *faultinjectorpb.RemoveFaultRequest) (*faultinjectorpb.RemoveFaultResponse, error) {
	val, ok := f.fii.services.Load(req.Service)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "service %s not found", req.Service)
	}
	service := val.(*serviceFaultInjectorInterceptor)

	service.methods.Delete(req.Method)
	return &faultinjectorpb.RemoveFaultResponse{}, nil
}
