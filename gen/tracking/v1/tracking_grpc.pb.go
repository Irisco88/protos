// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tracking/v1/tracking.proto

package trackingv1

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

// TrackingServiceClient is the client API for TrackingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrackingServiceClient interface {
	DevicesStreams(ctx context.Context, in *DevicesStreamsRequest, opts ...grpc.CallOption) (TrackingService_DevicesStreamsClient, error)
}

type trackingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrackingServiceClient(cc grpc.ClientConnInterface) TrackingServiceClient {
	return &trackingServiceClient{cc}
}

func (c *trackingServiceClient) DevicesStreams(ctx context.Context, in *DevicesStreamsRequest, opts ...grpc.CallOption) (TrackingService_DevicesStreamsClient, error) {
	stream, err := c.cc.NewStream(ctx, &TrackingService_ServiceDesc.Streams[0], "/tracking.v1.TrackingService/DevicesStreams", opts...)
	if err != nil {
		return nil, err
	}
	x := &trackingServiceDevicesStreamsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TrackingService_DevicesStreamsClient interface {
	Recv() (*DevicesStreamsResponse, error)
	grpc.ClientStream
}

type trackingServiceDevicesStreamsClient struct {
	grpc.ClientStream
}

func (x *trackingServiceDevicesStreamsClient) Recv() (*DevicesStreamsResponse, error) {
	m := new(DevicesStreamsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TrackingServiceServer is the server API for TrackingService service.
// All implementations must embed UnimplementedTrackingServiceServer
// for forward compatibility
type TrackingServiceServer interface {
	DevicesStreams(*DevicesStreamsRequest, TrackingService_DevicesStreamsServer) error
	mustEmbedUnimplementedTrackingServiceServer()
}

// UnimplementedTrackingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrackingServiceServer struct {
}

func (UnimplementedTrackingServiceServer) DevicesStreams(*DevicesStreamsRequest, TrackingService_DevicesStreamsServer) error {
	return status.Errorf(codes.Unimplemented, "method DevicesStreams not implemented")
}
func (UnimplementedTrackingServiceServer) mustEmbedUnimplementedTrackingServiceServer() {}

// UnsafeTrackingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrackingServiceServer will
// result in compilation errors.
type UnsafeTrackingServiceServer interface {
	mustEmbedUnimplementedTrackingServiceServer()
}

func RegisterTrackingServiceServer(s grpc.ServiceRegistrar, srv TrackingServiceServer) {
	s.RegisterService(&TrackingService_ServiceDesc, srv)
}

func _TrackingService_DevicesStreams_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DevicesStreamsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrackingServiceServer).DevicesStreams(m, &trackingServiceDevicesStreamsServer{stream})
}

type TrackingService_DevicesStreamsServer interface {
	Send(*DevicesStreamsResponse) error
	grpc.ServerStream
}

type trackingServiceDevicesStreamsServer struct {
	grpc.ServerStream
}

func (x *trackingServiceDevicesStreamsServer) Send(m *DevicesStreamsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// TrackingService_ServiceDesc is the grpc.ServiceDesc for TrackingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrackingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tracking.v1.TrackingService",
	HandlerType: (*TrackingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DevicesStreams",
			Handler:       _TrackingService_DevicesStreams_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tracking/v1/tracking.proto",
}
