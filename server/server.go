package server

import (
	"github.com/xdatk/micro/registry"
)

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

type Handler interface {
	Name() string
	Handler() interface{}
	Endpoint() []*registry.Endpoint
	Options() HandlerOptions
}

type Subscriber interface {
	Topic() string
	Subscriber() interface{}
	Endpoint() []*registry.Endpoint
	Options() SubscriberOptions
}
