package trace

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
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
	var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, span := tr.Start(r.Context(), "testHandler")
		span.End()
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("test"))
	})

	m := gee.NewRouter()
	m.Use(New(WithAppName("test")))
	m.Get("/dummy/impl", testHandler)

	ts := httptest.NewServer(m)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/dummy/impl")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	assert.Equal(t, 500, resp.StatusCode)
	assert.JSONEq(t, `test`, string(respBody))

	spans := sr.Ended()
	for _, ss := range spans {
		b, _ := json.Marshal(ss.Attributes())
		fmt.Println("Trace:", ss.Name(), string(b))
	}
}
