package gee

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRouterGroupBasicHandle(t *testing.T) {
	performRequestInGroup(t, http.MethodGet)
	performRequestInGroup(t, http.MethodPost)
	performRequestInGroup(t, http.MethodPut)
	performRequestInGroup(t, http.MethodPatch)
	performRequestInGroup(t, http.MethodDelete)
	performRequestInGroup(t, http.MethodHead)
	performRequestInGroup(t, http.MethodOptions)
}

func TestRouterGroup_Any(t *testing.T) {
	r := New()
	var testHander = func(c *Context) error {
		return c.JSON(gin.H{"code": 10000, "msg": "ok", "data": nil})
	}
	r.Any("/", testHander)
	w := performRequest(r, "POST", "/")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"code\":10000,\"data\":null,\"msg\":\"ok\"}", w.Body.String())
}

func TestRouterGroup_Handle(t *testing.T) {
	r := New()
	var testHander = func(c *Context) error {
		return c.JSON(gin.H{"code": 10000, "msg": "ok", "data": nil})
	}
	r.Handle(http.MethodPost, "/", testHander)
	w := performRequest(r, "POST", "/")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"code\":10000,\"data\":null,\"msg\":\"ok\"}", w.Body.String())
}

func performRequestInGroup(t *testing.T, method string) {
	router := New()
	v1 := router.Group("v1", func(c *Context) error { return nil })
	assert.Equal(t, "/v1", v1.BasePath())

	login := v1.Group("/login/", func(c *Context) error { return nil }, func(c *Context) error { return nil })
	assert.Equal(t, "/v1/login/", login.BasePath())

	handler := func(c *Context) error {
		c.httpStatus = http.StatusBadRequest
		return c.String("the method was %s and uri %s", c.Request.Method, c.Request.URL)
	}

	switch method {
	case http.MethodGet:
		v1.GET("/test", handler)
		login.GET("/test", handler)
	case http.MethodPost:
		v1.POST("/test", handler)
		login.POST("/test", handler)
	case http.MethodPut:
		v1.PUT("/test", handler)
		login.PUT("/test", handler)
	case http.MethodPatch:
		v1.PATCH("/test", handler)
		login.PATCH("/test", handler)
	case http.MethodDelete:
		v1.DELETE("/test", handler)
		login.DELETE("/test", handler)
	case http.MethodHead:
		v1.HEAD("/test", handler)
		login.HEAD("/test", handler)
	case http.MethodOptions:
		v1.OPTIONS("/test", handler)
		login.OPTIONS("/test", handler)
	default:
		panic("unknown method")
	}

	w := performRequest(router, method, "/v1/login/test")
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "the method was "+method+" and uri /v1/login/test", w.Body.String())

	w = performRequest(router, method, "/v1/test")
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "the method was "+method+" and uri /v1/test", w.Body.String())
}
