package trace

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/goapt/gee"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("github.com/goapt/gee")

type Option func(c *tracing)

func WithAppName(name string) Option {
	return func(c *tracing) {
		c.appName = name
	}
}

func WithPropagator(propagator propagation.TextMapPropagator) Option {
	return func(c *tracing) {
		c.propagator = propagator
	}
}

func WithTracer(tracer trace.Tracer) Option {
	return func(c *tracing) {
		c.tracer = tracer
	}
}

type tracing struct {
	tracer     trace.Tracer
	appName    string
	propagator propagation.TextMapPropagator
}

func (t *tracing) apply(c *gee.Context) {
	savedCtx := c.Request.Context()
	defer func() {
		c.Request = c.Request.WithContext(savedCtx)
	}()

	ctx := t.propagator.Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
	opts := []trace.SpanStartOption{
		trace.WithAttributes(semconv.NetAttributesFromHTTPRequest("tcp", c.Request)...),
		trace.WithAttributes(semconv.EndUserAttributesFromHTTPRequest(c.Request)...),
		trace.WithAttributes(semconv.HTTPServerAttributesFromHTTPRequest(t.appName, c.FullPath(), c.Request)...),
		trace.WithSpanKind(trace.SpanKindServer),
	}

	spanName := c.FullPath()
	if spanName == "" {
		spanName = fmt.Sprintf("HTTP %s route not found", c.Request.Method)
	}

	ctx, span := t.tracer.Start(ctx, spanName, opts...)
	defer span.End()
	c.Request = c.Request.WithContext(ctx)

	c.Next()

	status := c.Writer.Status()
	if status != http.StatusOK {
		span.RecordError(errors.New(string(c.Response.Body())))
	}

	attrs := semconv.HTTPAttributesFromHTTPStatusCode(status)
	spanStatus, spanMessage := semconv.SpanStatusFromHTTPStatusCode(status)
	span.SetAttributes(attrs...)
	span.SetStatus(spanStatus, spanMessage)
}

func New(opts ...Option) gee.Handler {
	tr := &tracing{
		appName:    "unknow",
		propagator: otel.GetTextMapPropagator(),
		tracer:     tracer,
	}

	for _, o := range opts {
		o(tr)
	}

	return func(c *gee.Context) error {
		tr.apply(c)
		return nil
	}
}
