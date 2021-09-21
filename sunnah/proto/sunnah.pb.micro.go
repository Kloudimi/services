// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/sunnah.proto

package sunnah

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Sunnah service

func NewSunnahEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Sunnah service

type SunnahService interface {
	Collections(ctx context.Context, in *CollectionsRequest, opts ...client.CallOption) (*CollectionsResponse, error)
	Books(ctx context.Context, in *BooksRequest, opts ...client.CallOption) (*BooksResponse, error)
	Chapters(ctx context.Context, in *ChaptersRequest, opts ...client.CallOption) (*ChaptersResponse, error)
	Hadiths(ctx context.Context, in *HadithsRequest, opts ...client.CallOption) (*HadithsResponse, error)
}

type sunnahService struct {
	c    client.Client
	name string
}

func NewSunnahService(name string, c client.Client) SunnahService {
	return &sunnahService{
		c:    c,
		name: name,
	}
}

func (c *sunnahService) Collections(ctx context.Context, in *CollectionsRequest, opts ...client.CallOption) (*CollectionsResponse, error) {
	req := c.c.NewRequest(c.name, "Sunnah.Collections", in)
	out := new(CollectionsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sunnahService) Books(ctx context.Context, in *BooksRequest, opts ...client.CallOption) (*BooksResponse, error) {
	req := c.c.NewRequest(c.name, "Sunnah.Books", in)
	out := new(BooksResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sunnahService) Chapters(ctx context.Context, in *ChaptersRequest, opts ...client.CallOption) (*ChaptersResponse, error) {
	req := c.c.NewRequest(c.name, "Sunnah.Chapters", in)
	out := new(ChaptersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sunnahService) Hadiths(ctx context.Context, in *HadithsRequest, opts ...client.CallOption) (*HadithsResponse, error) {
	req := c.c.NewRequest(c.name, "Sunnah.Hadiths", in)
	out := new(HadithsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sunnah service

type SunnahHandler interface {
	Collections(context.Context, *CollectionsRequest, *CollectionsResponse) error
	Books(context.Context, *BooksRequest, *BooksResponse) error
	Chapters(context.Context, *ChaptersRequest, *ChaptersResponse) error
	Hadiths(context.Context, *HadithsRequest, *HadithsResponse) error
}

func RegisterSunnahHandler(s server.Server, hdlr SunnahHandler, opts ...server.HandlerOption) error {
	type sunnah interface {
		Collections(ctx context.Context, in *CollectionsRequest, out *CollectionsResponse) error
		Books(ctx context.Context, in *BooksRequest, out *BooksResponse) error
		Chapters(ctx context.Context, in *ChaptersRequest, out *ChaptersResponse) error
		Hadiths(ctx context.Context, in *HadithsRequest, out *HadithsResponse) error
	}
	type Sunnah struct {
		sunnah
	}
	h := &sunnahHandler{hdlr}
	return s.Handle(s.NewHandler(&Sunnah{h}, opts...))
}

type sunnahHandler struct {
	SunnahHandler
}

func (h *sunnahHandler) Collections(ctx context.Context, in *CollectionsRequest, out *CollectionsResponse) error {
	return h.SunnahHandler.Collections(ctx, in, out)
}

func (h *sunnahHandler) Books(ctx context.Context, in *BooksRequest, out *BooksResponse) error {
	return h.SunnahHandler.Books(ctx, in, out)
}

func (h *sunnahHandler) Chapters(ctx context.Context, in *ChaptersRequest, out *ChaptersResponse) error {
	return h.SunnahHandler.Chapters(ctx, in, out)
}

func (h *sunnahHandler) Hadiths(ctx context.Context, in *HadithsRequest, out *HadithsResponse) error {
	return h.SunnahHandler.Hadiths(ctx, in, out)
}
