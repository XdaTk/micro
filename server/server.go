package server

import (
	"context"
	"github.com/xdatk/micro/codec"
	"github.com/xdatk/micro/registry"
)

type Handler interface {
	Name() string
	Handler() interface{}
	Endpoint() []*registry.Endpoint
	Options() HandlerOptions
}

type Request interface {
	Service() string
	Method() string
	Endpoint() string
	ContentType() string
	Header() map[string]string
	Body() interface{}
	Read() ([]byte, error)
	Codec() codec.Reader
	Stream() bool
}

type Response interface {
	Codec() codec.Writer
	WriteHeader(map[string]string)
	Write([]byte) error
}

type Stream interface {
	Context() context.Context
	Request() Request
	Send(interface{}) error
	Recv(interface{}) error
	Error() error
	Close() error
}

type HandlerFunc func(context.Context, Request, interface{}) error

type HandlerWrapper func(HandlerFunc) HandlerFunc

type StreamWrapper func(Stream) Stream

type Subscriber interface {
	Topic() string
	Subscriber() interface{}
	Endpoint() []*registry.Endpoint
	Options() SubscriberOptions
}

type Message interface {
	Topic() string
	Payload() interface{}
	ContentType() string
	Header() map[string]string
	Body() []byte
	Codec() codec.Reader
}

type Server interface {
	Init(...Option) error
	Options() Options
	NewHandler(interface{}, ...HandlerOption) Handler
	Handle(Handler) error
	NewSubscriber(string, interface{}, ...SubscriberOption) Subscriber
	Subscribe(Subscriber) error
	Start() error
	Stop() error
	String() string
}

type SubscriberFunc func(ctx context.Context, msg Message) error

type SubscriberWrapper func(SubscriberFunc) SubscriberFunc
