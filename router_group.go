package gee

import (
	"github.com/gin-gonic/gin"
)

// IRoutes defines all router handle interface.
type IRoutes interface {
	Use(...HandlerFunc) IRoutes
	Any(string, ...HandlerFunc) IRoutes
	GET(string, ...HandlerFunc) IRoutes
	POST(string, ...HandlerFunc) IRoutes
	DELETE(string, ...HandlerFunc) IRoutes
	PATCH(string, ...HandlerFunc) IRoutes
	PUT(string, ...HandlerFunc) IRoutes
	OPTIONS(string, ...HandlerFunc) IRoutes
	HEAD(string, ...HandlerFunc) IRoutes
}

type RouterGroup struct {
	*gin.RouterGroup
}

func (r *RouterGroup) warp(handlers []HandlerFunc) []gin.HandlerFunc {
	var ginHandlers []gin.HandlerFunc
	for _, v := range handlers {
		ginHandlers = append(ginHandlers, HandleFunc(v))
	}

	return ginHandlers
}

func (r *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
	r.RouterGroup.Use(r.warp(middleware)...)
	return r
}

func (r *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.RouterGroup.Any(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.RouterGroup.POST(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.RouterGroup.GET(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.RouterGroup.DELETE(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.RouterGroup.OPTIONS(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.RouterGroup.PUT(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.RouterGroup.PATCH(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.RouterGroup.HEAD(relativePath, r.warp(handlers)...)
	return r
}

func (r *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) IRoutes {
	return &RouterGroup{
		r.RouterGroup.Group(relativePath, r.warp(handlers)...),
	}
}
