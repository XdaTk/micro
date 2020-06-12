package registry

import (
	"context"
	"time"
)

type Options struct {
	Context context.Context
}

type Option func(*Options)

type RegisterOptions struct {
	TTL     time.Duration
	Context context.Context
}

type RegisterOption func(*RegisterOptions)

func RegisterTTL(t time.Duration) RegisterOption {
	return func(o *RegisterOptions) {
		o.TTL = t
	}
}

func RegisterContext(ctx context.Context) RegisterOption {
	return func(o *RegisterOptions) {
		o.Context = ctx
	}
}

type DeregisterOptions struct {
	Context context.Context
}

type DeregisterOption func(*DeregisterOptions)

func DeregisterContext(ctx context.Context) DeregisterOption {
	return func(o *DeregisterOptions) {
		o.Context = ctx
	}
}

type GetOptions struct {
	Context context.Context
}

type GetOption func(*GetOptions)

func GetContext(ctx context.Context) GetOption {
	return func(o *GetOptions) {
		o.Context = ctx
	}
}

type ListOptions struct {
	Context context.Context
}

type ListOption func(*ListOptions)

func ListContext(ctx context.Context) ListOption {
	return func(o *ListOptions) {
		o.Context = ctx
	}
}

type WatchOptions struct {
	Service string
	Context context.Context
}

type WatchOption func(*WatchOptions)

func WatchService(name string) WatchOption {
	return func(o *WatchOptions) {
		o.Service = name
	}
}

func WatchContext(ctx context.Context) WatchOption {
	return func(o *WatchOptions) {
		o.Context = ctx
	}
}
