package gee

import (
	"time"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Handle(c *Context) Response
}

type HandlerFunc func(c *Context) Response

func (h HandlerFunc) Handle(c *Context) Response {
	return h(c)
}

const (
	contextKey = "__gee_context__"
)

func Handle(handler IHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := getContext(c)
		if resp := handler.Handle(ctx); resp != nil {
			ctx.Response = resp
			ctx.Response.Render()
		}
	}
}

func HandleFunc(handler func(c *Context) Response) gin.HandlerFunc {
	return Handle(HandlerFunc(handler))
}

func Wrap(h gin.HandlerFunc) HandlerFunc {
	return func(c *Context) Response {
		h(c.Context)
		return nil
	}
}

func getContext(c *gin.Context) *Context {
	var ctx1 *Context
	if ctx, ok := c.Get(contextKey); !ok {
		ctx1 = &Context{
			Context:   c,
			StartTime: time.Now(),
		}
		c.Set(contextKey, ctx1)
	} else {
		ctx1 = ctx.(*Context)
	}
	return ctx1
}
