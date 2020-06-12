package server

import "context"

type Options struct {
	Name string
}

type Option func(*Options)

func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

type HandlerOptions struct {
	Internal bool
	Metadata map[string]map[string]string
}

type HandlerOption func(*HandlerOptions)

func InternalHandler(b bool) HandlerOption {
	return func(o *HandlerOptions) {
		o.Internal = b
	}
}

func EndpointMetadata(name string, md map[string]string) HandlerOption {
	return func(o *HandlerOptions) {
		o.Metadata[name] = md
	}
}

type SubscriberOptions struct {
	AutoAck  bool
	Queue    string
	Internal bool
	Context  context.Context
}

type SubscriberOption func(*SubscriberOptions)

func DisableAutoAck() SubscriberOption {
	return func(o *SubscriberOptions) {
		o.AutoAck = false
	}
}

func SubscriberQueue(n string) SubscriberOption {
	return func(o *SubscriberOptions) {
		o.Queue = n
	}
}

func InternalSubscriber(b bool) SubscriberOption {
	return func(o *SubscriberOptions) {
		o.Internal = b
	}
}

func SubscriberContext(ctx context.Context) SubscriberOption {
	return func(o *SubscriberOptions) {
		o.Context = ctx
	}
}
