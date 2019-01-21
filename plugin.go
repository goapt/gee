package very

import "github.com/gin-gonic/gin"

type IPlugin interface {
	Name() string
}

type IGetToken interface {
	GetToken(c *Context) (string , error)
}


var PluginMiddleware = func(plugins ...IPlugin) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := getContext(c)
		ctx.Plugins = plugins
	}
}