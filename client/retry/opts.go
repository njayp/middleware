package retry

import (
	"time"
)

type MiddlewareOpts func(*MiddlewareProps)

// WithCount sets the maximum number of retries.
// The default is 10 retries.
func WithCount(count int) MiddlewareOpts {
	return func(props *MiddlewareProps) {
		props.Count = count
	}
}

// WithBackoff sets the time to wait before retrying a request after a rate limit error.
// The default is 1 second.
func WithBackoff(interval time.Duration) MiddlewareOpts {
	return func(props *MiddlewareProps) {
		props.BackOff = interval
	}
}
