package client

import (
	"context"
	"github.com/xdatk/micro/codec"
	"github.com/xdatk/micro/registry"
)

type Request interface {
	Service() string
	Method() string
	Endpoint() string
	ContentType() string
	Body() interface{}
	Codec() codec.Writer
	Stream() bool
}

type Response interface {
	Codec() codec.Reader
	Header() map[string]string
	Read() ([]byte, error)
}

type Stream interface {
	Context() context.Context
	Request() Request
	Response() Response
	Send(interface{}) error
	Recv(interface{}) error
	Error() error
	Close() error
}

type CallFunc func(ctx context.Context, node *registry.Node, req Request, rsp interface{}, opts CallOptions) error

type CallWrapper func(CallFunc) CallFunc

type StreamWrapper func(Stream) Stream

type Message interface {
	Topic() string
	Payload() interface{}
	ContentType() string
}

type Client interface {
	Init(...Option) error
	Options() Options
	NewRequest(string, string, interface{}, ...RequestOption) Request
	Call(context.Context, Request, interface{}, ...CallOption) error
	Stream(context.Context, Request, ...CallOption) (Stream, error)
	NewMessage(string, interface{}, ...MessageOption) Message
	Publish(context.Context, Message, ...PublishOption) error
	String() string
}

type Wrapper func(Client) Client
