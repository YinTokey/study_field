// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: unsp.proto

package unsplashPackage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UnPictureService service

type UnPictureService interface {
	GetUnPictureInfo(ctx context.Context, in *UnPictureRequest, opts ...client.CallOption) (*UnPictureInfo, error)
}

type unPictureService struct {
	c    client.Client
	name string
}

func NewUnPictureService(name string, c client.Client) UnPictureService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "unsplashPackage"
	}
	return &unPictureService{
		c:    c,
		name: name,
	}
}

func (c *unPictureService) GetUnPictureInfo(ctx context.Context, in *UnPictureRequest, opts ...client.CallOption) (*UnPictureInfo, error) {
	req := c.c.NewRequest(c.name, "UnPictureService.GetUnPictureInfo", in)
	out := new(UnPictureInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UnPictureService service

type UnPictureServiceHandler interface {
	GetUnPictureInfo(context.Context, *UnPictureRequest, *UnPictureInfo) error
}

func RegisterUnPictureServiceHandler(s server.Server, hdlr UnPictureServiceHandler, opts ...server.HandlerOption) error {
	type unPictureService interface {
		GetUnPictureInfo(ctx context.Context, in *UnPictureRequest, out *UnPictureInfo) error
	}
	type UnPictureService struct {
		unPictureService
	}
	h := &unPictureServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UnPictureService{h}, opts...))
}

type unPictureServiceHandler struct {
	UnPictureServiceHandler
}

func (h *unPictureServiceHandler) GetUnPictureInfo(ctx context.Context, in *UnPictureRequest, out *UnPictureInfo) error {
	return h.UnPictureServiceHandler.GetUnPictureInfo(ctx, in, out)
}