package service

import (
	"context"
	"net"
	"os"
	"reflect"
	"runtime"
	"testing"
	"time"

	faultinjectorpb "github.com/Netflix-skunkworks/grpc_fault/faultinjector"
	helloworld "github.com/Netflix-skunkworks/grpc_fault/internal/testing/helloworld/protos"
	"github.com/golang/protobuf/proto" // nolint: staticcheck
	"github.com/rs/zerolog"
	"golang.org/x/net/nettest"
	"google.golang.org/grpc"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

type testContext struct {
	interceptor            *faultInjectorInterceptor
	faultinjectionserver   *grpc.Server
	faultinjectionlistener net.Listener
	// connection *to* the fault injection server
	faultinjectorconn   *grpc.ClientConn
	faultinjectorclient faultinjectorpb.FaultInjectorClient

	helloworldserver   *grpc.Server
	helloworldlistener net.Listener
}

type testFunc func(ctx context.Context, t *testing.T, tc testContext)

func TestHarness(t *testing.T) {
	tests := []testFunc{
		testBasic,
		testFaultInjectionInterceptor,
		testRegistration,
	}
	for idx := range tests {
		testFunc := tests[idx]
		name := runtime.FuncForPC(reflect.ValueOf(testFunc).Pointer()).Name()
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			logger := zerolog.New(os.Stderr).Level(zerolog.DebugLevel)
			ctx = logger.WithContext(ctx)

			faultinjectionlistener, err := nettest.NewLocalListener("tcp")
			assert.NilError(t, err)
			faultinjectionserver := grpc.NewServer()
			interceptor := NewInterceptor(ctx, faultinjectionserver)
			go func() {
				_ = faultinjectionserver.Serve(faultinjectionlistener)
			}()

			go func() {
				<-ctx.Done()
				faultinjectionserver.GracefulStop()
				_ = faultinjectionlistener.Close()
			}()

			helloworldlistener, err := nettest.NewLocalListener("tcp")
			assert.NilError(t, err)
			helloworldserver := grpc.NewServer()
			helloworld.RegisterGreeterServer(helloworldserver, &testGreeterService{})
			go func() {
				_ = helloworldserver.Serve(helloworldlistener)
			}()

			go func() {
				<-ctx.Done()
				helloworldserver.GracefulStop()
				helloworldlistener.Close()
			}()

			faultinjectorconn, err := grpc.DialContext(ctx, faultinjectionlistener.Addr().String(), grpc.WithInsecure())
			assert.NilError(t, err)
			faultinjectorclient := faultinjectorpb.NewFaultInjectorClient(faultinjectorconn)

			testFunc(ctx, t, testContext{
				faultinjectionlistener: faultinjectionlistener,
				faultinjectionserver:   faultinjectionserver,
				faultinjectorconn:      faultinjectorconn,
				faultinjectorclient:    faultinjectorclient,
				interceptor:            interceptor.(*faultInjectorInterceptor),

				helloworldlistener: helloworldlistener,
				helloworldserver:   helloworldserver,
			})
			cancel()
		})
	}
}

func testBasic(ctx context.Context, t *testing.T, tc testContext) {
	enumerateServicesResponse, err := tc.faultinjectorclient.EnumerateServices(ctx, &faultinjectorpb.EnumerateServicesRequest{})
	assert.NilError(t, err)
	assert.Assert(t, proto.Equal(enumerateServicesResponse, &faultinjectorpb.EnumerateServicesResponse{}))
	assert.Assert(t, is.Len(enumerateServicesResponse.Services, 0))

	tc.interceptor.RegisterService(helloworld.Greeter_ServiceDesc)

	enumerateServicesResponse, err = tc.faultinjectorclient.EnumerateServices(ctx, &faultinjectorpb.EnumerateServicesRequest{})
	assert.NilError(t, err)
	assert.Assert(t, is.Len(enumerateServicesResponse.Services, 1))
	t.Log(enumerateServicesResponse.String())
}

func testFaultInjectionInterceptor(ctx context.Context, t *testing.T, tc testContext) {
	helloworldconn, err := grpc.DialContext(ctx, tc.helloworldlistener.Addr().String(), grpc.WithInsecure(), grpc.WithUnaryInterceptor(tc.interceptor.UnaryClientInterceptor))
	assert.NilError(t, err)
	helloworldclient := helloworld.NewGreeterClient(helloworldconn)
	resp, err := helloworldclient.SayHello(ctx, &helloworld.HelloRequest{
		Name: "Sargun",
	})
	assert.NilError(t, err)
	assert.Assert(t, is.Equal("Sargun", resp.Message))
}

func testRegistration(ctx context.Context, t *testing.T, tc testContext) {
	tc.interceptor.RegisterService(helloworld.Greeter_ServiceDesc)

	_, err := tc.faultinjectorclient.RegisterFault(ctx, &faultinjectorpb.RegisterFaultRequest{
		Service:    "helloworld.Greeter",
		Method:     "SayHello",
		Expression: `req.Name == "Sargun" ? PermissionDenied("Sargun Not Allowed") : OK()`,
	})
	assert.NilError(t, err)

	helloworldconn, err := grpc.DialContext(ctx, tc.helloworldlistener.Addr().String(), grpc.WithInsecure(), grpc.WithUnaryInterceptor(tc.interceptor.UnaryClientInterceptor))
	assert.NilError(t, err)

	helloworldclient := helloworld.NewGreeterClient(helloworldconn)
	_, err = helloworldclient.SayHello(ctx, &helloworld.HelloRequest{
		Name: "Sargun",
	})
	assert.ErrorContains(t, err, "PermissionDenied")

	_, err = helloworldclient.SayHello(ctx, &helloworld.HelloRequest{
		Name: "NotSargun",
	})
	assert.NilError(t, err)

	services, err := tc.faultinjectorclient.EnumerateServices(ctx, &faultinjectorpb.EnumerateServicesRequest{})
	assert.NilError(t, err)
	assert.Assert(t, is.Len(services.Services, 1))
	assert.Assert(t, is.Len(services.Services[0].Methods, 1))
	assert.Assert(t, services.Services[0].Methods[0].Expression != "")

	_, err = tc.faultinjectorclient.RemoveFault(ctx, &faultinjectorpb.RemoveFaultRequest{
		Service: "helloworld.Greeter",
		Method:  "SayHello",
	})
	assert.NilError(t, err)

	services, err = tc.faultinjectorclient.EnumerateServices(ctx, &faultinjectorpb.EnumerateServicesRequest{})
	assert.NilError(t, err)
	assert.Assert(t, services.Services[0].Methods[0].Expression == "")

	_, err = helloworldclient.SayHello(ctx, &helloworld.HelloRequest{
		Name: "Sargun",
	})
	assert.NilError(t, err)
}

type testGreeterService struct {
	helloworld.UnimplementedGreeterServer
}

func (t *testGreeterService) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: request.Name,
	}, nil
}
