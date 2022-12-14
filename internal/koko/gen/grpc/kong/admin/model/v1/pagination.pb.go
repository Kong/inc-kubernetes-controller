// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: kong/admin/model/v1/pagination.proto

package v1

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

type PaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size   int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Number int32 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	// Allows callers to provide a CEL expression to filter results on `list` API calls.
	//
	// For example, given a resource with the following tags, `tag1` & `tag2`, the
	// following CEL expressions are supported:
	//
	// - Matches resources that have `tag1` as any tag:
	//   - `"tag1" in tags`
	//
	// - Matches all resources that have both `tag1` & `tag2`:
	//   - `["tag1", "tag2"].all(x, x in tags)`
	//   - `"tag1" in tags && "tag2" in tags`
	//
	// - Matches resources that have `tag1` or `tag2`:
	//   - `["tag1", "tag2"].exists(x, x in tags)`
	//   - `"tag1" in tags || "tag2" in tags`
	//
	// Limitations:
	// Currently, it is only possible to filter on tags, and supported logical
	// operators/macros are limited to only what is documented above.
	//
	// For further information, you may view the CEL Specification:
	// https://github.com/google/cel-spec
	Filter string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_model_v1_pagination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_model_v1_pagination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_model_v1_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *PaginationRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PaginationRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *PaginationRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type PaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount  int32 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	NextPageNum int32 `protobuf:"varint,2,opt,name=next_page_num,json=nextPageNum,proto3" json:"next_page_num,omitempty"`
}

func (x *PaginationResponse) Reset() {
	*x = PaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_model_v1_pagination_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResponse) ProtoMessage() {}

func (x *PaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_model_v1_pagination_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationResponse.ProtoReflect.Descriptor instead.
func (*PaginationResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_model_v1_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *PaginationResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *PaginationResponse) GetNextPageNum() int32 {
	if x != nil {
		return x.NextPageNum
	}
	return 0
}

var File_kong_admin_model_v1_pagination_proto protoreflect.FileDescriptor

var file_kong_admin_model_v1_pagination_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x57, 0x0a, 0x11, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x22, 0x59, 0x0a, 0x12, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x42,
	0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f,
	0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kong_admin_model_v1_pagination_proto_rawDescOnce sync.Once
	file_kong_admin_model_v1_pagination_proto_rawDescData = file_kong_admin_model_v1_pagination_proto_rawDesc
)

func file_kong_admin_model_v1_pagination_proto_rawDescGZIP() []byte {
	file_kong_admin_model_v1_pagination_proto_rawDescOnce.Do(func() {
		file_kong_admin_model_v1_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_kong_admin_model_v1_pagination_proto_rawDescData)
	})
	return file_kong_admin_model_v1_pagination_proto_rawDescData
}

var file_kong_admin_model_v1_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kong_admin_model_v1_pagination_proto_goTypes = []interface{}{
	(*PaginationRequest)(nil),  // 0: kong.admin.model.v1.PaginationRequest
	(*PaginationResponse)(nil), // 1: kong.admin.model.v1.PaginationResponse
}
var file_kong_admin_model_v1_pagination_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kong_admin_model_v1_pagination_proto_init() }
func file_kong_admin_model_v1_pagination_proto_init() {
	if File_kong_admin_model_v1_pagination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kong_admin_model_v1_pagination_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationRequest); i {
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
		file_kong_admin_model_v1_pagination_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationResponse); i {
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
			RawDescriptor: file_kong_admin_model_v1_pagination_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kong_admin_model_v1_pagination_proto_goTypes,
		DependencyIndexes: file_kong_admin_model_v1_pagination_proto_depIdxs,
		MessageInfos:      file_kong_admin_model_v1_pagination_proto_msgTypes,
	}.Build()
	File_kong_admin_model_v1_pagination_proto = out.File
	file_kong_admin_model_v1_pagination_proto_rawDesc = nil
	file_kong_admin_model_v1_pagination_proto_goTypes = nil
	file_kong_admin_model_v1_pagination_proto_depIdxs = nil
}