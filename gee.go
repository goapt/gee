package gee

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Engine struct {
	*gin.Engine
	RouterGroup
}

func Default() *Engine {
	engine := gin.Default()
	binding.Validator = new(DefaultValidator)
	return &Engine{
		engine,
		RouterGroup{
			&engine.RouterGroup,
		},
	}
}

func New() *Engine {
	engine := gin.New()
	binding.Validator = new(DefaultValidator)
	return &Engine{
		engine,
		RouterGroup{
			&engine.RouterGroup,
		},
	}
}

func (e *Engine) Use(middleware ...HandlerFunc) IRoutes {
	e.Engine.Use(e.RouterGroup.warp(middleware)...)
	return e
}

func (e *Engine) NoRoute(handlers ...HandlerFunc) {
	e.Engine.NoRoute(e.RouterGroup.warp(handlers)...)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	e.Engine.ServeHTTP(w, req)
}

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

// MarshalXML allows type H to be used with xml.Marshal.
func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{
		Space: "",
		Local: "map",
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for key, value := range h {
		elem := xml.StartElement{
			Name: xml.Name{Space: "", Local: key},
			Attr: []xml.Attr{},
		}
		if err := e.EncodeElement(value, elem); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
