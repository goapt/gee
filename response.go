package gee

import (
	"github.com/goapt/gee/render"
)

type Response interface {
	Render()
}

type RenderFunc func()

func (e RenderFunc) Render() { e() }

type JSONResponse struct {
	Context *Context    `json:"-"`
	Data    interface{} `json:"data"`
}

func (c *JSONResponse) Render() {
	c.Context.Render(&render.JSON{Data: c.Data})
}

type XMLResponse struct {
	Context *Context    `json:"-"`
	Data    interface{} `json:"data"`
}

func (c *XMLResponse) Render() {
	c.Context.Render(&render.XML{Data: c.Data})
}

type RedirectResponse struct {
	Context  *Context `json:"-"`
	Location string
}

func (c *RedirectResponse) Render() {
	c.Context.Context.Redirect(getHttpStatus(c.Context, 302), c.Location)
}

type StringResponse struct {
	Context *Context `json:"-"`
	Format  string
	Data    []interface{}
}

func (c *StringResponse) Render() {
	c.Context.Render(&render.String{Format: c.Format, Data: c.Data})
}
