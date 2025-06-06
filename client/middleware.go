package client

import (
	"net/http"
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
