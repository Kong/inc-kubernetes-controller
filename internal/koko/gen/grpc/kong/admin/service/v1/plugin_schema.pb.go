// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: kong/admin/service/v1/plugin_schema.proto

package v1

import (
	v1 "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/admin/model/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateLuaPluginSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item    *v1.PluginSchema   `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *CreateLuaPluginSchemaRequest) Reset() {
	*x = CreateLuaPluginSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLuaPluginSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLuaPluginSchemaRequest) ProtoMessage() {}

func (x *CreateLuaPluginSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLuaPluginSchemaRequest.ProtoReflect.Descriptor instead.
func (*CreateLuaPluginSchemaRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_plugin_schema_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLuaPluginSchemaRequest) GetItem() *v1.PluginSchema {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *CreateLuaPluginSchemaRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type CreateLuaPluginSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *v1.PluginSchema `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *CreateLuaPluginSchemaResponse) Reset() {
	*x = CreateLuaPluginSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLuaPluginSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLuaPluginSchemaResponse) ProtoMessage() {}

func (x *CreateLuaPluginSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLuaPluginSchemaResponse.ProtoReflect.Descriptor instead.
func (*CreateLuaPluginSchemaResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_plugin_schema_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLuaPluginSchemaResponse) GetItem() *v1.PluginSchema {
	if x != nil {
		return x.Item
	}
	return nil
}

type GetLuaPluginSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *GetLuaPluginSchemaRequest) Reset() {
	*x = GetLuaPluginSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLuaPluginSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLuaPluginSchemaRequest) ProtoMessage() {}

func (x *GetLuaPluginSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLuaPluginSchemaRequest.ProtoReflect.Descriptor instead.
func (*GetLuaPluginSchemaRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_plugin_schema_proto_rawDescGZIP(), []int{2}
}

func (x *GetLuaPluginSchemaRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetLuaPluginSchemaRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type GetLuaPluginSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *v1.PluginSchema `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetLuaPluginSchemaResponse) Reset() {
	*x = GetLuaPluginSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLuaPluginSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLuaPluginSchemaResponse) ProtoMessage() {}

func (x *GetLuaPluginSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLuaPluginSchemaResponse.ProtoReflect.Descriptor instead.
func (*GetLuaPluginSchemaResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_plugin_schema_proto_rawDescGZIP(), []int{3}
}

func (x *GetLuaPluginSchemaResponse) GetItem() *v1.PluginSchema {
	if x != nil {
		return x.Item
	}
	return nil
}

type ListLuaPluginSchemasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster *v1.RequestCluster    `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Page    *v1.PaginationRequest `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListLuaPluginSchemasRequest) Reset() {
	*x = ListLuaPluginSchemasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLuaPluginSchemasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLuaPluginSchemasRequest) ProtoMessage() {}

func (x *ListLuaPluginSchemasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLuaPluginSchemasRequest.ProtoReflect.Descriptor instead.
func (*ListLuaPluginSchemasRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_plugin_schema_proto_rawDescGZIP(), []int{4}
}

func (x *ListLuaPluginSchemasRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *ListLuaPluginSchemasRequest) GetPage() *v1.PaginationRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

type ListLuaPluginSchemasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*v1.PluginSchema     `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Page  *v1.PaginationResponse `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListLuaPluginSchemasResponse) Reset() {
	*x = ListLuaPluginSchemasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLuaPluginSchemasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLuaPluginSchemasResponse) ProtoMessage() {}

func (x *ListLuaPluginSchemasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLuaPluginSchemasResponse.ProtoReflect.Descriptor instead.
func (*ListLuaPluginSchemasResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_plugin_schema_proto_rawDescGZIP(), []int{5}
}

func (x *ListLuaPluginSchemasResponse) GetItems() []*v1.PluginSchema {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListLuaPluginSchemasResponse) GetPage() *v1.PaginationResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

type UpsertLuaPluginSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Item    *v1.PluginSchema   `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *UpsertLuaPluginSchemaRequest) Reset() {
	*x = UpsertLuaPluginSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertLuaPluginSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertLuaPluginSchemaRequest) ProtoMessage() {}

func (x *UpsertLuaPluginSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertLuaPluginSchemaRequest.ProtoReflect.Descriptor instead.
func (*UpsertLuaPluginSchemaRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_plugin_schema_proto_rawDescGZIP(), []int{6}
}

func (x *UpsertLuaPluginSchemaRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpsertLuaPluginSchemaRequest) GetItem() *v1.PluginSchema {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *UpsertLuaPluginSchemaRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type UpsertLuaPluginSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *v1.PluginSchema `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *UpsertLuaPluginSchemaResponse) Reset() {
	*x = UpsertLuaPluginSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertLuaPluginSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertLuaPluginSchemaResponse) ProtoMessage() {}

func (x *UpsertLuaPluginSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertLuaPluginSchemaResponse.ProtoReflect.Descriptor instead.
func (*UpsertLuaPluginSchemaResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_plugin_schema_proto_rawDescGZIP(), []int{7}
}

func (x *UpsertLuaPluginSchemaResponse) GetItem() *v1.PluginSchema {
	if x != nil {
		return x.Item
	}
	return nil
}

type DeleteLuaPluginSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *DeleteLuaPluginSchemaRequest) Reset() {
	*x = DeleteLuaPluginSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLuaPluginSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLuaPluginSchemaRequest) ProtoMessage() {}

func (x *DeleteLuaPluginSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLuaPluginSchemaRequest.ProtoReflect.Descriptor instead.
func (*DeleteLuaPluginSchemaRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_plugin_schema_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteLuaPluginSchemaRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteLuaPluginSchemaRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type DeleteLuaPluginSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteLuaPluginSchemaResponse) Reset() {
	*x = DeleteLuaPluginSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLuaPluginSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLuaPluginSchemaResponse) ProtoMessage() {}

func (x *DeleteLuaPluginSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_plugin_schema_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLuaPluginSchemaResponse.ProtoReflect.Descriptor instead.
func (*DeleteLuaPluginSchemaResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_plugin_schema_proto_rawDescGZIP(), []int{9}
}

var File_kong_admin_service_v1_plugin_schema_proto protoreflect.FileDescriptor

var file_kong_admin_service_v1_plugin_schema_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6b, 0x6f, 0x6e,
	0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6b, 0x6f, 0x6e, 0x67, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x75, 0x61,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f,
	0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x56, 0x0a, 0x1d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x22, 0x6e, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x22, 0x53, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x98, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x4c,
	0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x22, 0x94, 0x01, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3b, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6b, 0x6f, 0x6e, 0x67,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x1c, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6b, 0x6f,
	0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x22, 0x56, 0x0a, 0x1d, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x75, 0x61,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x71, 0x0a, 0x1c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x1f,
	0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xcf, 0x06, 0x0a, 0x13, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa4, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x12, 0x33, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0x9c,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x30, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x9b, 0x01,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0x32, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6b, 0x6f, 0x6e,
	0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0xab, 0x01, 0x0a, 0x15,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x33, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x6b, 0x6f, 0x6e,
	0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x1a, 0x19,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0xa5, 0x01, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x33, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x75, 0x61, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x2a, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kong_admin_service_v1_plugin_schema_proto_rawDescOnce sync.Once
	file_kong_admin_service_v1_plugin_schema_proto_rawDescData = file_kong_admin_service_v1_plugin_schema_proto_rawDesc
)

func file_kong_admin_service_v1_plugin_schema_proto_rawDescGZIP() []byte {
	file_kong_admin_service_v1_plugin_schema_proto_rawDescOnce.Do(func() {
		file_kong_admin_service_v1_plugin_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_kong_admin_service_v1_plugin_schema_proto_rawDescData)
	})
	return file_kong_admin_service_v1_plugin_schema_proto_rawDescData
}

var file_kong_admin_service_v1_plugin_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_kong_admin_service_v1_plugin_schema_proto_goTypes = []interface{}{
	(*CreateLuaPluginSchemaRequest)(nil),  // 0: kong.admin.service.v1.CreateLuaPluginSchemaRequest
	(*CreateLuaPluginSchemaResponse)(nil), // 1: kong.admin.service.v1.CreateLuaPluginSchemaResponse
	(*GetLuaPluginSchemaRequest)(nil),     // 2: kong.admin.service.v1.GetLuaPluginSchemaRequest
	(*GetLuaPluginSchemaResponse)(nil),    // 3: kong.admin.service.v1.GetLuaPluginSchemaResponse
	(*ListLuaPluginSchemasRequest)(nil),   // 4: kong.admin.service.v1.ListLuaPluginSchemasRequest
	(*ListLuaPluginSchemasResponse)(nil),  // 5: kong.admin.service.v1.ListLuaPluginSchemasResponse
	(*UpsertLuaPluginSchemaRequest)(nil),  // 6: kong.admin.service.v1.UpsertLuaPluginSchemaRequest
	(*UpsertLuaPluginSchemaResponse)(nil), // 7: kong.admin.service.v1.UpsertLuaPluginSchemaResponse
	(*DeleteLuaPluginSchemaRequest)(nil),  // 8: kong.admin.service.v1.DeleteLuaPluginSchemaRequest
	(*DeleteLuaPluginSchemaResponse)(nil), // 9: kong.admin.service.v1.DeleteLuaPluginSchemaResponse
	(*v1.PluginSchema)(nil),               // 10: kong.admin.model.v1.PluginSchema
	(*v1.RequestCluster)(nil),             // 11: kong.admin.model.v1.RequestCluster
	(*v1.PaginationRequest)(nil),          // 12: kong.admin.model.v1.PaginationRequest
	(*v1.PaginationResponse)(nil),         // 13: kong.admin.model.v1.PaginationResponse
}
var file_kong_admin_service_v1_plugin_schema_proto_depIdxs = []int32{
	10, // 0: kong.admin.service.v1.CreateLuaPluginSchemaRequest.item:type_name -> kong.admin.model.v1.PluginSchema
	11, // 1: kong.admin.service.v1.CreateLuaPluginSchemaRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	10, // 2: kong.admin.service.v1.CreateLuaPluginSchemaResponse.item:type_name -> kong.admin.model.v1.PluginSchema
	11, // 3: kong.admin.service.v1.GetLuaPluginSchemaRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	10, // 4: kong.admin.service.v1.GetLuaPluginSchemaResponse.item:type_name -> kong.admin.model.v1.PluginSchema
	11, // 5: kong.admin.service.v1.ListLuaPluginSchemasRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	12, // 6: kong.admin.service.v1.ListLuaPluginSchemasRequest.page:type_name -> kong.admin.model.v1.PaginationRequest
	10, // 7: kong.admin.service.v1.ListLuaPluginSchemasResponse.items:type_name -> kong.admin.model.v1.PluginSchema
	13, // 8: kong.admin.service.v1.ListLuaPluginSchemasResponse.page:type_name -> kong.admin.model.v1.PaginationResponse
	10, // 9: kong.admin.service.v1.UpsertLuaPluginSchemaRequest.item:type_name -> kong.admin.model.v1.PluginSchema
	11, // 10: kong.admin.service.v1.UpsertLuaPluginSchemaRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	10, // 11: kong.admin.service.v1.UpsertLuaPluginSchemaResponse.item:type_name -> kong.admin.model.v1.PluginSchema
	11, // 12: kong.admin.service.v1.DeleteLuaPluginSchemaRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	0,  // 13: kong.admin.service.v1.PluginSchemaService.CreateLuaPluginSchema:input_type -> kong.admin.service.v1.CreateLuaPluginSchemaRequest
	2,  // 14: kong.admin.service.v1.PluginSchemaService.GetLuaPluginSchema:input_type -> kong.admin.service.v1.GetLuaPluginSchemaRequest
	4,  // 15: kong.admin.service.v1.PluginSchemaService.ListLuaPluginSchemas:input_type -> kong.admin.service.v1.ListLuaPluginSchemasRequest
	6,  // 16: kong.admin.service.v1.PluginSchemaService.UpsertLuaPluginSchema:input_type -> kong.admin.service.v1.UpsertLuaPluginSchemaRequest
	8,  // 17: kong.admin.service.v1.PluginSchemaService.DeleteLuaPluginSchema:input_type -> kong.admin.service.v1.DeleteLuaPluginSchemaRequest
	1,  // 18: kong.admin.service.v1.PluginSchemaService.CreateLuaPluginSchema:output_type -> kong.admin.service.v1.CreateLuaPluginSchemaResponse
	3,  // 19: kong.admin.service.v1.PluginSchemaService.GetLuaPluginSchema:output_type -> kong.admin.service.v1.GetLuaPluginSchemaResponse
	5,  // 20: kong.admin.service.v1.PluginSchemaService.ListLuaPluginSchemas:output_type -> kong.admin.service.v1.ListLuaPluginSchemasResponse
	7,  // 21: kong.admin.service.v1.PluginSchemaService.UpsertLuaPluginSchema:output_type -> kong.admin.service.v1.UpsertLuaPluginSchemaResponse
	9,  // 22: kong.admin.service.v1.PluginSchemaService.DeleteLuaPluginSchema:output_type -> kong.admin.service.v1.DeleteLuaPluginSchemaResponse
	18, // [18:23] is the sub-list for method output_type
	13, // [13:18] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_kong_admin_service_v1_plugin_schema_proto_init() }
func file_kong_admin_service_v1_plugin_schema_proto_init() {
	if File_kong_admin_service_v1_plugin_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kong_admin_service_v1_plugin_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLuaPluginSchemaRequest); i {
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
		file_kong_admin_service_v1_plugin_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLuaPluginSchemaResponse); i {
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
		file_kong_admin_service_v1_plugin_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLuaPluginSchemaRequest); i {
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
		file_kong_admin_service_v1_plugin_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLuaPluginSchemaResponse); i {
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
		file_kong_admin_service_v1_plugin_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLuaPluginSchemasRequest); i {
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
		file_kong_admin_service_v1_plugin_schema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLuaPluginSchemasResponse); i {
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
		file_kong_admin_service_v1_plugin_schema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertLuaPluginSchemaRequest); i {
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
		file_kong_admin_service_v1_plugin_schema_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertLuaPluginSchemaResponse); i {
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
		file_kong_admin_service_v1_plugin_schema_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLuaPluginSchemaRequest); i {
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
		file_kong_admin_service_v1_plugin_schema_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLuaPluginSchemaResponse); i {
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
			RawDescriptor: file_kong_admin_service_v1_plugin_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kong_admin_service_v1_plugin_schema_proto_goTypes,
		DependencyIndexes: file_kong_admin_service_v1_plugin_schema_proto_depIdxs,
		MessageInfos:      file_kong_admin_service_v1_plugin_schema_proto_msgTypes,
	}.Build()
	File_kong_admin_service_v1_plugin_schema_proto = out.File
	file_kong_admin_service_v1_plugin_schema_proto_rawDesc = nil
	file_kong_admin_service_v1_plugin_schema_proto_goTypes = nil
	file_kong_admin_service_v1_plugin_schema_proto_depIdxs = nil
}
