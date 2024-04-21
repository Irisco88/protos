// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: imeis/v1/imeis.proto

package imeisv1

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

// ImeisServiceClient is the client API for ImeisService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImeisServiceClient interface {
	CreateImeis(ctx context.Context, in *CreateImeisRequest, opts ...grpc.CallOption) (*CreateImeisResponse, error)
	UpdateImeis(ctx context.Context, in *UpdateImeisRequest, opts ...grpc.CallOption) (*UpdateImeisResponse, error)
	DeleteImeis(ctx context.Context, in *DeleteImeisRequest, opts ...grpc.CallOption) (*DeleteImeisResponse, error)
	ListImeiss(ctx context.Context, in *ListImeissRequest, opts ...grpc.CallOption) (*ListImeissResponse, error)
}

type imeisServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImeisServiceClient(cc grpc.ClientConnInterface) ImeisServiceClient {
	return &imeisServiceClient{cc}
}

func (c *imeisServiceClient) CreateImeis(ctx context.Context, in *CreateImeisRequest, opts ...grpc.CallOption) (*CreateImeisResponse, error) {
	out := new(CreateImeisResponse)
	err := c.cc.Invoke(ctx, "/imeis.v1.ImeisService/CreateImeis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imeisServiceClient) UpdateImeis(ctx context.Context, in *UpdateImeisRequest, opts ...grpc.CallOption) (*UpdateImeisResponse, error) {
	out := new(UpdateImeisResponse)
	err := c.cc.Invoke(ctx, "/imeis.v1.ImeisService/UpdateImeis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imeisServiceClient) DeleteImeis(ctx context.Context, in *DeleteImeisRequest, opts ...grpc.CallOption) (*DeleteImeisResponse, error) {
	out := new(DeleteImeisResponse)
	err := c.cc.Invoke(ctx, "/imeis.v1.ImeisService/DeleteImeis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imeisServiceClient) ListImeiss(ctx context.Context, in *ListImeissRequest, opts ...grpc.CallOption) (*ListImeissResponse, error) {
	out := new(ListImeissResponse)
	err := c.cc.Invoke(ctx, "/imeis.v1.ImeisService/ListImeiss", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImeisServiceServer is the server API for ImeisService service.
// All implementations must embed UnimplementedImeisServiceServer
// for forward compatibility
type ImeisServiceServer interface {
	CreateImeis(context.Context, *CreateImeisRequest) (*CreateImeisResponse, error)
	UpdateImeis(context.Context, *UpdateImeisRequest) (*UpdateImeisResponse, error)
	DeleteImeis(context.Context, *DeleteImeisRequest) (*DeleteImeisResponse, error)
	ListImeiss(context.Context, *ListImeissRequest) (*ListImeissResponse, error)
	mustEmbedUnimplementedImeisServiceServer()
}

// UnimplementedImeisServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImeisServiceServer struct {
}

func (UnimplementedImeisServiceServer) CreateImeis(context.Context, *CreateImeisRequest) (*CreateImeisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImeis not implemented")
}
func (UnimplementedImeisServiceServer) UpdateImeis(context.Context, *UpdateImeisRequest) (*UpdateImeisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateImeis not implemented")
}
func (UnimplementedImeisServiceServer) DeleteImeis(context.Context, *DeleteImeisRequest) (*DeleteImeisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImeis not implemented")
}
func (UnimplementedImeisServiceServer) ListImeiss(context.Context, *ListImeissRequest) (*ListImeissResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImeiss not implemented")
}
func (UnimplementedImeisServiceServer) mustEmbedUnimplementedImeisServiceServer() {}

// UnsafeImeisServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImeisServiceServer will
// result in compilation errors.
type UnsafeImeisServiceServer interface {
	mustEmbedUnimplementedImeisServiceServer()
}

func RegisterImeisServiceServer(s grpc.ServiceRegistrar, srv ImeisServiceServer) {
	s.RegisterService(&ImeisService_ServiceDesc, srv)
}

func _ImeisService_CreateImeis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImeisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImeisServiceServer).CreateImeis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imeis.v1.ImeisService/CreateImeis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImeisServiceServer).CreateImeis(ctx, req.(*CreateImeisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImeisService_UpdateImeis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImeisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImeisServiceServer).UpdateImeis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imeis.v1.ImeisService/UpdateImeis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImeisServiceServer).UpdateImeis(ctx, req.(*UpdateImeisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImeisService_DeleteImeis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImeisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImeisServiceServer).DeleteImeis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imeis.v1.ImeisService/DeleteImeis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImeisServiceServer).DeleteImeis(ctx, req.(*DeleteImeisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImeisService_ListImeiss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImeissRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImeisServiceServer).ListImeiss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imeis.v1.ImeisService/ListImeiss",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImeisServiceServer).ListImeiss(ctx, req.(*ListImeissRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImeisService_ServiceDesc is the grpc.ServiceDesc for ImeisService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImeisService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "imeis.v1.ImeisService",
	HandlerType: (*ImeisServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateImeis",
			Handler:    _ImeisService_CreateImeis_Handler,
		},
		{
			MethodName: "UpdateImeis",
			Handler:    _ImeisService_UpdateImeis_Handler,
		},
		{
			MethodName: "DeleteImeis",
			Handler:    _ImeisService_DeleteImeis_Handler,
		},
		{
			MethodName: "ListImeiss",
			Handler:    _ImeisService_ListImeiss_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "imeis/v1/imeis.proto",
}
