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

func (t *tracing) apply(handler http.Handler, w http.ResponseWriter, r *http.Request) {
	savedCtx := r.Context()
	defer func() {
		r = r.WithContext(savedCtx)
	}()

	ctx := t.propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
	opts := []trace.SpanStartOption{
		trace.WithAttributes(semconv.NetAttributesFromHTTPRequest("tcp", r)...),
		trace.WithAttributes(semconv.EndUserAttributesFromHTTPRequest(r)...),
		trace.WithAttributes(semconv.HTTPServerAttributesFromHTTPRequest(t.appName, r.URL.Path, r)...),
		trace.WithSpanKind(trace.SpanKindServer),
	}

	spanName := r.URL.Path
	if spanName == "" {
		spanName = fmt.Sprintf("HTTP %s route not found", r.Method)
	}

	ctx, span := t.tracer.Start(ctx, spanName, opts...)
	defer span.End()
	r = r.WithContext(ctx)

	handler.ServeHTTP(w, r)
	ww := gee.Response(w)
	status := ww.Status()
	if status != http.StatusOK {
		span.RecordError(errors.New(string(ww.Body())))
	}

	attrs := semconv.HTTPAttributesFromHTTPStatusCode(status)
	spanStatus, spanMessage := semconv.SpanStatusFromHTTPStatusCode(status)
	span.SetAttributes(attrs...)
	span.SetStatus(spanStatus, spanMessage)
}

func New(opts ...Option) gee.Middleware {
	tr := &tracing{
		appName:    "unknow",
		propagator: otel.GetTextMapPropagator(),
		tracer:     tracer,
	}

	for _, o := range opts {
		o(tr)
	}

	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tr.apply(handler, w, r)
		})
	}
}
