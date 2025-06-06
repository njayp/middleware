package limiter

import (
	"time"
)

type MiddlewareOpts func(*MiddlewareProps)

// WithCount sets the maximum number of requests per second.
func WithCount(count int) MiddlewareOpts {
	return func(props *MiddlewareProps) {
		props.Count = count
	}
}

// WithInterval sets the time interval for the rate limit.
func WithInterval(interval time.Duration) MiddlewareOpts {
	return func(props *MiddlewareProps) {
		props.Interval = interval
	}
}

// WithStagger sets the time to wait before releasing the next token.
func WithStagger(stagger time.Duration) MiddlewareOpts {
	return func(props *MiddlewareProps) {
		props.Stagger = stagger
	}
}
