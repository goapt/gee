package gee

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestContext_ShouldBindJSON(t *testing.T) {
	ctx, _ := CreateTestContext(httptest.NewRecorder())
	body := `{"page":1,"num":10,"search":"你好","created":"2020-09-17 12:34:23"}`
	ctx.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(body))

	p := &struct {
		Page    int       `json:"page"`
		Num     int       `json:"num"`
		Search  string    `json:"search"`
		Created time.Time `json:"created" time_format:"2006-01-02 15:04:05"`
	}{}

	err := ctx.ShouldBindJSON(p)
	assert.NoError(t, err)
	assert.Equal(t, p.Page, 1)
	assert.Equal(t, p.Num, 10)
	assert.Equal(t, p.Created.Format("2006-01-02 15:04:05"), "2020-09-17 12:34:23")
	assert.Equal(t, p.Search, "你好")
}

func TestContext_BindJSON(t *testing.T) {
	ctx, _ := CreateTestContext(httptest.NewRecorder())
	body := `{"page":1,"num":10,"search":"你好"}`
	ctx.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(body))

	p := &struct {
		Page   int    `json:"page"`
		Num    int    `json:"num"`
		Search string `json:"search"`
	}{}
	err := ctx.BindJSON(p)
	assert.NoError(t, err)
	assert.Equal(t, p.Page, 1)
	assert.Equal(t, p.Num, 10)
	assert.Equal(t, p.Search, "你好")
}

func TestContext_ShouldBindBodyJSON(t *testing.T) {
	ctx, _ := CreateTestContext(httptest.NewRecorder())
	body := `{"page":1,"num":10,"search":"你好","created":"2020-09-17 12:34:23"}`
	ctx.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(body))
	p := &struct {
		Page    int       `json:"page"`
		Num     int       `json:"num"`
		Search  string    `json:"search"`
		Created time.Time `json:"created" time_format:"2006-01-02 15:04:05"`
	}{}
	err := ctx.ShouldBindBodyJSON(p)
	assert.NoError(t, err)
	assert.Equal(t, p.Page, 1)
	assert.Equal(t, p.Num, 10)
	assert.Equal(t, p.Search, "你好")
	assert.Equal(t, p.Created.Format("2006-01-02 15:04:05"), "2020-09-17 12:34:23")
	body2, _ := ctx.Get(gin.BodyBytesKey)
	assert.Equal(t, string(body2.([]byte)), body)
}

func TestContext_GetBody(t *testing.T) {
	ctx, _ := CreateTestContext(httptest.NewRecorder())
	body := `{"page":1,"num":10,"search":"你好"}`
	ctx.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(body))

	body2, err := ctx.GetBody()
	assert.NoError(t, err)
	assert.Equal(t, string(body2), body)
	body3, _ := ioutil.ReadAll(ctx.Request.Body)
	assert.Equal(t, string(body3), body)
}

func TestContext_ResponseWriter(t *testing.T) {
	ctx, _ := CreateTestContext(httptest.NewRecorder())
	resp := ctx.JSON(H{
		"foo": "bar",
	})
	resp.Render()
	assert.Equal(t, `{"foo":"bar"}`, string(ctx.ResponseBody()))
	assert.NotNil(t, ctx.ResponseWriter())
}

func TestContext_AddRenderHook(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := CreateTestContext(w)
	ctx.AddRenderHook(func(body []byte) {
		ctx.Writer.Header().Set("X-Response-Len", strconv.Itoa(len(body)))
	})

	resp := ctx.JSON(H{
		"foo": "bar",
	})
	resp.Render()
	assert.Equal(t, `{"foo":"bar"}`, string(ctx.ResponseBody()))
	assert.Equal(t, strconv.Itoa(len(ctx.ResponseBody())), w.Header().Get("X-Response-Len"))
}

func TestContext_JSON(t *testing.T) {
	type data struct {
		Id      int       `json:"id"`
		Name    string    `json:"name"`
		Created time.Time `json:"created" time_format:"2006-01-02 15:04:05"`
	}

	tests := []struct {
		name string
		args interface{}
		want string
	}{
		{name: "biz error", args: &data{
			Id:      1,
			Name:    "test",
			Created: time.Date(2019, 11, 21, 07, 49, 0, 0, time.Local),
		}, want: `{"id":1,"name":"test","created":"2019-11-21 07:49:00"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := CreateTestContext(w)

			resp := ctx.JSON(tt.args)
			resp.Render()
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

			if got := w.Body.String(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_String(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{name: "biz error", args: "foo bar", want: `foo bar`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := CreateTestContext(w)

			resp := ctx.String(tt.args)
			resp.Render()
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))

			if got := w.Body.String(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_XML(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := CreateTestContext(w)

	resp := ctx.XML(H{
		"foo": "bar",
	})
	resp.Render()
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "<map><foo>bar</foo></map>", w.Body.String())
	assert.Equal(t, "application/xml; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestContext_RequestId(t *testing.T) {
	t.Run("get header request id", func(t *testing.T) {
		ctx, _ := CreateTestContext(httptest.NewRecorder())
		ctx.Request, _ = http.NewRequest("POST", "/", nil)
		ctx.Request.Header.Set("X-Request-ID", "65ce669b-ed17-4255-8c0b-c287ed17a01e")
		assert.Equal(t, "65ce669b-ed17-4255-8c0b-c287ed17a01e", ctx.RequestId())
	})

	t.Run("set header request id", func(t *testing.T) {
		ctx, _ := CreateTestContext(httptest.NewRecorder())
		ctx.Request, _ = http.NewRequest("POST", "/", nil)
		assert.Len(t, ctx.RequestId(), 36)
	})
}

func TestContext_GetToken(t *testing.T) {
	t.Run("get access token", func(t *testing.T) {
		ctx, _ := CreateTestContext(httptest.NewRecorder())
		ctx.Request, _ = http.NewRequest("POST", "/", nil)
		ctx.Request.Header.Set("Access-Token", "c287ed17a01e")
		assert.Equal(t, "c287ed17a01e", ctx.Request.Header.Get("Access-Token"))
	})
}

func TestContext_Status(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := CreateTestContext(w)
	ctx.Status(403)

	resp := ctx.String("foo")
	resp.Render()
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestContext_Redirect(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := CreateTestContext(w)
	resp := ctx.Redirect("/new/location")
	assert.PanicsWithValue(t, "Cannot redirect with status code 200", func() {
		ctx.Status(http.StatusOK)
		resp.Render()
	})
}

func TestContext_SetLogInfo(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/", nil)
	ctx2 := context.WithValue(context.Background(), "__info__", map[string]interface{}{"foo": "bar"})
	ctx.Request = ctx.Request.WithContext(ctx2)

	assert.Equal(t, map[string]interface{}{"foo": "bar"}, ctx.Request.Context().Value("__info__"))
}
