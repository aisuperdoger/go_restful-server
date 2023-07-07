// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: yolov5_flame_smog.proto

package yolov5_flame_smog

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

// Yolov5FlameSmogServiceClient is the client API for Yolov5FlameSmogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Yolov5FlameSmogServiceClient interface {
	Yolov5FlameSmogImage(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Image, error)
	Yolov5FlameSmogVideo(ctx context.Context, in *VideoPath, opts ...grpc.CallOption) (*VideoPath, error)
}

type yolov5FlameSmogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYolov5FlameSmogServiceClient(cc grpc.ClientConnInterface) Yolov5FlameSmogServiceClient {
	return &yolov5FlameSmogServiceClient{cc}
}

func (c *yolov5FlameSmogServiceClient) Yolov5FlameSmogImage(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, "/yolov5_flame_smog.Yolov5FlameSmogService/yolov5FlameSmogImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yolov5FlameSmogServiceClient) Yolov5FlameSmogVideo(ctx context.Context, in *VideoPath, opts ...grpc.CallOption) (*VideoPath, error) {
	out := new(VideoPath)
	err := c.cc.Invoke(ctx, "/yolov5_flame_smog.Yolov5FlameSmogService/yolov5FlameSmogVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Yolov5FlameSmogServiceServer is the server API for Yolov5FlameSmogService service.
// All implementations must embed UnimplementedYolov5FlameSmogServiceServer
// for forward compatibility
type Yolov5FlameSmogServiceServer interface {
	Yolov5FlameSmogImage(context.Context, *Image) (*Image, error)
	Yolov5FlameSmogVideo(context.Context, *VideoPath) (*VideoPath, error)
	mustEmbedUnimplementedYolov5FlameSmogServiceServer()
}

// UnimplementedYolov5FlameSmogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYolov5FlameSmogServiceServer struct {
}

func (UnimplementedYolov5FlameSmogServiceServer) Yolov5FlameSmogImage(context.Context, *Image) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Yolov5FlameSmogImage not implemented")
}
func (UnimplementedYolov5FlameSmogServiceServer) Yolov5FlameSmogVideo(context.Context, *VideoPath) (*VideoPath, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Yolov5FlameSmogVideo not implemented")
}
func (UnimplementedYolov5FlameSmogServiceServer) mustEmbedUnimplementedYolov5FlameSmogServiceServer() {
}

// UnsafeYolov5FlameSmogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Yolov5FlameSmogServiceServer will
// result in compilation errors.
type UnsafeYolov5FlameSmogServiceServer interface {
	mustEmbedUnimplementedYolov5FlameSmogServiceServer()
}

func RegisterYolov5FlameSmogServiceServer(s grpc.ServiceRegistrar, srv Yolov5FlameSmogServiceServer) {
	s.RegisterService(&Yolov5FlameSmogService_ServiceDesc, srv)
}

func _Yolov5FlameSmogService_Yolov5FlameSmogImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Image)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Yolov5FlameSmogServiceServer).Yolov5FlameSmogImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yolov5_flame_smog.Yolov5FlameSmogService/yolov5FlameSmogImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Yolov5FlameSmogServiceServer).Yolov5FlameSmogImage(ctx, req.(*Image))
	}
	return interceptor(ctx, in, info, handler)
}

func _Yolov5FlameSmogService_Yolov5FlameSmogVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoPath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Yolov5FlameSmogServiceServer).Yolov5FlameSmogVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yolov5_flame_smog.Yolov5FlameSmogService/yolov5FlameSmogVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Yolov5FlameSmogServiceServer).Yolov5FlameSmogVideo(ctx, req.(*VideoPath))
	}
	return interceptor(ctx, in, info, handler)
}

// Yolov5FlameSmogService_ServiceDesc is the grpc.ServiceDesc for Yolov5FlameSmogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Yolov5FlameSmogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yolov5_flame_smog.Yolov5FlameSmogService",
	HandlerType: (*Yolov5FlameSmogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "yolov5FlameSmogImage",
			Handler:    _Yolov5FlameSmogService_Yolov5FlameSmogImage_Handler,
		},
		{
			MethodName: "yolov5FlameSmogVideo",
			Handler:    _Yolov5FlameSmogService_Yolov5FlameSmogVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yolov5_flame_smog.proto",
}