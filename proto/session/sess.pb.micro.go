// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/session/sess.proto

package session

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Client API for SessionService service

type SessionService interface {
	Create(ctx context.Context, in *ReqSessionAdd, opts ...client.CallOption) (*ReplyInfo, error)
	CheckAvailable(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyAvailable, error)
	Remove(ctx context.Context, in *ReqSessionRemove, opts ...client.CallOption) (*ReplyInfo, error)
}

type sessionService struct {
	c    client.Client
	name string
}

func NewSessionService(name string, c client.Client) SessionService {
	return &sessionService{
		c:    c,
		name: name,
	}
}

func (c *sessionService) Create(ctx context.Context, in *ReqSessionAdd, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "SessionService.Create", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionService) CheckAvailable(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyAvailable, error) {
	req := c.c.NewRequest(c.name, "SessionService.CheckAvailable", in)
	out := new(ReplyAvailable)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionService) Remove(ctx context.Context, in *ReqSessionRemove, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "SessionService.Remove", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SessionService service

type SessionServiceHandler interface {
	Create(context.Context, *ReqSessionAdd, *ReplyInfo) error
	CheckAvailable(context.Context, *RequestInfo, *ReplyAvailable) error
	Remove(context.Context, *ReqSessionRemove, *ReplyInfo) error
}

func RegisterSessionServiceHandler(s server.Server, hdlr SessionServiceHandler, opts ...server.HandlerOption) error {
	type sessionService interface {
		Create(ctx context.Context, in *ReqSessionAdd, out *ReplyInfo) error
		CheckAvailable(ctx context.Context, in *RequestInfo, out *ReplyAvailable) error
		Remove(ctx context.Context, in *ReqSessionRemove, out *ReplyInfo) error
	}
	type SessionService struct {
		sessionService
	}
	h := &sessionServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SessionService{h}, opts...))
}

type sessionServiceHandler struct {
	SessionServiceHandler
}

func (h *sessionServiceHandler) Create(ctx context.Context, in *ReqSessionAdd, out *ReplyInfo) error {
	return h.SessionServiceHandler.Create(ctx, in, out)
}

func (h *sessionServiceHandler) CheckAvailable(ctx context.Context, in *RequestInfo, out *ReplyAvailable) error {
	return h.SessionServiceHandler.CheckAvailable(ctx, in, out)
}

func (h *sessionServiceHandler) Remove(ctx context.Context, in *ReqSessionRemove, out *ReplyInfo) error {
	return h.SessionServiceHandler.Remove(ctx, in, out)
}
