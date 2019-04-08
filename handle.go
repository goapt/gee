package very

import (
	"github.com/gin-gonic/gin"
	"time"
)

type IHandler interface {
	Handle(c *Context) Response
}

type HandlerFunc func(c *Context) Response

func (h HandlerFunc) Handle(c *Context) Response {
	return h(c)
}

const contextKey = "__context"

func Handle(handler IHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := getContext(c)
		resp := handler.Handle(ctx)
		if resp != nil {
			ctx.Response = resp
			ctx.Response.Render()
		}
	}
}

func Middleware(handler IHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := getContext(c)
		resp := handler.Handle(ctx)
		if resp != nil {
			c.Abort()
			ctx.Response = resp
			ctx.Response.Render()
		}
	}
}

func getContext(c *gin.Context) *Context {
	ctx, ok := c.Get(contextKey)
	var ctx1 *Context
	if !ok {
		ctx1 = &Context{
			Context:   c,
			LogInfo:   make(map[string]interface{}),
			StartTime: time.Now(),
		}
		c.Set(contextKey, ctx1)
	} else {
		ctx1 = ctx.(*Context)
	}
	return ctx1
}
