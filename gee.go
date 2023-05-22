package gee

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	vbinding "github.com/goapt/gee/binding"
)

func init() {
	binding.Validator = new(vbinding.DefaultValidator)
}

type Engine struct {
	*gin.Engine
	RouterGroup
}

func Default() *Engine {
	engine := gin.Default()
	return &Engine{
		engine,
		RouterGroup{
			&engine.RouterGroup,
		},
	}
}

func New() *Engine {
	engine := gin.New()
	return &Engine{
		engine,
		RouterGroup{
			&engine.RouterGroup,
		},
	}
}

func (e *Engine) Use(middleware ...Handler) IRoutes {
	e.Engine.Use(e.RouterGroup.warp(middleware)...)
	return e
}

func (e *Engine) NoRoute(handlers ...Handler) {
	e.Engine.NoRoute(e.RouterGroup.warp(handlers)...)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	e.Engine.ServeHTTP(w, req)
}

// H is a shortcut for map[string]interface{}
type H map[string]any

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

func CreateTestContext(w http.ResponseWriter) (*Context, *Engine) {
	c, engine := gin.CreateTestContext(w)
	return getContext(c), &Engine{
		engine,
		RouterGroup{
			&engine.RouterGroup,
		},
	}
}

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = gin.DebugMode
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = gin.ReleaseMode
	// TestMode indicates gin mode is test.
	TestMode = gin.TestMode
)

func SetMode(value string) {
	gin.SetMode(value)
}
