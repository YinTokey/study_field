// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package timestamppb

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

// AcgServiceClient is the client API for AcgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AcgServiceClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Random(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (*RandomResponse, error)
}

type acgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAcgServiceClient(cc grpc.ClientConnInterface) AcgServiceClient {
	return &acgServiceClient{cc}
}

func (c *acgServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/acgPackage.AcgService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *acgServiceClient) Random(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (*RandomResponse, error) {
	out := new(RandomResponse)
	err := c.cc.Invoke(ctx, "/acgPackage.AcgService/Random", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AcgServiceServer is the server API for AcgService service.
// All implementations must embed UnimplementedAcgServiceServer
// for forward compatibility
type AcgServiceServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Random(context.Context, *RandomRequest) (*RandomResponse, error)
	mustEmbedUnimplementedAcgServiceServer()
}

// UnimplementedAcgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAcgServiceServer struct {
}

func (UnimplementedAcgServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAcgServiceServer) Random(context.Context, *RandomRequest) (*RandomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Random not implemented")
}
func (UnimplementedAcgServiceServer) mustEmbedUnimplementedAcgServiceServer() {}

// UnsafeAcgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AcgServiceServer will
// result in compilation errors.
type UnsafeAcgServiceServer interface {
	mustEmbedUnimplementedAcgServiceServer()
}

func RegisterAcgServiceServer(s grpc.ServiceRegistrar, srv AcgServiceServer) {
	s.RegisterService(&AcgService_ServiceDesc, srv)
}

func _AcgService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcgServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acgPackage.AcgService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcgServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AcgService_Random_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcgServiceServer).Random(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acgPackage.AcgService/Random",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcgServiceServer).Random(ctx, req.(*RandomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AcgService_ServiceDesc is the grpc.ServiceDesc for AcgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AcgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "acgPackage.AcgService",
	HandlerType: (*AcgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _AcgService_List_Handler,
		},
		{
			MethodName: "Random",
			Handler:    _AcgService_Random_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "acg_server/acg.proto",
}