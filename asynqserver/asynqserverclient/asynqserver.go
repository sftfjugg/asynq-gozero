// Code generated by goctl. DO NOT EDIT!
// Source: asynqserver.proto

package asynqserverclient

import (
	"context"

	"gozero-asynq/asynqserver/asynqserver"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = asynqserver.Request
	Response = asynqserver.Response

	Asynqserver interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultAsynqserver struct {
		cli zrpc.Client
	}
)

func NewAsynqserver(cli zrpc.Client) Asynqserver {
	return &defaultAsynqserver{
		cli: cli,
	}
}

func (m *defaultAsynqserver) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := asynqserver.NewAsynqserverClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
