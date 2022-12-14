// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: kong/admin/service/v1/service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceServiceClient is the client API for ServiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceServiceClient interface {
	GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error)
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error)
	UpsertService(ctx context.Context, in *UpsertServiceRequest, opts ...grpc.CallOption) (*UpsertServiceResponse, error)
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error)
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
}

type serviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceServiceClient(cc grpc.ClientConnInterface) ServiceServiceClient {
	return &serviceServiceClient{cc}
}

func (c *serviceServiceClient) GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error) {
	out := new(GetServiceResponse)
	err := c.cc.Invoke(ctx, "/kong.admin.service.v1.ServiceService/GetService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error) {
	out := new(CreateServiceResponse)
	err := c.cc.Invoke(ctx, "/kong.admin.service.v1.ServiceService/CreateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) UpsertService(ctx context.Context, in *UpsertServiceRequest, opts ...grpc.CallOption) (*UpsertServiceResponse, error) {
	out := new(UpsertServiceResponse)
	err := c.cc.Invoke(ctx, "/kong.admin.service.v1.ServiceService/UpsertService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error) {
	out := new(DeleteServiceResponse)
	err := c.cc.Invoke(ctx, "/kong.admin.service.v1.ServiceService/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, "/kong.admin.service.v1.ServiceService/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServiceServer is the server API for ServiceService service.
// All implementations must embed UnimplementedServiceServiceServer
// for forward compatibility
type ServiceServiceServer interface {
	GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error)
	CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error)
	UpsertService(context.Context, *UpsertServiceRequest) (*UpsertServiceResponse, error)
	DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error)
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	mustEmbedUnimplementedServiceServiceServer()
}

// UnimplementedServiceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServiceServer struct {
}

func (UnimplementedServiceServiceServer) GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}
func (UnimplementedServiceServiceServer) CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedServiceServiceServer) UpsertService(context.Context, *UpsertServiceRequest) (*UpsertServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertService not implemented")
}
func (UnimplementedServiceServiceServer) DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedServiceServiceServer) ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedServiceServiceServer) mustEmbedUnimplementedServiceServiceServer() {}

// UnsafeServiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServiceServer will
// result in compilation errors.
type UnsafeServiceServiceServer interface {
	mustEmbedUnimplementedServiceServiceServer()
}

func RegisterServiceServiceServer(s grpc.ServiceRegistrar, srv ServiceServiceServer) {
	s.RegisterService(&ServiceService_ServiceDesc, srv)
}

func _ServiceService_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kong.admin.service.v1.ServiceService/GetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).GetService(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kong.admin.service.v1.ServiceService/CreateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).CreateService(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_UpsertService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).UpsertService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kong.admin.service.v1.ServiceService/UpsertService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).UpsertService(ctx, req.(*UpsertServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kong.admin.service.v1.ServiceService/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kong.admin.service.v1.ServiceService/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceService_ServiceDesc is the grpc.ServiceDesc for ServiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kong.admin.service.v1.ServiceService",
	HandlerType: (*ServiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetService",
			Handler:    _ServiceService_GetService_Handler,
		},
		{
			MethodName: "CreateService",
			Handler:    _ServiceService_CreateService_Handler,
		},
		{
			MethodName: "UpsertService",
			Handler:    _ServiceService_UpsertService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _ServiceService_DeleteService_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _ServiceService_ListServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kong/admin/service/v1/service.proto",
}