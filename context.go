package gee

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/goapt/gee/binding"
	"github.com/goapt/gee/render"
)

func getHttpStatus(c *Context, status int) int {
	if c.httpStatus == 0 {
		return status
	}
	return c.httpStatus
}

type Context struct {
	*gin.Context
	httpStatus int
	StartTime  time.Time
	renderHook []render.Hook
}

func (c *Context) Render(r render.Render) {
	r.Hooks(c.renderHook)
	c.Context.Render(getHttpStatus(c, 200), r)
}

func (c *Context) ShouldBindJSON(obj interface{}) error {
	return c.ShouldBindWith(obj, binding.JSON)
}

func (c *Context) ShouldBindBodyJSON(obj interface{}) error {
	return c.ShouldBindBodyWith(obj, binding.JSON)
}

// GetBody read body data and restore request body data
func (c *Context) GetBody() ([]byte, error) {
	body, err := c.Context.GetRawData()
	if err != nil {
		return nil, err
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
	return body, nil
}

func (c *Context) AddRenderHook(fn render.Hook) {
	c.renderHook = append(c.renderHook, fn)
}

func (c *Context) ResponseWriter() *ResponseWriter {
	if resp, ok := c.Writer.(*ResponseWriter); ok {
		return resp
	}
	return nil
}

func (c *Context) ResponseBody() []byte {
	if resp := c.ResponseWriter(); resp != nil {
		return resp.Body()
	}
	return nil
}

func (c *Context) BindJSON(obj interface{}) error {
	return c.MustBindWith(obj, binding.JSON)
}

func (c *Context) Status(status int) {
	c.httpStatus = status
}

func (c *Context) JSON(data interface{}) Response {
	return &JSONResponse{
		Context: c,
		Data:    data,
	}
}

func (c *Context) XML(data interface{}) Response {
	return &XMLResponse{
		Context: c,
		Data:    data,
	}
}

func (c *Context) Redirect(location string) Response {
	return &RedirectResponse{
		Context:  c,
		Location: location,
	}
}

func (c *Context) String(format string, values ...interface{}) Response {
	return &StringResponse{
		Context: c,
		Format:  format,
		Data:    values,
	}
}

func (c *Context) RequestId() string {
	requestId := c.Request.Header.Get("X-Request-ID")
	if requestId == "" {
		requestId = uuid.New().String()
		c.Request.Header.Set("X-Request-ID", requestId)
	}

	return requestId
}
