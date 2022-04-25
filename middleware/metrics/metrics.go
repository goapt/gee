package metrics

import (
	"strconv"
	"time"

	"github.com/goapt/gee"
)

// Counter is metrics counter.
type Counter interface {
	With(lvs ...string) Counter
	Inc()
	Add(delta float64)
}

// Gauge is metrics gauge.
type Gauge interface {
	With(lvs ...string) Gauge
	Set(value float64)
	Add(delta float64)
	Sub(delta float64)
}

// Observer is metrics observer.
type Observer interface {
	With(lvs ...string) Observer
	Observe(float64)
}

// Option is metrics option.
type Option func(*options)

// WithRequests with requests counter.
func WithRequests(c Counter) Option {
	return func(o *options) {
		o.requests = c
	}
}

// WithSeconds with seconds histogram.
func WithSeconds(c Observer) Option {
	return func(o *options) {
		o.seconds = c
	}
}

type options struct {
	// counter: <client/server>_requests_code_total{operation, code, reason}
	requests Counter
	// histogram: <client/server>_requests_seconds_bucket{operation}
	seconds Observer
}

// New is middleware server-side metrics.
func New(opts ...Option) gee.Handler {
	options := options{}
	for _, o := range opts {
		o(&options)
	}
	return func(c *gee.Context) error {
		var (
			code      int
			reason    string // TODO It's not done
			operation string
		)
		startTime := time.Now()
		c.Next()
		code = c.Writer.Status()

		if options.requests != nil {
			options.requests.With(operation, strconv.Itoa(code), reason).Inc()
		}
		if options.seconds != nil {
			options.seconds.With(operation).Observe(time.Since(startTime).Seconds())
		}
		return nil
	}
}
