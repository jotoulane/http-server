// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.3
// - protoc             v5.29.2
// source: api/ping/ping.proto

package ping

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPingGetPing = "/api.ping.Ping/GetPing"

type PingHTTPServer interface {
	GetPing(context.Context, *GetPingRequest) (*GetPingReply, error)
}

func RegisterPingHTTPServer(s *http.Server, srv PingHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/ping", _Ping_GetPing0_HTTP_Handler(srv))
}

func _Ping_GetPing0_HTTP_Handler(srv PingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPingGetPing)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPing(ctx, req.(*GetPingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPingReply)
		return ctx.Result(200, reply)
	}
}

type PingHTTPClient interface {
	GetPing(ctx context.Context, req *GetPingRequest, opts ...http.CallOption) (rsp *GetPingReply, err error)
}

type PingHTTPClientImpl struct {
	cc *http.Client
}

func NewPingHTTPClient(client *http.Client) PingHTTPClient {
	return &PingHTTPClientImpl{client}
}

func (c *PingHTTPClientImpl) GetPing(ctx context.Context, in *GetPingRequest, opts ...http.CallOption) (*GetPingReply, error) {
	var out GetPingReply
	pattern := "/v1/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPingGetPing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
