// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: message.proto

package proto

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

// PingProcessorClient is the client API for PingProcessor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingProcessorClient interface {
	ProcessPing(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type pingProcessorClient struct {
	cc grpc.ClientConnInterface
}

func NewPingProcessorClient(cc grpc.ClientConnInterface) PingProcessorClient {
	return &pingProcessorClient{cc}
}

func (c *pingProcessorClient) ProcessPing(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/proto.PingProcessor/ProcessPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingProcessorServer is the server API for PingProcessor service.
// All implementations must embed UnimplementedPingProcessorServer
// for forward compatibility
type PingProcessorServer interface {
	ProcessPing(context.Context, *Ping) (*Pong, error)
	mustEmbedUnimplementedPingProcessorServer()
}

// UnimplementedPingProcessorServer must be embedded to have forward compatible implementations.
type UnimplementedPingProcessorServer struct {
}

func (UnimplementedPingProcessorServer) ProcessPing(context.Context, *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessPing not implemented")
}
func (UnimplementedPingProcessorServer) mustEmbedUnimplementedPingProcessorServer() {}

// UnsafePingProcessorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingProcessorServer will
// result in compilation errors.
type UnsafePingProcessorServer interface {
	mustEmbedUnimplementedPingProcessorServer()
}

func RegisterPingProcessorServer(s grpc.ServiceRegistrar, srv PingProcessorServer) {
	s.RegisterService(&PingProcessor_ServiceDesc, srv)
}

func _PingProcessor_ProcessPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingProcessorServer).ProcessPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PingProcessor/ProcessPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingProcessorServer).ProcessPing(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

// PingProcessor_ServiceDesc is the grpc.ServiceDesc for PingProcessor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PingProcessor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PingProcessor",
	HandlerType: (*PingProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessPing",
			Handler:    _PingProcessor_ProcessPing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
