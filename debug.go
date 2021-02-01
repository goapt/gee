package gee

import (
	"path/filepath"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var debugToken = "058759a8c4d920576348854616a58f3f"

func SetDebugToken(token string) {
	debugToken = token
}

func authDebug() gin.HandlerFunc {
	return func(c *gin.Context) {
		if pwd := c.Query("debug_token"); pwd != debugToken {
			c.JSON(403, map[string]string{"code": "40000", "message": "access denied"})
			c.Abort()
		}
		c.Next()
	}
}

func DebugRoute(router *Engine, path ...string) {
	p := "/debug"
	// custom router path
	if len(path) > 0 && path[0] != "" {
		p = filepath.Join(path[0], p)
	}
	debugger := router.Group(p, Wrap(authDebug()))
	pprof.RouteRegister(debugger.RouterGroup, "pprof")
	debugger.GET("/metrics", WrapH(promhttp.Handler()))
}

func DebugWithRouteGroup(router *RouterGroup) {
	debugger := router.Group("/debug", Wrap(authDebug()))
	pprof.RouteRegister(debugger.RouterGroup, "pprof")
	debugger.GET("/metrics", WrapH(promhttp.Handler()))
}
