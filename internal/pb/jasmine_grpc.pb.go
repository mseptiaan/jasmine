// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: jasmine.proto

package pb

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
	JasmineEndpoint_GetData_FullMethodName        = "/JasmineEndpoint/GetData"
	JasmineEndpoint_GetDataGeoJson_FullMethodName = "/JasmineEndpoint/GetDataGeoJson"
	JasmineEndpoint_PostNearby_FullMethodName     = "/JasmineEndpoint/PostNearby"
	JasmineEndpoint_PostStore_FullMethodName      = "/JasmineEndpoint/PostStore"
)

// JasmineEndpointClient is the client API for JasmineEndpoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JasmineEndpointClient interface {
	GetData(ctx context.Context, in *RequestGet, opts ...grpc.CallOption) (*ResponseGet, error)
	GetDataGeoJson(ctx context.Context, in *RequestGet, opts ...grpc.CallOption) (*ResponseGetGeoJson, error)
	PostNearby(ctx context.Context, in *RequestNearby, opts ...grpc.CallOption) (*ResponseNearby, error)
	PostStore(ctx context.Context, in *RequestStore, opts ...grpc.CallOption) (*ResponseStore, error)
}

type jasmineEndpointClient struct {
	cc grpc.ClientConnInterface
}

func NewJasmineEndpointClient(cc grpc.ClientConnInterface) JasmineEndpointClient {
	return &jasmineEndpointClient{cc}
}

func (c *jasmineEndpointClient) GetData(ctx context.Context, in *RequestGet, opts ...grpc.CallOption) (*ResponseGet, error) {
	out := new(ResponseGet)
	err := c.cc.Invoke(ctx, JasmineEndpoint_GetData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jasmineEndpointClient) GetDataGeoJson(ctx context.Context, in *RequestGet, opts ...grpc.CallOption) (*ResponseGetGeoJson, error) {
	out := new(ResponseGetGeoJson)
	err := c.cc.Invoke(ctx, JasmineEndpoint_GetDataGeoJson_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jasmineEndpointClient) PostNearby(ctx context.Context, in *RequestNearby, opts ...grpc.CallOption) (*ResponseNearby, error) {
	out := new(ResponseNearby)
	err := c.cc.Invoke(ctx, JasmineEndpoint_PostNearby_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jasmineEndpointClient) PostStore(ctx context.Context, in *RequestStore, opts ...grpc.CallOption) (*ResponseStore, error) {
	out := new(ResponseStore)
	err := c.cc.Invoke(ctx, JasmineEndpoint_PostStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JasmineEndpointServer is the server API for JasmineEndpoint service.
// All implementations must embed UnimplementedJasmineEndpointServer
// for forward compatibility
type JasmineEndpointServer interface {
	GetData(context.Context, *RequestGet) (*ResponseGet, error)
	GetDataGeoJson(context.Context, *RequestGet) (*ResponseGetGeoJson, error)
	PostNearby(context.Context, *RequestNearby) (*ResponseNearby, error)
	PostStore(context.Context, *RequestStore) (*ResponseStore, error)
	mustEmbedUnimplementedJasmineEndpointServer()
}

// UnimplementedJasmineEndpointServer must be embedded to have forward compatible implementations.
type UnimplementedJasmineEndpointServer struct {
}

func (UnimplementedJasmineEndpointServer) GetData(context.Context, *RequestGet) (*ResponseGet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedJasmineEndpointServer) GetDataGeoJson(context.Context, *RequestGet) (*ResponseGetGeoJson, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataGeoJson not implemented")
}
func (UnimplementedJasmineEndpointServer) PostNearby(context.Context, *RequestNearby) (*ResponseNearby, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostNearby not implemented")
}
func (UnimplementedJasmineEndpointServer) PostStore(context.Context, *RequestStore) (*ResponseStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostStore not implemented")
}
func (UnimplementedJasmineEndpointServer) mustEmbedUnimplementedJasmineEndpointServer() {}

// UnsafeJasmineEndpointServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JasmineEndpointServer will
// result in compilation errors.
type UnsafeJasmineEndpointServer interface {
	mustEmbedUnimplementedJasmineEndpointServer()
}

func RegisterJasmineEndpointServer(s grpc.ServiceRegistrar, srv JasmineEndpointServer) {
	s.RegisterService(&JasmineEndpoint_ServiceDesc, srv)
}

func _JasmineEndpoint_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JasmineEndpointServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JasmineEndpoint_GetData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JasmineEndpointServer).GetData(ctx, req.(*RequestGet))
	}
	return interceptor(ctx, in, info, handler)
}

func _JasmineEndpoint_GetDataGeoJson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JasmineEndpointServer).GetDataGeoJson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JasmineEndpoint_GetDataGeoJson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JasmineEndpointServer).GetDataGeoJson(ctx, req.(*RequestGet))
	}
	return interceptor(ctx, in, info, handler)
}

func _JasmineEndpoint_PostNearby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestNearby)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JasmineEndpointServer).PostNearby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JasmineEndpoint_PostNearby_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JasmineEndpointServer).PostNearby(ctx, req.(*RequestNearby))
	}
	return interceptor(ctx, in, info, handler)
}

func _JasmineEndpoint_PostStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JasmineEndpointServer).PostStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JasmineEndpoint_PostStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JasmineEndpointServer).PostStore(ctx, req.(*RequestStore))
	}
	return interceptor(ctx, in, info, handler)
}

// JasmineEndpoint_ServiceDesc is the grpc.ServiceDesc for JasmineEndpoint service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JasmineEndpoint_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "JasmineEndpoint",
	HandlerType: (*JasmineEndpointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _JasmineEndpoint_GetData_Handler,
		},
		{
			MethodName: "GetDataGeoJson",
			Handler:    _JasmineEndpoint_GetDataGeoJson_Handler,
		},
		{
			MethodName: "PostNearby",
			Handler:    _JasmineEndpoint_PostNearby_Handler,
		},
		{
			MethodName: "PostStore",
			Handler:    _JasmineEndpoint_PostStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jasmine.proto",
}
