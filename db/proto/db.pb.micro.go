// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/db.proto

package db

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/struct"
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

// Api Endpoints for Db service

func NewDbEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Db service

type DbService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Truncate(ctx context.Context, in *TruncateRequest, opts ...client.CallOption) (*TruncateResponse, error)
	Count(ctx context.Context, in *CountRequest, opts ...client.CallOption) (*CountResponse, error)
	RenameTable(ctx context.Context, in *RenameTableRequest, opts ...client.CallOption) (*RenameTableResponse, error)
	ListTables(ctx context.Context, in *ListTablesRequest, opts ...client.CallOption) (*ListTablesResponse, error)
	DropTable(ctx context.Context, in *DropTableRequest, opts ...client.CallOption) (*DropTableResponse, error)
}

type dbService struct {
	c    client.Client
	name string
}

func NewDbService(name string, c client.Client) DbService {
	return &dbService{
		c:    c,
		name: name,
	}
}

func (c *dbService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Db.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Db.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Db.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Db.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbService) Truncate(ctx context.Context, in *TruncateRequest, opts ...client.CallOption) (*TruncateResponse, error) {
	req := c.c.NewRequest(c.name, "Db.Truncate", in)
	out := new(TruncateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbService) Count(ctx context.Context, in *CountRequest, opts ...client.CallOption) (*CountResponse, error) {
	req := c.c.NewRequest(c.name, "Db.Count", in)
	out := new(CountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbService) RenameTable(ctx context.Context, in *RenameTableRequest, opts ...client.CallOption) (*RenameTableResponse, error) {
	req := c.c.NewRequest(c.name, "Db.RenameTable", in)
	out := new(RenameTableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbService) ListTables(ctx context.Context, in *ListTablesRequest, opts ...client.CallOption) (*ListTablesResponse, error) {
	req := c.c.NewRequest(c.name, "Db.ListTables", in)
	out := new(ListTablesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbService) DropTable(ctx context.Context, in *DropTableRequest, opts ...client.CallOption) (*DropTableResponse, error) {
	req := c.c.NewRequest(c.name, "Db.DropTable", in)
	out := new(DropTableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Db service

type DbHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Truncate(context.Context, *TruncateRequest, *TruncateResponse) error
	Count(context.Context, *CountRequest, *CountResponse) error
	RenameTable(context.Context, *RenameTableRequest, *RenameTableResponse) error
	ListTables(context.Context, *ListTablesRequest, *ListTablesResponse) error
	DropTable(context.Context, *DropTableRequest, *DropTableResponse) error
}

func RegisterDbHandler(s server.Server, hdlr DbHandler, opts ...server.HandlerOption) error {
	type db interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Truncate(ctx context.Context, in *TruncateRequest, out *TruncateResponse) error
		Count(ctx context.Context, in *CountRequest, out *CountResponse) error
		RenameTable(ctx context.Context, in *RenameTableRequest, out *RenameTableResponse) error
		ListTables(ctx context.Context, in *ListTablesRequest, out *ListTablesResponse) error
		DropTable(ctx context.Context, in *DropTableRequest, out *DropTableResponse) error
	}
	type Db struct {
		db
	}
	h := &dbHandler{hdlr}
	return s.Handle(s.NewHandler(&Db{h}, opts...))
}

type dbHandler struct {
	DbHandler
}

func (h *dbHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.DbHandler.Create(ctx, in, out)
}

func (h *dbHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.DbHandler.Read(ctx, in, out)
}

func (h *dbHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.DbHandler.Update(ctx, in, out)
}

func (h *dbHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.DbHandler.Delete(ctx, in, out)
}

func (h *dbHandler) Truncate(ctx context.Context, in *TruncateRequest, out *TruncateResponse) error {
	return h.DbHandler.Truncate(ctx, in, out)
}

func (h *dbHandler) Count(ctx context.Context, in *CountRequest, out *CountResponse) error {
	return h.DbHandler.Count(ctx, in, out)
}

func (h *dbHandler) RenameTable(ctx context.Context, in *RenameTableRequest, out *RenameTableResponse) error {
	return h.DbHandler.RenameTable(ctx, in, out)
}

func (h *dbHandler) ListTables(ctx context.Context, in *ListTablesRequest, out *ListTablesResponse) error {
	return h.DbHandler.ListTables(ctx, in, out)
}

func (h *dbHandler) DropTable(ctx context.Context, in *DropTableRequest, out *DropTableResponse) error {
	return h.DbHandler.DropTable(ctx, in, out)
}
