package gee

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	Buffer     *bytes.Buffer
	HttpStatus int
}

func NewResponseWriter(w gin.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		Buffer:         &bytes.Buffer{},
	}
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	w.Buffer.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *ResponseWriter) Status() int {
	httpStatus := w.ResponseWriter.Status()
	w.HttpStatus = httpStatus
	return httpStatus
}

func (w *ResponseWriter) WriteString(s string) (int, error) {
	w.Buffer.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
