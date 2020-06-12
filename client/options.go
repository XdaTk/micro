package client

import "context"

type Options struct {
}

type Option func(*Options)

type RequestOptions struct {
	ContentType string
	Stream      bool
	Context     context.Context
}

type RequestOption func(*RequestOptions)

type CallOptions struct {
}

type CallOption func(*CallOptions)

type MessageOptions struct {
	ContentType string
}

type MessageOption func(*MessageOptions)

type PublishOptions struct {
	Context context.Context
}

type PublishOption func(*PublishOptions)
