package gee

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResponseWriter_Write(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	nw := NewResponseWriter(c.Writer)
	c.Writer = nw
	_, err := c.Writer.WriteString("hello")
	assert.NoError(t, err)
	assert.Equal(t, nw.Buffer.String(), "hello")
}
