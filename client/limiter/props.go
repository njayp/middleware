package limiter

import (
	"time"
)

type MiddlewareProps struct {
	// Count is the maximum number of requests per interval. The default is 10.
	Count int
	// Interval is the time interval for the rate limit. The default is 1 second.
	Interval time.Duration
	// Stagger is the time to wait before releasing the next token. The default is 50 milliseconds.
	Stagger time.Duration
}

// NewMiddlewareProps creates a new ClientProps with default values.
func NewMiddlewareProps(opts ...MiddlewareOpts) *MiddlewareProps {
	props := &MiddlewareProps{
		Count:    10,          // default limit
		Interval: time.Second, // default interval
	}

	for _, opt := range opts {
		opt(props)
	}

	return props
}
