// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: kong/admin/service/v1/sni.proto

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

type GetSNIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *GetSNIRequest) Reset() {
	*x = GetSNIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_sni_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSNIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSNIRequest) ProtoMessage() {}

func (x *GetSNIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_sni_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSNIRequest.ProtoReflect.Descriptor instead.
func (*GetSNIRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_sni_proto_rawDescGZIP(), []int{0}
}

func (x *GetSNIRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetSNIRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type GetSNIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *v1.SNI `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetSNIResponse) Reset() {
	*x = GetSNIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_sni_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSNIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSNIResponse) ProtoMessage() {}

func (x *GetSNIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_sni_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSNIResponse.ProtoReflect.Descriptor instead.
func (*GetSNIResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_sni_proto_rawDescGZIP(), []int{1}
}

func (x *GetSNIResponse) GetItem() *v1.SNI {
	if x != nil {
		return x.Item
	}
	return nil
}

type CreateSNIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item    *v1.SNI            `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *CreateSNIRequest) Reset() {
	*x = CreateSNIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_sni_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSNIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSNIRequest) ProtoMessage() {}

func (x *CreateSNIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_sni_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSNIRequest.ProtoReflect.Descriptor instead.
func (*CreateSNIRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_sni_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSNIRequest) GetItem() *v1.SNI {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *CreateSNIRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type CreateSNIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *v1.SNI `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *CreateSNIResponse) Reset() {
	*x = CreateSNIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_sni_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSNIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSNIResponse) ProtoMessage() {}

func (x *CreateSNIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_sni_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSNIResponse.ProtoReflect.Descriptor instead.
func (*CreateSNIResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_sni_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSNIResponse) GetItem() *v1.SNI {
	if x != nil {
		return x.Item
	}
	return nil
}

type UpsertSNIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item    *v1.SNI            `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *UpsertSNIRequest) Reset() {
	*x = UpsertSNIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_sni_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertSNIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertSNIRequest) ProtoMessage() {}

func (x *UpsertSNIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_sni_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertSNIRequest.ProtoReflect.Descriptor instead.
func (*UpsertSNIRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_sni_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertSNIRequest) GetItem() *v1.SNI {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *UpsertSNIRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type UpsertSNIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *v1.SNI `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *UpsertSNIResponse) Reset() {
	*x = UpsertSNIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_sni_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertSNIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertSNIResponse) ProtoMessage() {}

func (x *UpsertSNIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_sni_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertSNIResponse.ProtoReflect.Descriptor instead.
func (*UpsertSNIResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_sni_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertSNIResponse) GetItem() *v1.SNI {
	if x != nil {
		return x.Item
	}
	return nil
}

type DeleteSNIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *DeleteSNIRequest) Reset() {
	*x = DeleteSNIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_sni_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSNIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSNIRequest) ProtoMessage() {}

func (x *DeleteSNIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_sni_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSNIRequest.ProtoReflect.Descriptor instead.
func (*DeleteSNIRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_sni_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSNIRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteSNIRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type DeleteSNIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSNIResponse) Reset() {
	*x = DeleteSNIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_sni_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSNIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSNIResponse) ProtoMessage() {}

func (x *DeleteSNIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_sni_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSNIResponse.ProtoReflect.Descriptor instead.
func (*DeleteSNIResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_sni_proto_rawDescGZIP(), []int{7}
}

type ListSNIsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster       *v1.RequestCluster    `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Page          *v1.PaginationRequest `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	CertificateId string                `protobuf:"bytes,3,opt,name=certificate_id,json=certificateId,proto3" json:"certificate_id,omitempty"`
}

func (x *ListSNIsRequest) Reset() {
	*x = ListSNIsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_sni_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSNIsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSNIsRequest) ProtoMessage() {}

func (x *ListSNIsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_sni_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSNIsRequest.ProtoReflect.Descriptor instead.
func (*ListSNIsRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_sni_proto_rawDescGZIP(), []int{8}
}

func (x *ListSNIsRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *ListSNIsRequest) GetPage() *v1.PaginationRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListSNIsRequest) GetCertificateId() string {
	if x != nil {
		return x.CertificateId
	}
	return ""
}

type ListSNIsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*v1.SNI              `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Page  *v1.PaginationResponse `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListSNIsResponse) Reset() {
	*x = ListSNIsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_sni_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSNIsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSNIsResponse) ProtoMessage() {}

func (x *ListSNIsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_sni_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSNIsResponse.ProtoReflect.Descriptor instead.
func (*ListSNIsResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_sni_proto_rawDescGZIP(), []int{9}
}

func (x *ListSNIsResponse) GetItems() []*v1.SNI {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListSNIsResponse) GetPage() *v1.PaginationResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_kong_admin_service_v1_sni_proto protoreflect.FileDescriptor

var file_kong_admin_service_v1_sni_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6b, 0x6f, 0x6e, 0x67, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x4e, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x3e,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x4e, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x4e, 0x49, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x7f,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x4e, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x4e, 0x49, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22,
	0x41, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x4e, 0x49, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x4e, 0x49, 0x52, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x22, 0x7f, 0x0a, 0x10, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x53, 0x4e, 0x49, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x4e, 0x49, 0x52, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x22, 0x41, 0x0a, 0x11, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x53, 0x4e, 0x49,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x4e, 0x49,
	0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x61, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x4e, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f,
	0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x4e, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb3,
	0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x4e, 0x49, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x3a, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x22, 0x7f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x4e, 0x49, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x4e,
	0x49, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0xdb, 0x04, 0x0a, 0x0a, 0x53, 0x4e, 0x49, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x53, 0x4e, 0x49, 0x12, 0x24,
	0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x4e, 0x49, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x4e, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e, 0x69, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x76, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x4e, 0x49, 0x12,
	0x27, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x4e,
	0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x4e, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e, 0x69, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x09, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x53, 0x4e, 0x49, 0x12, 0x27, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x53, 0x4e, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x53, 0x4e, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x3a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x1a, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x6e, 0x69, 0x73, 0x2f, 0x7b, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x69, 0x64, 0x7d, 0x12, 0x75, 0x0a,
	0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x4e, 0x49, 0x12, 0x27, 0x2e, 0x6b, 0x6f, 0x6e,
	0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x4e, 0x49, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x4e, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e, 0x69, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6d, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x4e, 0x49, 0x73,
	0x12, 0x26, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x4e, 0x49,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x4e, 0x49, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x6e, 0x69, 0x73, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kong_admin_service_v1_sni_proto_rawDescOnce sync.Once
	file_kong_admin_service_v1_sni_proto_rawDescData = file_kong_admin_service_v1_sni_proto_rawDesc
)

func file_kong_admin_service_v1_sni_proto_rawDescGZIP() []byte {
	file_kong_admin_service_v1_sni_proto_rawDescOnce.Do(func() {
		file_kong_admin_service_v1_sni_proto_rawDescData = protoimpl.X.CompressGZIP(file_kong_admin_service_v1_sni_proto_rawDescData)
	})
	return file_kong_admin_service_v1_sni_proto_rawDescData
}

var file_kong_admin_service_v1_sni_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_kong_admin_service_v1_sni_proto_goTypes = []interface{}{
	(*GetSNIRequest)(nil),         // 0: kong.admin.service.v1.GetSNIRequest
	(*GetSNIResponse)(nil),        // 1: kong.admin.service.v1.GetSNIResponse
	(*CreateSNIRequest)(nil),      // 2: kong.admin.service.v1.CreateSNIRequest
	(*CreateSNIResponse)(nil),     // 3: kong.admin.service.v1.CreateSNIResponse
	(*UpsertSNIRequest)(nil),      // 4: kong.admin.service.v1.UpsertSNIRequest
	(*UpsertSNIResponse)(nil),     // 5: kong.admin.service.v1.UpsertSNIResponse
	(*DeleteSNIRequest)(nil),      // 6: kong.admin.service.v1.DeleteSNIRequest
	(*DeleteSNIResponse)(nil),     // 7: kong.admin.service.v1.DeleteSNIResponse
	(*ListSNIsRequest)(nil),       // 8: kong.admin.service.v1.ListSNIsRequest
	(*ListSNIsResponse)(nil),      // 9: kong.admin.service.v1.ListSNIsResponse
	(*v1.RequestCluster)(nil),     // 10: kong.admin.model.v1.RequestCluster
	(*v1.SNI)(nil),                // 11: kong.admin.model.v1.SNI
	(*v1.PaginationRequest)(nil),  // 12: kong.admin.model.v1.PaginationRequest
	(*v1.PaginationResponse)(nil), // 13: kong.admin.model.v1.PaginationResponse
}
var file_kong_admin_service_v1_sni_proto_depIdxs = []int32{
	10, // 0: kong.admin.service.v1.GetSNIRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	11, // 1: kong.admin.service.v1.GetSNIResponse.item:type_name -> kong.admin.model.v1.SNI
	11, // 2: kong.admin.service.v1.CreateSNIRequest.item:type_name -> kong.admin.model.v1.SNI
	10, // 3: kong.admin.service.v1.CreateSNIRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	11, // 4: kong.admin.service.v1.CreateSNIResponse.item:type_name -> kong.admin.model.v1.SNI
	11, // 5: kong.admin.service.v1.UpsertSNIRequest.item:type_name -> kong.admin.model.v1.SNI
	10, // 6: kong.admin.service.v1.UpsertSNIRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	11, // 7: kong.admin.service.v1.UpsertSNIResponse.item:type_name -> kong.admin.model.v1.SNI
	10, // 8: kong.admin.service.v1.DeleteSNIRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	10, // 9: kong.admin.service.v1.ListSNIsRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	12, // 10: kong.admin.service.v1.ListSNIsRequest.page:type_name -> kong.admin.model.v1.PaginationRequest
	11, // 11: kong.admin.service.v1.ListSNIsResponse.items:type_name -> kong.admin.model.v1.SNI
	13, // 12: kong.admin.service.v1.ListSNIsResponse.page:type_name -> kong.admin.model.v1.PaginationResponse
	0,  // 13: kong.admin.service.v1.SNIService.GetSNI:input_type -> kong.admin.service.v1.GetSNIRequest
	2,  // 14: kong.admin.service.v1.SNIService.CreateSNI:input_type -> kong.admin.service.v1.CreateSNIRequest
	4,  // 15: kong.admin.service.v1.SNIService.UpsertSNI:input_type -> kong.admin.service.v1.UpsertSNIRequest
	6,  // 16: kong.admin.service.v1.SNIService.DeleteSNI:input_type -> kong.admin.service.v1.DeleteSNIRequest
	8,  // 17: kong.admin.service.v1.SNIService.ListSNIs:input_type -> kong.admin.service.v1.ListSNIsRequest
	1,  // 18: kong.admin.service.v1.SNIService.GetSNI:output_type -> kong.admin.service.v1.GetSNIResponse
	3,  // 19: kong.admin.service.v1.SNIService.CreateSNI:output_type -> kong.admin.service.v1.CreateSNIResponse
	5,  // 20: kong.admin.service.v1.SNIService.UpsertSNI:output_type -> kong.admin.service.v1.UpsertSNIResponse
	7,  // 21: kong.admin.service.v1.SNIService.DeleteSNI:output_type -> kong.admin.service.v1.DeleteSNIResponse
	9,  // 22: kong.admin.service.v1.SNIService.ListSNIs:output_type -> kong.admin.service.v1.ListSNIsResponse
	18, // [18:23] is the sub-list for method output_type
	13, // [13:18] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_kong_admin_service_v1_sni_proto_init() }
func file_kong_admin_service_v1_sni_proto_init() {
	if File_kong_admin_service_v1_sni_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kong_admin_service_v1_sni_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSNIRequest); i {
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
		file_kong_admin_service_v1_sni_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSNIResponse); i {
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
		file_kong_admin_service_v1_sni_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSNIRequest); i {
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
		file_kong_admin_service_v1_sni_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSNIResponse); i {
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
		file_kong_admin_service_v1_sni_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertSNIRequest); i {
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
		file_kong_admin_service_v1_sni_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertSNIResponse); i {
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
		file_kong_admin_service_v1_sni_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSNIRequest); i {
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
		file_kong_admin_service_v1_sni_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSNIResponse); i {
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
		file_kong_admin_service_v1_sni_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSNIsRequest); i {
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
		file_kong_admin_service_v1_sni_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSNIsResponse); i {
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
			RawDescriptor: file_kong_admin_service_v1_sni_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kong_admin_service_v1_sni_proto_goTypes,
		DependencyIndexes: file_kong_admin_service_v1_sni_proto_depIdxs,
		MessageInfos:      file_kong_admin_service_v1_sni_proto_msgTypes,
	}.Build()
	File_kong_admin_service_v1_sni_proto = out.File
	file_kong_admin_service_v1_sni_proto_rawDesc = nil
	file_kong_admin_service_v1_sni_proto_goTypes = nil
	file_kong_admin_service_v1_sni_proto_depIdxs = nil
}
