package gee

import (
	"github.com/gin-gonic/gin"
)

// IRoutes defines all router handle interface.
type IRoutes interface {
	Use(...Handler) IRoutes
	Any(string, ...Handler) IRoutes
	GET(string, ...Handler) IRoutes
	POST(string, ...Handler) IRoutes
	DELETE(string, ...Handler) IRoutes
	PATCH(string, ...Handler) IRoutes
	PUT(string, ...Handler) IRoutes
	OPTIONS(string, ...Handler) IRoutes
	HEAD(string, ...Handler) IRoutes
}

type RouterGroup struct {
	*gin.RouterGroup
}

func (r *RouterGroup) warp(handlers []Handler) []gin.HandlerFunc {
	var ginHandlers []gin.HandlerFunc
	for _, v := range handlers {
		ginHandlers = append(ginHandlers, HandlerFunc(v))
	}

	return ginHandlers
}

func (r *RouterGroup) Use(middleware ...Handler) IRoutes {
	r.RouterGroup.Use(r.warp(middleware)...)
	return r
}

func (r *RouterGroup) Any(relativePath string, handlers ...Handler) IRoutes {
	r.RouterGroup.Any(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) POST(relativePath string, handlers ...Handler) IRoutes {
	r.RouterGroup.POST(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) GET(relativePath string, handlers ...Handler) IRoutes {
	r.RouterGroup.GET(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) DELETE(relativePath string, handlers ...Handler) IRoutes {
	r.RouterGroup.DELETE(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) OPTIONS(relativePath string, handlers ...Handler) IRoutes {
	r.RouterGroup.OPTIONS(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) PUT(relativePath string, handlers ...Handler) IRoutes {
	r.RouterGroup.PUT(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) PATCH(relativePath string, handlers ...Handler) IRoutes {
	r.RouterGroup.PATCH(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) HEAD(relativePath string, handlers ...Handler) IRoutes {
	r.RouterGroup.HEAD(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) Handle(method string, relativePath string, handlers ...Handler) IRoutes {
	r.RouterGroup.Handle(method, relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) Group(relativePath string, handlers ...Handler) *RouterGroup {
	return &RouterGroup{
		r.RouterGroup.Group(relativePath, r.warp(handlers)...),
	}
}
