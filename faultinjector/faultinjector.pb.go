// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: faultinjector.proto

package faultinjector

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

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
	0x65, 0x32, 0xaf, 0x02, 0x0a, 0x0d, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x6a, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x68, 0x0a, 0x11, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x27, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a,
	0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x23,
	0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0b, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x21, 0x2e, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x65, 0x74, 0x66, 0x6c, 0x69, 0x78, 0x2d, 0x73, 0x6b, 0x75, 0x6e, 0x6b, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2f,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_faultinjector_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_faultinjector_proto_goTypes = []interface{}{
	(*EnumerateServicesRequest)(nil),  // 0: faultinjector.EnumerateServicesRequest
	(*EnumerateServicesResponse)(nil), // 1: faultinjector.EnumerateServicesResponse
	(*Service)(nil),                   // 2: faultinjector.Service
	(*Method)(nil),                    // 3: faultinjector.Method
	(*RegisterFaultRequest)(nil),      // 4: faultinjector.RegisterFaultRequest
	(*RegisterFaultResponse)(nil),     // 5: faultinjector.RegisterFaultResponse
	(*RemoveFaultRequest)(nil),        // 6: faultinjector.RemoveFaultRequest
	(*RemoveFaultResponse)(nil),       // 7: faultinjector.RemoveFaultResponse
}
var file_faultinjector_proto_depIdxs = []int32{
	2, // 0: faultinjector.EnumerateServicesResponse.services:type_name -> faultinjector.Service
	3, // 1: faultinjector.Service.methods:type_name -> faultinjector.Method
	0, // 2: faultinjector.FaultInjector.EnumerateServices:input_type -> faultinjector.EnumerateServicesRequest
	4, // 3: faultinjector.FaultInjector.RegisterFault:input_type -> faultinjector.RegisterFaultRequest
	6, // 4: faultinjector.FaultInjector.RemoveFault:input_type -> faultinjector.RemoveFaultRequest
	1, // 5: faultinjector.FaultInjector.EnumerateServices:output_type -> faultinjector.EnumerateServicesResponse
	5, // 6: faultinjector.FaultInjector.RegisterFault:output_type -> faultinjector.RegisterFaultResponse
	7, // 7: faultinjector.FaultInjector.RemoveFault:output_type -> faultinjector.RemoveFaultResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_faultinjector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
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
