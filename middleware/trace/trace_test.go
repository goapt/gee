package trace

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/goapt/gee"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestNewTrace(t *testing.T) {
	sr := &tracetest.SpanRecorder{}
	tp := trace.NewTracerProvider(trace.WithSpanProcessor(sr))
	tr := tp.Tracer("trace/test")
	var testHandler = func(c *gee.Context) error {
		_, span := tr.Start(c.Request.Context(), "testHandler")
		span.End()
		c.Status(http.StatusInternalServerError)
		return c.JSON(gee.H{
			"code": "SystemError",
			"msg":  "系统错误，请重试",
		})
	}
	req := gee.NewTestRequest("/dummy/impl", New(WithAppName("test")), testHandler)
	resp, err := req.Get()
	assert.NoError(t, err)
	assert.Equal(t, 500, resp.Code)
	assert.JSONEq(t, `{"code":"SystemError","msg":"系统错误，请重试"}`, resp.GetBodyString())

	spans := sr.Ended()

	for _, ss := range spans {
		b, _ := json.Marshal(ss.Attributes())
		fmt.Println("Trace:", ss.Name(), string(b))
	}
}
