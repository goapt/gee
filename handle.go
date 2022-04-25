package gee

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler = func(c *Context) error

var ErrorHandler = func(c *Context, err error) error {
	c.Abort()
	return c.String(err.Error())
}

func HandlerFunc(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := getContext(c)
		if err := h(ctx); err != nil {
			if err := ErrorHandler(ctx, err); err != nil {
				ctx.Status(http.StatusInternalServerError)
				log.Printf("[ERROR] %s\n", err)
			}
		}
	}
}

func Wrap(h gin.HandlerFunc) Handler {
	return func(c *Context) error {
		h(c.Context)
		return nil
	}
}

func WrapH(h http.Handler) Handler {
	return Wrap(gin.WrapH(h))
}

const contextKey = "__context"

func getContext(c *gin.Context) *Context {
	var ctx1 *Context
	if ctx, ok := c.Get(contextKey); !ok {
		resp := newResponseWriter(c.Writer)
		c.Writer = resp
		ctx1 = &Context{
			Context:  c,
			Response: resp,
		}
		c.Set(contextKey, ctx1)
	} else {
		ctx1 = ctx.(*Context)
	}
	return ctx1
}
