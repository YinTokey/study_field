// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package unsplashPackage

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

// UnPictureServiceClient is the client API for UnPictureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UnPictureServiceClient interface {
	GetUnPictureInfo(ctx context.Context, in *UnPictureRequest, opts ...grpc.CallOption) (*UnPictureInfo, error)
}

type unPictureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUnPictureServiceClient(cc grpc.ClientConnInterface) UnPictureServiceClient {
	return &unPictureServiceClient{cc}
}

func (c *unPictureServiceClient) GetUnPictureInfo(ctx context.Context, in *UnPictureRequest, opts ...grpc.CallOption) (*UnPictureInfo, error) {
	out := new(UnPictureInfo)
	err := c.cc.Invoke(ctx, "/unsplashPackage.UnPictureService/GetUnPictureInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnPictureServiceServer is the server API for UnPictureService service.
// All implementations must embed UnimplementedUnPictureServiceServer
// for forward compatibility
type UnPictureServiceServer interface {
	GetUnPictureInfo(context.Context, *UnPictureRequest) (*UnPictureInfo, error)
	mustEmbedUnimplementedUnPictureServiceServer()
}

// UnimplementedUnPictureServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUnPictureServiceServer struct {
}

func (UnimplementedUnPictureServiceServer) GetUnPictureInfo(context.Context, *UnPictureRequest) (*UnPictureInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnPictureInfo not implemented")
}
func (UnimplementedUnPictureServiceServer) mustEmbedUnimplementedUnPictureServiceServer() {}

// UnsafeUnPictureServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UnPictureServiceServer will
// result in compilation errors.
type UnsafeUnPictureServiceServer interface {
	mustEmbedUnimplementedUnPictureServiceServer()
}

func RegisterUnPictureServiceServer(s grpc.ServiceRegistrar, srv UnPictureServiceServer) {
	s.RegisterService(&UnPictureService_ServiceDesc, srv)
}

func _UnPictureService_GetUnPictureInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnPictureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnPictureServiceServer).GetUnPictureInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/unsplashPackage.UnPictureService/GetUnPictureInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnPictureServiceServer).GetUnPictureInfo(ctx, req.(*UnPictureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UnPictureService_ServiceDesc is the grpc.ServiceDesc for UnPictureService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UnPictureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "unsplashPackage.UnPictureService",
	HandlerType: (*UnPictureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUnPictureInfo",
			Handler:    _UnPictureService_GetUnPictureInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "unsp.proto",
}