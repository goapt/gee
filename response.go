package gee

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

type Response struct {
	gin.ResponseWriter
	body   *bytes.Buffer
	before func(w *Response)
}

func newResponseWriter(w gin.ResponseWriter) *Response {
	return &Response{
		ResponseWriter: w,
		body:           &bytes.Buffer{},
	}
}

// Before is a function which will be called before the response is written to the client.
func (w *Response) Before(fn func(w *Response)) {
	w.before = fn
}

func (w *Response) Write(b []byte) (int, error) {
	w.body.Write(b)
	if w.before != nil {
		w.before(w)
	}
	return w.ResponseWriter.Write(b)
}

func (w *Response) Body() []byte {
	return w.body.Bytes()
}

func (w *Response) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	if w.before != nil {
		w.before(w)
	}
	return w.ResponseWriter.WriteString(s)
}
