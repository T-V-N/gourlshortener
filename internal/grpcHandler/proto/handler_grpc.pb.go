// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/handler.proto

package handler

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
	URLShortener_GetURL_FullMethodName          = "/handler.URLShortener/GetURL"
	URLShortener_ShortenURL_FullMethodName      = "/handler.URLShortener/ShortenURL"
	URLShortener_ShortenBatchURL_FullMethodName = "/handler.URLShortener/ShortenBatchURL"
	URLShortener_Ping_FullMethodName            = "/handler.URLShortener/Ping"
	URLShortener_GetStats_FullMethodName        = "/handler.URLShortener/GetStats"
	URLShortener_GetListURL_FullMethodName      = "/handler.URLShortener/GetListURL"
	URLShortener_DeleteListURL_FullMethodName   = "/handler.URLShortener/DeleteListURL"
)

// URLShortenerClient is the client API for URLShortener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type URLShortenerClient interface {
	GetURL(ctx context.Context, in *GetURLRequest, opts ...grpc.CallOption) (*GetURLResponse, error)
	ShortenURL(ctx context.Context, in *ShortenURLRequest, opts ...grpc.CallOption) (*ShortenURLResponse, error)
	ShortenBatchURL(ctx context.Context, in *ShortenBatchURLRequest, opts ...grpc.CallOption) (*ShortenBatchURLResponse, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResponse, error)
	GetStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetStatsResponse, error)
	GetListURL(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetListURLResponse, error)
	DeleteListURL(ctx context.Context, in *DeleteListURLRequest, opts ...grpc.CallOption) (*Empty, error)
}

type uRLShortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewURLShortenerClient(cc grpc.ClientConnInterface) URLShortenerClient {
	return &uRLShortenerClient{cc}
}

func (c *uRLShortenerClient) GetURL(ctx context.Context, in *GetURLRequest, opts ...grpc.CallOption) (*GetURLResponse, error) {
	out := new(GetURLResponse)
	err := c.cc.Invoke(ctx, URLShortener_GetURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLShortenerClient) ShortenURL(ctx context.Context, in *ShortenURLRequest, opts ...grpc.CallOption) (*ShortenURLResponse, error) {
	out := new(ShortenURLResponse)
	err := c.cc.Invoke(ctx, URLShortener_ShortenURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLShortenerClient) ShortenBatchURL(ctx context.Context, in *ShortenBatchURLRequest, opts ...grpc.CallOption) (*ShortenBatchURLResponse, error) {
	out := new(ShortenBatchURLResponse)
	err := c.cc.Invoke(ctx, URLShortener_ShortenBatchURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLShortenerClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, URLShortener_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLShortenerClient) GetStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, URLShortener_GetStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLShortenerClient) GetListURL(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetListURLResponse, error) {
	out := new(GetListURLResponse)
	err := c.cc.Invoke(ctx, URLShortener_GetListURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLShortenerClient) DeleteListURL(ctx context.Context, in *DeleteListURLRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, URLShortener_DeleteListURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// URLShortenerServer is the server API for URLShortener service.
// All implementations must embed UnimplementedURLShortenerServer
// for forward compatibility
type URLShortenerServer interface {
	GetURL(context.Context, *GetURLRequest) (*GetURLResponse, error)
	ShortenURL(context.Context, *ShortenURLRequest) (*ShortenURLResponse, error)
	ShortenBatchURL(context.Context, *ShortenBatchURLRequest) (*ShortenBatchURLResponse, error)
	Ping(context.Context, *Empty) (*PingResponse, error)
	GetStats(context.Context, *Empty) (*GetStatsResponse, error)
	GetListURL(context.Context, *Empty) (*GetListURLResponse, error)
	DeleteListURL(context.Context, *DeleteListURLRequest) (*Empty, error)
	mustEmbedUnimplementedURLShortenerServer()
}

// UnimplementedURLShortenerServer must be embedded to have forward compatible implementations.
type UnimplementedURLShortenerServer struct {
}

func (UnimplementedURLShortenerServer) GetURL(context.Context, *GetURLRequest) (*GetURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetURL not implemented")
}
func (UnimplementedURLShortenerServer) ShortenURL(context.Context, *ShortenURLRequest) (*ShortenURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortenURL not implemented")
}
func (UnimplementedURLShortenerServer) ShortenBatchURL(context.Context, *ShortenBatchURLRequest) (*ShortenBatchURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortenBatchURL not implemented")
}
func (UnimplementedURLShortenerServer) Ping(context.Context, *Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedURLShortenerServer) GetStats(context.Context, *Empty) (*GetStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedURLShortenerServer) GetListURL(context.Context, *Empty) (*GetListURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListURL not implemented")
}
func (UnimplementedURLShortenerServer) DeleteListURL(context.Context, *DeleteListURLRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteListURL not implemented")
}
func (UnimplementedURLShortenerServer) mustEmbedUnimplementedURLShortenerServer() {}

// UnsafeURLShortenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to URLShortenerServer will
// result in compilation errors.
type UnsafeURLShortenerServer interface {
	mustEmbedUnimplementedURLShortenerServer()
}

func RegisterURLShortenerServer(s grpc.ServiceRegistrar, srv URLShortenerServer) {
	s.RegisterService(&URLShortener_ServiceDesc, srv)
}

func _URLShortener_GetURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShortenerServer).GetURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: URLShortener_GetURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLShortenerServer).GetURL(ctx, req.(*GetURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLShortener_ShortenURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShortenerServer).ShortenURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: URLShortener_ShortenURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLShortenerServer).ShortenURL(ctx, req.(*ShortenURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLShortener_ShortenBatchURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenBatchURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShortenerServer).ShortenBatchURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: URLShortener_ShortenBatchURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLShortenerServer).ShortenBatchURL(ctx, req.(*ShortenBatchURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLShortener_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShortenerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: URLShortener_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLShortenerServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLShortener_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShortenerServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: URLShortener_GetStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLShortenerServer).GetStats(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLShortener_GetListURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShortenerServer).GetListURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: URLShortener_GetListURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLShortenerServer).GetListURL(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLShortener_DeleteListURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteListURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShortenerServer).DeleteListURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: URLShortener_DeleteListURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLShortenerServer).DeleteListURL(ctx, req.(*DeleteListURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// URLShortener_ServiceDesc is the grpc.ServiceDesc for URLShortener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var URLShortener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "handler.URLShortener",
	HandlerType: (*URLShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetURL",
			Handler:    _URLShortener_GetURL_Handler,
		},
		{
			MethodName: "ShortenURL",
			Handler:    _URLShortener_ShortenURL_Handler,
		},
		{
			MethodName: "ShortenBatchURL",
			Handler:    _URLShortener_ShortenBatchURL_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _URLShortener_Ping_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _URLShortener_GetStats_Handler,
		},
		{
			MethodName: "GetListURL",
			Handler:    _URLShortener_GetListURL_Handler,
		},
		{
			MethodName: "DeleteListURL",
			Handler:    _URLShortener_DeleteListURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/handler.proto",
}