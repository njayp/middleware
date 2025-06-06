package retry

import (
	"log/slog"
	"net/http"
)

func NewMiddlewareRoundTripper(rt http.RoundTripper, opts ...MiddlewareOpts) *RetryRoundTripper {
	props := NewMiddlewareProps(opts...)

	if rt == nil {
		rt = http.DefaultTransport // Use default transport if none provided
	}

	return &RetryRoundTripper{
		MaxRetries: props.Count,
		Next:       rt,
	}
}

type RetryRoundTripper struct {
	MaxRetries int
	Next       http.RoundTripper
}

func (r *RetryRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for range r.MaxRetries {
		resp, err = r.Next.RoundTrip(req)
		if err == nil && resp.StatusCode < 500 {
			return resp, nil // Return response if no error and status code is not a server error
		}
	}

	slog.Error("RetryRoundTripper failed after retries",
		"method", req.Method,
		"url", req.URL.String(),
		"error", err,
		"status", resp.StatusCode,
	)

	return resp, err // Return the last response or error after retries
}
