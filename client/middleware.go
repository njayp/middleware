package middleware

import (
	"net/http"

	"github.com/njayp/middleware/client/limiter"
	"github.com/njayp/middleware/client/retry"
)

// BuildTransport constructs an http.RoundTripper with the provided middleware options.
// It applies each middleware function to the default transport in the order they are provided.
// This allows for flexible composition of middleware functionalities, such as rate limiting, logging, etc.
func BuildTransport(opts ...Middleware) http.RoundTripper {
	transport := http.DefaultTransport

	for _, opt := range opts {
		transport = opt(transport)
	}

	return transport
}

type Middleware func(http.RoundTripper) http.RoundTripper

func WithLimiter(opts ...limiter.MiddlewareOpts) Middleware {
	return func(rt http.RoundTripper) http.RoundTripper {
		return limiter.NewMiddlewareRoundTripper(rt, opts...)
	}
}

func WithRetries(opts ...retry.MiddlewareOpts) Middleware {
	return func(rt http.RoundTripper) http.RoundTripper {
		return retry.NewMiddlewareRoundTripper(rt, opts...)
	}
}
