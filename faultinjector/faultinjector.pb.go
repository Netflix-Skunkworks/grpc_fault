// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: faultinjector.proto

package faultinjector

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EnumerateServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnumerateServicesRequest) Reset() {
	*x = EnumerateServicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faultinjector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumerateServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumerateServicesRequest) ProtoMessage() {}

func (x *EnumerateServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_faultinjector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumerateServicesRequest.ProtoReflect.Descriptor instead.
func (*EnumerateServicesRequest) Descriptor() ([]byte, []int) {
	return file_faultinjector_proto_rawDescGZIP(), []int{0}
}

type EnumerateServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *EnumerateServicesResponse) Reset() {
	*x = EnumerateServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faultinjector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumerateServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumerateServicesResponse) ProtoMessage() {}

func (x *EnumerateServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_faultinjector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumerateServicesResponse.ProtoReflect.Descriptor instead.
func (*EnumerateServicesResponse) Descriptor() ([]byte, []int) {
	return file_faultinjector_proto_rawDescGZIP(), []int{1}
}

func (x *EnumerateServicesResponse) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Methods []*Method `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faultinjector_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_faultinjector_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_faultinjector_proto_rawDescGZIP(), []int{2}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetMethods() []*Method {
	if x != nil {
		return x.Methods
	}
	return nil
}

type Method struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Expression string `protobuf:"bytes,2,opt,name=expression,proto3" json:"expression,omitempty"`
}

func (x *Method) Reset() {
	*x = Method{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faultinjector_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Method) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Method) ProtoMessage() {}

func (x *Method) ProtoReflect() protoreflect.Message {
	mi := &file_faultinjector_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Method.ProtoReflect.Descriptor instead.
func (*Method) Descriptor() ([]byte, []int) {
	return file_faultinjector_proto_rawDescGZIP(), []int{3}
}

func (x *Method) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Method) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

type RegisterFaultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service    string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Method     string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Expression string `protobuf:"bytes,3,opt,name=expression,proto3" json:"expression,omitempty"`
}

func (x *RegisterFaultRequest) Reset() {
	*x = RegisterFaultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faultinjector_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterFaultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterFaultRequest) ProtoMessage() {}

func (x *RegisterFaultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_faultinjector_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterFaultRequest.ProtoReflect.Descriptor instead.
func (*RegisterFaultRequest) Descriptor() ([]byte, []int) {
	return file_faultinjector_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterFaultRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *RegisterFaultRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *RegisterFaultRequest) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

type RegisterFaultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterFaultResponse) Reset() {
	*x = RegisterFaultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faultinjector_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterFaultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterFaultResponse) ProtoMessage() {}

func (x *RegisterFaultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_faultinjector_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterFaultResponse.ProtoReflect.Descriptor instead.
func (*RegisterFaultResponse) Descriptor() ([]byte, []int) {
	return file_faultinjector_proto_rawDescGZIP(), []int{5}
}

type RemoveFaultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Method  string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *RemoveFaultRequest) Reset() {
	*x = RemoveFaultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faultinjector_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFaultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFaultRequest) ProtoMessage() {}

func (x *RemoveFaultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_faultinjector_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFaultRequest.ProtoReflect.Descriptor instead.
func (*RemoveFaultRequest) Descriptor() ([]byte, []int) {
	return file_faultinjector_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveFaultRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *RemoveFaultRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type RemoveFaultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveFaultResponse) Reset() {
	*x = RemoveFaultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faultinjector_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFaultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFaultResponse) ProtoMessage() {}

func (x *RemoveFaultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_faultinjector_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFaultResponse.ProtoReflect.Descriptor instead.
func (*RemoveFaultResponse) Descriptor() ([]byte, []int) {
	return file_faultinjector_proto_rawDescGZIP(), []int{7}
}

type ListenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Method  string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *ListenRequest) Reset() {
	*x = ListenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faultinjector_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenRequest) ProtoMessage() {}

func (x *ListenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_faultinjector_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenRequest.ProtoReflect.Descriptor instead.
func (*ListenRequest) Descriptor() ([]byte, []int) {
	return file_faultinjector_proto_rawDescGZIP(), []int{8}
}

func (x *ListenRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ListenRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type ListenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Reply   string `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ListenResponse) Reset() {
	*x = ListenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faultinjector_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenResponse) ProtoMessage() {}

