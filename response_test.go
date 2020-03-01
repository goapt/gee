package gee

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestApiResponse_Render(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	resp := &JSONResponse{
		HttpStatus: 200,
		Context:    ctx,
		Data: map[string]interface{}{
			"id":   1,
			"name": "test",
		},
	}

	resp.Render()
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"id":1,"name":"test"}`, w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}
