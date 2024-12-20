// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.6
// source: api/proto/api.proto

package postgres

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PostgresService_CreatePostgres_FullMethodName = "/proto.PostgresService/CreatePostgres"
	PostgresService_UpdatePostgres_FullMethodName = "/proto.PostgresService/UpdatePostgres"
	PostgresService_DeletePostgres_FullMethodName = "/proto.PostgresService/DeletePostgres"
)

// PostgresServiceClient is the client API for PostgresService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostgresServiceClient interface {
	CreatePostgres(ctx context.Context, in *CreatePostgresRequest, opts ...grpc.CallOption) (*CreatePostgresResponse, error)
	UpdatePostgres(ctx context.Context, in *UpdatePostgresRequest, opts ...grpc.CallOption) (*UpdatePostgresResponse, error)
	DeletePostgres(ctx context.Context, in *DeletePostgresRequest, opts ...grpc.CallOption) (*DeletePostgresResponse, error)
}

type postgresServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostgresServiceClient(cc grpc.ClientConnInterface) PostgresServiceClient {
	return &postgresServiceClient{cc}
}

func (c *postgresServiceClient) CreatePostgres(ctx context.Context, in *CreatePostgresRequest, opts ...grpc.CallOption) (*CreatePostgresResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePostgresResponse)
	err := c.cc.Invoke(ctx, PostgresService_CreatePostgres_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresServiceClient) UpdatePostgres(ctx context.Context, in *UpdatePostgresRequest, opts ...grpc.CallOption) (*UpdatePostgresResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePostgresResponse)
	err := c.cc.Invoke(ctx, PostgresService_UpdatePostgres_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresServiceClient) DeletePostgres(ctx context.Context, in *DeletePostgresRequest, opts ...grpc.CallOption) (*DeletePostgresResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePostgresResponse)
	err := c.cc.Invoke(ctx, PostgresService_DeletePostgres_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostgresServiceServer is the server API for PostgresService service.
// All implementations must embed UnimplementedPostgresServiceServer
// for forward compatibility.
type PostgresServiceServer interface {
	CreatePostgres(context.Context, *CreatePostgresRequest) (*CreatePostgresResponse, error)
	UpdatePostgres(context.Context, *UpdatePostgresRequest) (*UpdatePostgresResponse, error)
	DeletePostgres(context.Context, *DeletePostgresRequest) (*DeletePostgresResponse, error)
	mustEmbedUnimplementedPostgresServiceServer()
}

// UnimplementedPostgresServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPostgresServiceServer struct{}

func (UnimplementedPostgresServiceServer) CreatePostgres(context.Context, *CreatePostgresRequest) (*CreatePostgresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePostgres not implemented")
}
func (UnimplementedPostgresServiceServer) UpdatePostgres(context.Context, *UpdatePostgresRequest) (*UpdatePostgresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePostgres not implemented")
}
func (UnimplementedPostgresServiceServer) DeletePostgres(context.Context, *DeletePostgresRequest) (*DeletePostgresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePostgres not implemented")
}
func (UnimplementedPostgresServiceServer) mustEmbedUnimplementedPostgresServiceServer() {}
func (UnimplementedPostgresServiceServer) testEmbeddedByValue()                         {}

// UnsafePostgresServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostgresServiceServer will
// result in compilation errors.
type UnsafePostgresServiceServer interface {
	mustEmbedUnimplementedPostgresServiceServer()
}

func RegisterPostgresServiceServer(s grpc.ServiceRegistrar, srv PostgresServiceServer) {
	// If the following call pancis, it indicates UnimplementedPostgresServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PostgresService_ServiceDesc, srv)
}

func _PostgresService_CreatePostgres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostgresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresServiceServer).CreatePostgres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresService_CreatePostgres_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresServiceServer).CreatePostgres(ctx, req.(*CreatePostgresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresService_UpdatePostgres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostgresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresServiceServer).UpdatePostgres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresService_UpdatePostgres_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresServiceServer).UpdatePostgres(ctx, req.(*UpdatePostgresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresService_DeletePostgres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostgresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresServiceServer).DeletePostgres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresService_DeletePostgres_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresServiceServer).DeletePostgres(ctx, req.(*DeletePostgresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostgresService_ServiceDesc is the grpc.ServiceDesc for PostgresService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostgresService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PostgresService",
	HandlerType: (*PostgresServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePostgres",
			Handler:    _PostgresService_CreatePostgres_Handler,
		},
		{
			MethodName: "UpdatePostgres",
			Handler:    _PostgresService_UpdatePostgres_Handler,
		},
		{
			MethodName: "DeletePostgres",
			Handler:    _PostgresService_DeletePostgres_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/api.proto",
}
