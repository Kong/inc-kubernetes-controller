// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: kong/admin/model/v1/plugin.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt int32                 `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int32                 `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Enabled   *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Protocols []string              `protobuf:"bytes,6,rep,name=protocols,proto3" json:"protocols,omitempty"`
	Tags      []string              `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	Service   *Service              `protobuf:"bytes,8,opt,name=service,proto3" json:"service,omitempty"`
	Route     *Route                `protobuf:"bytes,9,opt,name=route,proto3" json:"route,omitempty"`
	Config    *structpb.Struct      `protobuf:"bytes,10,opt,name=config,proto3" json:"config,omitempty"`
	Consumer  *Consumer             `protobuf:"bytes,11,opt,name=consumer,proto3" json:"consumer,omitempty"`
	Ordering  *Ordering             `protobuf:"bytes,12,opt,name=ordering,proto3" json:"ordering,omitempty"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_model_v1_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_model_v1_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_kong_admin_model_v1_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *Plugin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Plugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plugin) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Plugin) GetUpdatedAt() int32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Plugin) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *Plugin) GetProtocols() []string {
	if x != nil {
		return x.Protocols
	}
	return nil
}

func (x *Plugin) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Plugin) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *Plugin) GetRoute() *Route {
	if x != nil {
		return x.Route
	}
	return nil
}

func (x *Plugin) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Plugin) GetConsumer() *Consumer {
	if x != nil {
		return x.Consumer
	}
	return nil
}

func (x *Plugin) GetOrdering() *Ordering {
	if x != nil {
		return x.Ordering
	}
	return nil
}

type Ordering struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Before *Order `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	After  *Order `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *Ordering) Reset() {
	*x = Ordering{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_model_v1_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ordering) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ordering) ProtoMessage() {}

func (x *Ordering) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_model_v1_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ordering.ProtoReflect.Descriptor instead.
func (*Ordering) Descriptor() ([]byte, []int) {
	return file_kong_admin_model_v1_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *Ordering) GetBefore() *Order {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *Ordering) GetAfter() *Order {
	if x != nil {
		return x.After
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Access []string `protobuf:"bytes,1,rep,name=access,proto3" json:"access,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_model_v1_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_model_v1_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_kong_admin_model_v1_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetAccess() []string {
	if x != nil {
		return x.Access
	}
	return nil
}

var File_kong_admin_model_v1_plugin_proto protoreflect.FileDescriptor

var file_kong_admin_model_v1_plugin_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6b, 0x6f, 0x6e, 0x67,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x03,
	0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x6f, 0x6e, 0x67,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x39, 0x0a,
	0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x08,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x6f, 0x6e,
	0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x22, 0x70, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x32, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x1f, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kong_admin_model_v1_plugin_proto_rawDescOnce sync.Once
	file_kong_admin_model_v1_plugin_proto_rawDescData = file_kong_admin_model_v1_plugin_proto_rawDesc
)

func file_kong_admin_model_v1_plugin_proto_rawDescGZIP() []byte {
	file_kong_admin_model_v1_plugin_proto_rawDescOnce.Do(func() {
		file_kong_admin_model_v1_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_kong_admin_model_v1_plugin_proto_rawDescData)
	})
	return file_kong_admin_model_v1_plugin_proto_rawDescData
}

var file_kong_admin_model_v1_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kong_admin_model_v1_plugin_proto_goTypes = []interface{}{
	(*Plugin)(nil),               // 0: kong.admin.model.v1.Plugin
	(*Ordering)(nil),             // 1: kong.admin.model.v1.Ordering
	(*Order)(nil),                // 2: kong.admin.model.v1.Order
	(*wrapperspb.BoolValue)(nil), // 3: google.protobuf.BoolValue
	(*Service)(nil),              // 4: kong.admin.model.v1.Service
	(*Route)(nil),                // 5: kong.admin.model.v1.Route
	(*structpb.Struct)(nil),      // 6: google.protobuf.Struct
	(*Consumer)(nil),             // 7: kong.admin.model.v1.Consumer
}
var file_kong_admin_model_v1_plugin_proto_depIdxs = []int32{
	3, // 0: kong.admin.model.v1.Plugin.enabled:type_name -> google.protobuf.BoolValue
	4, // 1: kong.admin.model.v1.Plugin.service:type_name -> kong.admin.model.v1.Service
	5, // 2: kong.admin.model.v1.Plugin.route:type_name -> kong.admin.model.v1.Route
	6, // 3: kong.admin.model.v1.Plugin.config:type_name -> google.protobuf.Struct
	7, // 4: kong.admin.model.v1.Plugin.consumer:type_name -> kong.admin.model.v1.Consumer
	1, // 5: kong.admin.model.v1.Plugin.ordering:type_name -> kong.admin.model.v1.Ordering
	2, // 6: kong.admin.model.v1.Ordering.before:type_name -> kong.admin.model.v1.Order
	2, // 7: kong.admin.model.v1.Ordering.after:type_name -> kong.admin.model.v1.Order
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_kong_admin_model_v1_plugin_proto_init() }
func file_kong_admin_model_v1_plugin_proto_init() {
	if File_kong_admin_model_v1_plugin_proto != nil {
		return
	}
	file_kong_admin_model_v1_consumer_proto_init()
	file_kong_admin_model_v1_route_proto_init()
	file_kong_admin_model_v1_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kong_admin_model_v1_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
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
		file_kong_admin_model_v1_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ordering); i {
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
		file_kong_admin_model_v1_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_kong_admin_model_v1_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kong_admin_model_v1_plugin_proto_goTypes,
		DependencyIndexes: file_kong_admin_model_v1_plugin_proto_depIdxs,
		MessageInfos:      file_kong_admin_model_v1_plugin_proto_msgTypes,
	}.Build()
	File_kong_admin_model_v1_plugin_proto = out.File
	file_kong_admin_model_v1_plugin_proto_rawDesc = nil
	file_kong_admin_model_v1_plugin_proto_goTypes = nil
	file_kong_admin_model_v1_plugin_proto_depIdxs = nil
}
