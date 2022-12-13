// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: kong/relay/service/v1/event.proto

package v1

import (
	v1 "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/admin/model/v1"
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

type FetchReconfigureEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster *v1.RequestCluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *FetchReconfigureEventsRequest) Reset() {
	*x = FetchReconfigureEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_relay_service_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchReconfigureEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchReconfigureEventsRequest) ProtoMessage() {}

func (x *FetchReconfigureEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_relay_service_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchReconfigureEventsRequest.ProtoReflect.Descriptor instead.
func (*FetchReconfigureEventsRequest) Descriptor() ([]byte, []int) {
	return file_kong_relay_service_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *FetchReconfigureEventsRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type FetchReconfigureEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchReconfigureEventsResponse) Reset() {
	*x = FetchReconfigureEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_relay_service_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchReconfigureEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchReconfigureEventsResponse) ProtoMessage() {}

func (x *FetchReconfigureEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_relay_service_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchReconfigureEventsResponse.ProtoReflect.Descriptor instead.
func (*FetchReconfigureEventsResponse) Descriptor() ([]byte, []int) {
	return file_kong_relay_service_v1_event_proto_rawDescGZIP(), []int{1}
}

var File_kong_relay_service_v1_event_proto protoreflect.FileDescriptor

var file_kong_relay_service_v1_event_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x6b, 0x6f, 0x6e, 0x67,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a,
	0x1d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d,
	0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x20, 0x0a,
	0x1e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x98, 0x01, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x87, 0x01, 0x0a, 0x16, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x34, 0x2e, 0x6b, 0x6f,
	0x6e, 0x67, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x35, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x6b, 0x6f,
	0x6b, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kong_relay_service_v1_event_proto_rawDescOnce sync.Once
	file_kong_relay_service_v1_event_proto_rawDescData = file_kong_relay_service_v1_event_proto_rawDesc
)

func file_kong_relay_service_v1_event_proto_rawDescGZIP() []byte {
	file_kong_relay_service_v1_event_proto_rawDescOnce.Do(func() {
		file_kong_relay_service_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_kong_relay_service_v1_event_proto_rawDescData)
	})
	return file_kong_relay_service_v1_event_proto_rawDescData
}

var file_kong_relay_service_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kong_relay_service_v1_event_proto_goTypes = []interface{}{
	(*FetchReconfigureEventsRequest)(nil),  // 0: kong.relay.service.v1.FetchReconfigureEventsRequest
	(*FetchReconfigureEventsResponse)(nil), // 1: kong.relay.service.v1.FetchReconfigureEventsResponse
	(*v1.RequestCluster)(nil),              // 2: kong.admin.model.v1.RequestCluster
}
var file_kong_relay_service_v1_event_proto_depIdxs = []int32{
	2, // 0: kong.relay.service.v1.FetchReconfigureEventsRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	0, // 1: kong.relay.service.v1.EventService.FetchReconfigureEvents:input_type -> kong.relay.service.v1.FetchReconfigureEventsRequest
	1, // 2: kong.relay.service.v1.EventService.FetchReconfigureEvents:output_type -> kong.relay.service.v1.FetchReconfigureEventsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kong_relay_service_v1_event_proto_init() }
func file_kong_relay_service_v1_event_proto_init() {
	if File_kong_relay_service_v1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kong_relay_service_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchReconfigureEventsRequest); i {
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
		file_kong_relay_service_v1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchReconfigureEventsResponse); i {
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
			RawDescriptor: file_kong_relay_service_v1_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kong_relay_service_v1_event_proto_goTypes,
		DependencyIndexes: file_kong_relay_service_v1_event_proto_depIdxs,
		MessageInfos:      file_kong_relay_service_v1_event_proto_msgTypes,
	}.Build()
	File_kong_relay_service_v1_event_proto = out.File
	file_kong_relay_service_v1_event_proto_rawDesc = nil
	file_kong_relay_service_v1_event_proto_goTypes = nil
	file_kong_relay_service_v1_event_proto_depIdxs = nil
}
