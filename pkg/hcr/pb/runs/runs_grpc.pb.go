// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: runs.proto

package runs

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

// RunsClient is the client API for Runs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RunsClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type runsClient struct {
	cc grpc.ClientConnInterface
}

func NewRunsClient(cc grpc.ClientConnInterface) RunsClient {
	return &runsClient{cc}
}

func (c *runsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/runs.Runs/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunsServer is the server API for Runs service.
// All implementations must embed UnimplementedRunsServer
// for forward compatibility
type RunsServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	mustEmbedUnimplementedRunsServer()
}

// UnimplementedRunsServer must be embedded to have forward compatible implementations.
type UnimplementedRunsServer struct {
}

func (UnimplementedRunsServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRunsServer) mustEmbedUnimplementedRunsServer() {}

// UnsafeRunsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RunsServer will
// result in compilation errors.
type UnsafeRunsServer interface {
	mustEmbedUnimplementedRunsServer()
}

func RegisterRunsServer(s grpc.ServiceRegistrar, srv RunsServer) {
	s.RegisterService(&Runs_ServiceDesc, srv)
}

func _Runs_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runs.Runs/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunsServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Runs_ServiceDesc is the grpc.ServiceDesc for Runs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Runs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "runs.Runs",
	HandlerType: (*RunsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Runs_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runs.proto",
}