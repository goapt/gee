package gee

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type header struct {
	Key   string
	Value string
}

func performRequest(r http.Handler, method, path string, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHandle(t *testing.T) {
	r := New()
	var testHander HandlerFunc = func(c *Context) Response {
		return c.JSON(gin.H{"code": 10000, "msg": "ok", "data": nil})
	}
	r.POST("/", testHander)
	w := performRequest(r, "POST", "/")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"code\":10000,\"data\":null,\"msg\":\"ok\"}", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestHandle_Wrap(t *testing.T) {
	r := New()
	var testHander gin.HandlerFunc = func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 10000, "msg": "ok", "data": nil})
	}
	r.POST("/", Wrap(testHander))
	w := performRequest(r, "POST", "/")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"code\":10000,\"data\":null,\"msg\":\"ok\"}", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestHandleFunc(t *testing.T) {
	r := New()
	var testHander HandlerFunc = func(c *Context) Response {
		return c.JSON(gin.H{"code": 10000, "msg": "ok", "data": nil})
	}
	r.POST("/", testHander)
	w := performRequest(r, "POST", "/")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"code\":10000,\"data\":null,\"msg\":\"ok\"}", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}
