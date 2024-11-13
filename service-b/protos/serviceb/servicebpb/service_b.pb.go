// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: protos/service_b.proto

package servicebpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CallServiceBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CallServiceBRequest) Reset() {
	*x = CallServiceBRequest{}
	mi := &file_protos_service_b_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CallServiceBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallServiceBRequest) ProtoMessage() {}

func (x *CallServiceBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_b_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallServiceBRequest.ProtoReflect.Descriptor instead.
func (*CallServiceBRequest) Descriptor() ([]byte, []int) {
	return file_protos_service_b_proto_rawDescGZIP(), []int{0}
}

type CallServiceBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CallServiceBResponse) Reset() {
	*x = CallServiceBResponse{}
	mi := &file_protos_service_b_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CallServiceBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallServiceBResponse) ProtoMessage() {}

func (x *CallServiceBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_b_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallServiceBResponse.ProtoReflect.Descriptor instead.
func (*CallServiceBResponse) Descriptor() ([]byte, []int) {
	return file_protos_service_b_proto_rawDescGZIP(), []int{1}
}

func (x *CallServiceBResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_service_b_proto protoreflect.FileDescriptor

var file_protos_service_b_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x14, 0x43, 0x61, 0x6c,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x59, 0x0a, 0x08, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x12, 0x4d, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_service_b_proto_rawDescOnce sync.Once
	file_protos_service_b_proto_rawDescData = file_protos_service_b_proto_rawDesc
)

func file_protos_service_b_proto_rawDescGZIP() []byte {
	file_protos_service_b_proto_rawDescOnce.Do(func() {
		file_protos_service_b_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_service_b_proto_rawDescData)
	})
	return file_protos_service_b_proto_rawDescData
}

var file_protos_service_b_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_service_b_proto_goTypes = []any{
	(*CallServiceBRequest)(nil),  // 0: serviceb.CallServiceBRequest
	(*CallServiceBResponse)(nil), // 1: serviceb.CallServiceBResponse
}
var file_protos_service_b_proto_depIdxs = []int32{
	0, // 0: serviceb.ServiceB.CallServiceB:input_type -> serviceb.CallServiceBRequest
	1, // 1: serviceb.ServiceB.CallServiceB:output_type -> serviceb.CallServiceBResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_service_b_proto_init() }
func file_protos_service_b_proto_init() {
	if File_protos_service_b_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_service_b_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_service_b_proto_goTypes,
		DependencyIndexes: file_protos_service_b_proto_depIdxs,
		MessageInfos:      file_protos_service_b_proto_msgTypes,
	}.Build()
	File_protos_service_b_proto = out.File
	file_protos_service_b_proto_rawDesc = nil
	file_protos_service_b_proto_goTypes = nil
	file_protos_service_b_proto_depIdxs = nil
}