func (x *ListenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_faultinjector_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenResponse.ProtoReflect.Descriptor instead.
func (*ListenResponse) Descriptor() ([]byte, []int) {
	return file_faultinjector_proto_rawDescGZIP(), []int{9}
}

func (x *ListenResponse) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *ListenResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

func (x *ListenResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_faultinjector_proto protoreflect.FileDescriptor

var file_faultinjector_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x22, 0x1a, 0x0a, 0x18, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x4f, 0x0a, 0x19, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x22, 0x4e, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x22, 0x3c, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x68, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x61, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x46, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x41, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x22, 0x56, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xfa, 0x02, 0x0a,
	0x0d, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x68,
	0x0a, 0x11, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x27, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x2e, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x46, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x21, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46,
	0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49,
	0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e,
	0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_faultinjector_proto_rawDescOnce sync.Once
	file_faultinjector_proto_rawDescData = file_faultinjector_proto_rawDesc
)

func file_faultinjector_proto_rawDescGZIP() []byte {
	file_faultinjector_proto_rawDescOnce.Do(func() {
		file_faultinjector_proto_rawDescData = protoimpl.X.CompressGZIP(file_faultinjector_proto_rawDescData)
	})
	return file_faultinjector_proto_rawDescData
}

var file_faultinjector_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_faultinjector_proto_goTypes = []interface{}{
	(*EnumerateServicesRequest)(nil),  // 0: faultinjector.EnumerateServicesRequest
	(*EnumerateServicesResponse)(nil), // 1: faultinjector.EnumerateServicesResponse
	(*Service)(nil),                   // 2: faultinjector.Service
	(*Method)(nil),                    // 3: faultinjector.Method
	(*RegisterFaultRequest)(nil),      // 4: faultinjector.RegisterFaultRequest
	(*RegisterFaultResponse)(nil),     // 5: faultinjector.RegisterFaultResponse
	(*RemoveFaultRequest)(nil),        // 6: faultinjector.RemoveFaultRequest
	(*RemoveFaultResponse)(nil),       // 7: faultinjector.RemoveFaultResponse
	(*ListenRequest)(nil),             // 8: faultinjector.ListenRequest
	(*ListenResponse)(nil),            // 9: faultinjector.ListenResponse
}
var file_faultinjector_proto_depIdxs = []int32{
	2, // 0: faultinjector.EnumerateServicesResponse.services:type_name -> faultinjector.Service
	3, // 1: faultinjector.Service.methods:type_name -> faultinjector.Method
	0, // 2: faultinjector.FaultInjector.EnumerateServices:input_type -> faultinjector.EnumerateServicesRequest
	4, // 3: faultinjector.FaultInjector.RegisterFault:input_type -> faultinjector.RegisterFaultRequest
	6, // 4: faultinjector.FaultInjector.RemoveFault:input_type -> faultinjector.RemoveFaultRequest
	8, // 5: faultinjector.FaultInjector.Listen:input_type -> faultinjector.ListenRequest
	1, // 6: faultinjector.FaultInjector.EnumerateServices:output_type -> faultinjector.EnumerateServicesResponse
	5, // 7: faultinjector.FaultInjector.RegisterFault:output_type -> faultinjector.RegisterFaultResponse
	7, // 8: faultinjector.FaultInjector.RemoveFault:output_type -> faultinjector.RemoveFaultResponse
	9, // 9: faultinjector.FaultInjector.Listen:output_type -> faultinjector.ListenResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_faultinjector_proto_init() }
func file_faultinjector_proto_init() {
	if File_faultinjector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_faultinjector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumerateServicesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_faultinjector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumerateServicesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_faultinjector_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_faultinjector_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Method); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_faultinjector_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterFaultRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_faultinjector_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterFaultResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_faultinjector_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFaultRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_faultinjector_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFaultResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_faultinjector_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_faultinjector_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_faultinjector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_faultinjector_proto_goTypes,
		DependencyIndexes: file_faultinjector_proto_depIdxs,
		MessageInfos:      file_faultinjector_proto_msgTypes,
	}.Build()
	File_faultinjector_proto = out.File
	file_faultinjector_proto_rawDesc = nil
	file_faultinjector_proto_goTypes = nil
	file_faultinjector_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FaultInjectorClient is the client API for FaultInjector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FaultInjectorClient interface {
	// Sends a greeting
	EnumerateServices(ctx context.Context, in *EnumerateServicesRequest, opts ...grpc.CallOption) (*EnumerateServicesResponse, error)
	RegisterFault(ctx context.Context, in *RegisterFaultRequest, opts ...grpc.CallOption) (*RegisterFaultResponse, error)
	RemoveFault(ctx context.Context, in *RemoveFaultRequest, opts ...grpc.CallOption) (*RemoveFaultResponse, error)
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (FaultInjector_ListenClient, error)
}

type faultInjectorClient struct {
	cc grpc.ClientConnInterface
}

func NewFaultInjectorClient(cc grpc.ClientConnInterface) FaultInjectorClient {
	return &faultInjectorClient{cc}
}

func (c *faultInjectorClient) EnumerateServices(ctx context.Context, in *EnumerateServicesRequest, opts ...grpc.CallOption) (*EnumerateServicesResponse, error) {
	out := new(EnumerateServicesResponse)
	err := c.cc.Invoke(ctx, "/faultinjector.FaultInjector/EnumerateServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faultInjectorClient) RegisterFault(ctx context.Context, in *RegisterFaultRequest, opts ...grpc.CallOption) (*RegisterFaultResponse, error) {
	out := new(RegisterFaultResponse)
	err := c.cc.Invoke(ctx, "/faultinjector.FaultInjector/RegisterFault", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faultInjectorClient) RemoveFault(ctx context.Context, in *RemoveFaultRequest, opts ...grpc.CallOption) (*RemoveFaultResponse, error) {
	out := new(RemoveFaultResponse)
	err := c.cc.Invoke(ctx, "/faultinjector.FaultInjector/RemoveFault", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faultInjectorClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (FaultInjector_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FaultInjector_serviceDesc.Streams[0], "/faultinjector.FaultInjector/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &faultInjectorListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FaultInjector_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type faultInjectorListenClient struct {
	grpc.ClientStream
}

func (x *faultInjectorListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FaultInjectorServer is the server API for FaultInjector service.
type FaultInjectorServer interface {
	// Sends a greeting
	EnumerateServices(context.Context, *EnumerateServicesRequest) (*EnumerateServicesResponse, error)
	RegisterFault(context.Context, *RegisterFaultRequest) (*RegisterFaultResponse, error)
	RemoveFault(context.Context, *RemoveFaultRequest) (*RemoveFaultResponse, error)
	Listen(*ListenRequest, FaultInjector_ListenServer) error
}

// UnimplementedFaultInjectorServer can be embedded to have forward compatible implementations.
type UnimplementedFaultInjectorServer struct {
}

func (*UnimplementedFaultInjectorServer) EnumerateServices(context.Context, *EnumerateServicesRequest) (*EnumerateServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnumerateServices not implemented")
}
func (*UnimplementedFaultInjectorServer) RegisterFault(context.Context, *RegisterFaultRequest) (*RegisterFaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterFault not implemented")
}
func (*UnimplementedFaultInjectorServer) RemoveFault(context.Context, *RemoveFaultRequest) (*RemoveFaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFault not implemented")
}
func (*UnimplementedFaultInjectorServer) Listen(*ListenRequest, FaultInjector_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}

func RegisterFaultInjectorServer(s *grpc.Server, srv FaultInjectorServer) {
	s.RegisterService(&_FaultInjector_serviceDesc, srv)
}

func _FaultInjector_EnumerateServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnumerateServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaultInjectorServer).EnumerateServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faultinjector.FaultInjector/EnumerateServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaultInjectorServer).EnumerateServices(ctx, req.(*EnumerateServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaultInjector_RegisterFault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterFaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaultInjectorServer).RegisterFault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faultinjector.FaultInjector/RegisterFault",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaultInjectorServer).RegisterFault(ctx, req.(*RegisterFaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaultInjector_RemoveFault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaultInjectorServer).RemoveFault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faultinjector.FaultInjector/RemoveFault",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaultInjectorServer).RemoveFault(ctx, req.(*RemoveFaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaultInjector_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FaultInjectorServer).Listen(m, &faultInjectorListenServer{stream})
}

type FaultInjector_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type faultInjectorListenServer struct {
	grpc.ServerStream
}

func (x *faultInjectorListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _FaultInjector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "faultinjector.FaultInjector",
	HandlerType: (*FaultInjectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnumerateServices",
			Handler:    _FaultInjector_EnumerateServices_Handler,
		},
		{
			MethodName: "RegisterFault",
			Handler:    _FaultInjector_RegisterFault_Handler,
		},
		{
			MethodName: "RemoveFault",
			Handler:    _FaultInjector_RemoveFault_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _FaultInjector_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "faultinjector.proto",
}
