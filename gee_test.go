package gee

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	router := Default()
	router.POST("/", func(c *Context) Response {
		return c.JSON(gin.H{"code": 10000, "msg": "ok", "data": nil})
	})
	w := performRequest(router, "POST", "/")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"code\":10000,\"data\":null,\"msg\":\"ok\"}", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestRouterGroup(t *testing.T) {
	router := Default()
	router.Use(func(c *Context) Response {
		c.Set("root", 123)
		return nil
	})

	api := router.Group("/api")
	api.Use(func(c *Context) Response {
		c.Set("code", 10000)
		return nil
	})
	api.POST("/user", func(c *Context) Response {
		return c.JSON(H{"code": c.MustGet("code"), "root": c.MustGet("root"), "msg": "ok", "data": nil})
	})
	w := performRequest(router, "POST", "/api/user")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"code":10000,"data":null,"msg":"ok","root":123}`, w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestEngine_NoRoute(t *testing.T) {
	router := Default()
	router.NoRoute(func(c *Context) Response {
		return c.String("not found")
	})
	w := performRequest(router, "POST", "/notfound")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "not found", w.Body.String())
}
