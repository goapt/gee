package gee

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
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
			c, _ := gin.CreateTestContext(w)

			ctx := &Context{
				Context:    c,
				HttpStatus: 200,
			}

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
