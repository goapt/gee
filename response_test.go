package gee

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestJSONResponse_Render(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	resp := &JSONResponse{
		Context: &Context{
			Context: ctx,
		},
		Data: map[string]interface{}{
			"id":   1,
			"name": "test",
		},
	}

	resp.Render()
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"id":1,"name":"test"}`, w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestXMLResponse_Render(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	resp := &XMLResponse{
		Context: &Context{
			Context: ctx,
		},
		Data: H{
			"foo": "bar",
		},
	}

	resp.Render()
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `<map><foo>bar</foo></map>`, w.Body.String())
	assert.Equal(t, "application/xml; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestStringResponse_Render(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	resp := &StringResponse{
		Context: &Context{
			Context: ctx,
		},
		Format: "hola %s %d",
		Data:   []interface{}{"manu", 2},
	}

	resp.Render()
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hola manu 2", w.Body.String())
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestRedirectResponse_Render(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	resp := &RedirectResponse{
		Context: &Context{
			Context: ctx,
		},
		Location: "/new/location",
	}

	w = httptest.NewRecorder()
	assert.PanicsWithValue(t, "Cannot redirect with status code 200", func() {
		resp.Context.httpStatus = http.StatusOK
		resp.Render()
	})
}
