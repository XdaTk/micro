package micro

import (
	"context"
	"github.com/xdatk/micro/client"
	"github.com/xdatk/micro/config/cmd"
	"github.com/xdatk/micro/server"
)

type Options struct {
	Cmd         cmd.Cmd
	Client      client.Client
	Server      server.Server
	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error
	Signal      bool
	Context     context.Context
}

type Option func(*Options)

func Cmd(c cmd.Cmd) Option {
	return func(o *Options) {
		o.Cmd = c
	}
}

func Client(c client.Client) Option {
	return func(o *Options) {
		o.Client = c
	}
}

func Server(s server.Server) Option {
	return func(o *Options) {
		o.Server = s
	}
}

func BeforeStart(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}

func BeforeStop(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStop = append(o.BeforeStop, fn)
	}
}

func AfterStart(fn func() error) Option {
	return func(o *Options) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}

func AfterStop(fn func() error) Option {
	return func(o *Options) {
		o.AfterStop = append(o.AfterStop, fn)
	}
}

func HandleSignal(b bool) Option {
	return func(o *Options) {
		o.Signal = b
	}
}

func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}
