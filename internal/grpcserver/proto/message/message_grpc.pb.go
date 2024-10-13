// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: internal/grpcserver/proto/message/message.proto

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

// HandleMessageClient is the client API for HandleMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HandleMessageClient interface {
	StoreMessage(ctx context.Context, in *MessageInfo, opts ...grpc.CallOption) (*Response, error)
}

type handleMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewHandleMessageClient(cc grpc.ClientConnInterface) HandleMessageClient {
	return &handleMessageClient{cc}
}

func (c *handleMessageClient) StoreMessage(ctx context.Context, in *MessageInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/message.HandleMessage/StoreMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandleMessageServer is the server API for HandleMessage service.
// All implementations must embed UnimplementedHandleMessageServer
// for forward compatibility
type HandleMessageServer interface {
	StoreMessage(context.Context, *MessageInfo) (*Response, error)
	mustEmbedUnimplementedHandleMessageServer()
}

// UnimplementedHandleMessageServer must be embedded to have forward compatible implementations.
type UnimplementedHandleMessageServer struct {
}

func (UnimplementedHandleMessageServer) StoreMessage(context.Context, *MessageInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreMessage not implemented")
}
func (UnimplementedHandleMessageServer) mustEmbedUnimplementedHandleMessageServer() {}

// UnsafeHandleMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HandleMessageServer will
// result in compilation errors.
type UnsafeHandleMessageServer interface {
	mustEmbedUnimplementedHandleMessageServer()
}

func RegisterHandleMessageServer(s grpc.ServiceRegistrar, srv HandleMessageServer) {
	s.RegisterService(&HandleMessage_ServiceDesc, srv)
}

func _HandleMessage_StoreMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandleMessageServer).StoreMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.HandleMessage/StoreMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandleMessageServer).StoreMessage(ctx, req.(*MessageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// HandleMessage_ServiceDesc is the grpc.ServiceDesc for HandleMessage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HandleMessage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.HandleMessage",
	HandlerType: (*HandleMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreMessage",
			Handler:    _HandleMessage_StoreMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/grpcserver/proto/message/message.proto",
}
