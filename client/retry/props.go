package retry

import (
	"time"
)

type MiddlewareProps struct {
	// Count is the maximum number of requests per interval. The default is 10.
	Count int
	// BackOff is the time to wait before retrying a request after a rate limit error. The default is 1 second.
	BackOff time.Duration
}

// NewMiddlewareProps creates a new ClientProps with default values.
func NewMiddlewareProps(opts ...MiddlewareOpts) *MiddlewareProps {
	props := &MiddlewareProps{
		Count:   10,          // default limit
		BackOff: time.Second, // default interval
	}

	for _, opt := range opts {
		opt(props)
	}

	return props
}
