// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: ioels/v1/ioels.proto

package ioelsv1

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

const (
	IoelsService_CreateIoels_FullMethodName = "/ioels.v1.IoelsService/CreateIoels"
	IoelsService_UpdateIoels_FullMethodName = "/ioels.v1.IoelsService/UpdateIoels"
	IoelsService_DeleteIoels_FullMethodName = "/ioels.v1.IoelsService/DeleteIoels"
	IoelsService_ListIoelss_FullMethodName  = "/ioels.v1.IoelsService/ListIoelss"
)

// IoelsServiceClient is the client API for IoelsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IoelsServiceClient interface {
	CreateIoels(ctx context.Context, in *CreateIoelsRequest, opts ...grpc.CallOption) (*CreateIoelsResponse, error)
	UpdateIoels(ctx context.Context, in *UpdateIoelsRequest, opts ...grpc.CallOption) (*UpdateIoelsResponse, error)
	DeleteIoels(ctx context.Context, in *DeleteIoelsRequest, opts ...grpc.CallOption) (*DeleteIoelsResponse, error)
	ListIoelss(ctx context.Context, in *ListIoelssRequest, opts ...grpc.CallOption) (*ListIoelssResponse, error)
}

type ioelsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIoelsServiceClient(cc grpc.ClientConnInterface) IoelsServiceClient {
	return &ioelsServiceClient{cc}
}

func (c *ioelsServiceClient) CreateIoels(ctx context.Context, in *CreateIoelsRequest, opts ...grpc.CallOption) (*CreateIoelsResponse, error) {
	out := new(CreateIoelsResponse)
	err := c.cc.Invoke(ctx, IoelsService_CreateIoels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioelsServiceClient) UpdateIoels(ctx context.Context, in *UpdateIoelsRequest, opts ...grpc.CallOption) (*UpdateIoelsResponse, error) {
	out := new(UpdateIoelsResponse)
	err := c.cc.Invoke(ctx, IoelsService_UpdateIoels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioelsServiceClient) DeleteIoels(ctx context.Context, in *DeleteIoelsRequest, opts ...grpc.CallOption) (*DeleteIoelsResponse, error) {
	out := new(DeleteIoelsResponse)
	err := c.cc.Invoke(ctx, IoelsService_DeleteIoels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioelsServiceClient) ListIoelss(ctx context.Context, in *ListIoelssRequest, opts ...grpc.CallOption) (*ListIoelssResponse, error) {
	out := new(ListIoelssResponse)
	err := c.cc.Invoke(ctx, IoelsService_ListIoelss_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IoelsServiceServer is the server API for IoelsService service.
// All implementations must embed UnimplementedIoelsServiceServer
// for forward compatibility
type IoelsServiceServer interface {
	CreateIoels(context.Context, *CreateIoelsRequest) (*CreateIoelsResponse, error)
	UpdateIoels(context.Context, *UpdateIoelsRequest) (*UpdateIoelsResponse, error)
	DeleteIoels(context.Context, *DeleteIoelsRequest) (*DeleteIoelsResponse, error)
	ListIoelss(context.Context, *ListIoelssRequest) (*ListIoelssResponse, error)
	mustEmbedUnimplementedIoelsServiceServer()
}

// UnimplementedIoelsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIoelsServiceServer struct {
}

func (UnimplementedIoelsServiceServer) CreateIoels(context.Context, *CreateIoelsRequest) (*CreateIoelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIoels not implemented")
}
func (UnimplementedIoelsServiceServer) UpdateIoels(context.Context, *UpdateIoelsRequest) (*UpdateIoelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIoels not implemented")
}
func (UnimplementedIoelsServiceServer) DeleteIoels(context.Context, *DeleteIoelsRequest) (*DeleteIoelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIoels not implemented")
}
func (UnimplementedIoelsServiceServer) ListIoelss(context.Context, *ListIoelssRequest) (*ListIoelssResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIoelss not implemented")
}
func (UnimplementedIoelsServiceServer) mustEmbedUnimplementedIoelsServiceServer() {}

// UnsafeIoelsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IoelsServiceServer will
// result in compilation errors.
type UnsafeIoelsServiceServer interface {
	mustEmbedUnimplementedIoelsServiceServer()
}

func RegisterIoelsServiceServer(s grpc.ServiceRegistrar, srv IoelsServiceServer) {
	s.RegisterService(&IoelsService_ServiceDesc, srv)
}

func _IoelsService_CreateIoels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIoelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoelsServiceServer).CreateIoels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IoelsService_CreateIoels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoelsServiceServer).CreateIoels(ctx, req.(*CreateIoelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoelsService_UpdateIoels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIoelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoelsServiceServer).UpdateIoels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IoelsService_UpdateIoels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoelsServiceServer).UpdateIoels(ctx, req.(*UpdateIoelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoelsService_DeleteIoels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIoelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoelsServiceServer).DeleteIoels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IoelsService_DeleteIoels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoelsServiceServer).DeleteIoels(ctx, req.(*DeleteIoelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoelsService_ListIoelss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIoelssRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoelsServiceServer).ListIoelss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IoelsService_ListIoelss_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoelsServiceServer).ListIoelss(ctx, req.(*ListIoelssRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IoelsService_ServiceDesc is the grpc.ServiceDesc for IoelsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IoelsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ioels.v1.IoelsService",
	HandlerType: (*IoelsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIoels",
			Handler:    _IoelsService_CreateIoels_Handler,
		},
		{
			MethodName: "UpdateIoels",
			Handler:    _IoelsService_UpdateIoels_Handler,
		},
		{
			MethodName: "DeleteIoels",
			Handler:    _IoelsService_DeleteIoels_Handler,
		},
		{
			MethodName: "ListIoelss",
			Handler:    _IoelsService_ListIoelss_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ioels/v1/ioels.proto",
}