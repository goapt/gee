package gee

import (
	"bytes"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/goapt/gee/binding"
	"github.com/goapt/gee/render"
)

func getHttpStatus(c *Context, status int) int {
	if c.HttpStatus == 0 {
		return status
	}
	return c.HttpStatus
}

type Context struct {
	*gin.Context
	HttpStatus int
	Response   Response
	LogInfo    map[string]interface{}
	StartTime  time.Time
	RenderHook []render.Hook
}

func (c *Context) Render(r render.Render) {
	r.Hooks(c.RenderHook)
	c.Context.Render(getHttpStatus(c, 200), r)
}

func (c *Context) ShouldBindJSON(obj interface{}) error {
	return c.ShouldBindWith(obj, binding.JSON)
}

// Read body data and restore request body data
func (c *Context) ShouldBindBodyJSON(obj interface{}) error {
	return c.ShouldBindBodyWith(obj, binding.JSON)
}

func (c *Context) GetBody() ([]byte, error) {
	body, err := c.Context.GetRawData()
	if err != nil {
		return nil, err
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
	return body, nil
}

func (c *Context) AddRenderHook(fn render.Hook) {
	c.RenderHook = append(c.RenderHook, fn)
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
	c.HttpStatus = status
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

func (c *Context) GetToken() string {
	return c.Request.Header.Get("Access-Token")
}

func (c *Context) RequestId() string {
	requestId := c.Request.Header.Get("X-Request-ID")
	if requestId == "" {
		requestId = uuid.New().String()
		c.Request.Header.Set("X-Request-ID", requestId)
	}

	return requestId
}

func (c *Context) SetLogInfo(key string, val interface{}) {
	c.LogInfo[key] = val
}

func (c *Context) RemoteIP() string {
	if ips := c.Request.Header.Get("X-Forwarded-For"); ips != "" {
		ipSli := strings.Split(ips, ",")
		for _, v := range ipSli {
			v = strings.TrimSpace(v)
			switch {
			case v == "":
				continue
			case v == "unknow":
				continue
			case v == "127.0.0.1":
				continue
			case strings.HasPrefix(v, "10."):
				continue
			case strings.HasPrefix(v, "172"):
				continue
			case strings.HasPrefix(v, "192"):
				continue
			}

			return v
		}
	} else if ip := c.Request.Header.Get("Client-Ip"); ip != "" {
		return strings.TrimSpace(ip)
	} else if ip := c.Request.Header.Get("Remote-Addr"); ip != "" {
		return strings.TrimSpace(ip)
	}

	return "-1"
}
