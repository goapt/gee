package gee

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	Buffer *bytes.Buffer
}

// 决定不导出，直接让项目报错，从而避免兼容性问题
func newResponseWriter(w gin.ResponseWriter) *ResponseWriter {

	// 为了兼容旧版本在项目里面直接使用NewResponseWriter的情况
	// 改用不导出方案
	// if resp, ok := w.(*ResponseWriter); ok {
	// 	w = resp.ResponseWriter
	// }

	return &ResponseWriter{
		ResponseWriter: w,
		Buffer:         &bytes.Buffer{},
	}
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	w.Buffer.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *ResponseWriter) Body() []byte {
	return w.Buffer.Bytes()
}

func (w *ResponseWriter) WriteString(s string) (int, error) {
	w.Buffer.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
