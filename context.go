package very

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/verystar/golib/convert"
	"strings"
	"time"
)

var (
	SuccessCode int = 10000
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
	Session    ISession
	StartTime  time.Time
}

func (c *Context) Status(status int) {
	c.HttpStatus = status
}

func (c *Context) Fail(code int, msg interface{}) Response {
	var message string
	if m, ok := msg.(error); ok {
		message = m.Error()
	} else {
		message = convert.ToStr(msg)
	}

	return &ApiResponse{
		HttpStatus: getHttpStatus(c, 200),
		Context:    c.Context,
		Code:       code,
		Msg:        message,
	}
}

func (c *Context) Success(data interface{}) Response {
	return &ApiResponse{
		HttpStatus: getHttpStatus(c, 200),
		Context:    c.Context,
		Code:       SuccessCode,
		Data:       data,
		Msg:        "ok",
	}
}

func (c *Context) JSON(data interface{}) Response {
	return &JSONResponse{
		HttpStatus: getHttpStatus(c, 200),
		Context:    c.Context,
		Data:       data,
	}
}

func (c *Context) Redirect(location string) Response {
	return &RedirectResponse{
		HttpStatus: getHttpStatus(c, 302),
		Context:    c.Context,
		Location:   location,
	}
}

func (c *Context) String(format string, values ...interface{}) Response {
	return &StringResponse{
		HttpStatus: getHttpStatus(c, 200),
		Context:    c.Context,
		Name:       format,
		Data:       values,
	}
}

func (c *Context) GetToken() string {
	return c.Request.Header.Get("Access-Token")
}

func (c *Context) BusinessError(msg interface{}) Response {
	return c.Fail(40002, msg)
}
func (c *Context) SystemError(msg interface{}) Response {
	return c.Fail(40003, msg)
}

func (c *Context) RequestId() string {
	requestId := c.Request.Header.Get("X-Request-ID")
	if requestId == "" {
		requestId = uuid.New().String()
		c.Request.Header.Set("X-Request-ID", requestId)
	}

	return requestId
}

func (c *Context) RemoteIP() string {
	var ip string
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
	} else if ip = c.Request.Header.Get("Client-Ip"); ip != "" {
		return strings.TrimSpace(ip)
	} else if ip = c.Request.Header.Get("Remote-Addr"); ip != "" {
		return strings.TrimSpace(ip)
	}

	if ip != "" {
		return ip
	}

	return "-1"
}
