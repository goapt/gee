package gee

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/goapt/gee/binding"
	"github.com/goapt/gee/render"
)

type Context struct {
	*gin.Context
	httpStatus int
	Response   *Response
}

func (c *Context) getHttpStatus(status int) int {
	if c.httpStatus == 0 {
		return status
	}
	return c.httpStatus
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

func (c *Context) BindJSON(obj interface{}) error {
	return c.MustBindWith(obj, binding.JSON)
}

func (c *Context) Status(status int) {
	c.httpStatus = status
}

func (c *Context) JSON(data interface{}) error {
	c.Context.Render(c.getHttpStatus(200), &render.JSON{Data: data})
	return nil
}

func (c *Context) XML(data interface{}) error {
	c.Context.XML(c.getHttpStatus(200), data)
	return nil
}

func (c *Context) YAML(data interface{}) error {
	c.Context.YAML(c.getHttpStatus(200), data)
	return nil
}

func (c *Context) Redirect(location string) error {
	c.Context.Redirect(c.getHttpStatus(302), location)
	return nil
}

func (c *Context) String(format string, values ...interface{}) error {
	c.Context.String(c.getHttpStatus(200), format, values...)
	return nil
}

func (c *Context) RequestId() string {
	requestId := c.Request.Header.Get("X-Request-ID")
	if requestId == "" {
		requestId = uuid.New().String()
		c.Request.Header.Set("X-Request-ID", requestId)
	}

	return requestId
}
