package render

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderString(t *testing.T) {
	w := httptest.NewRecorder()

	(String{
		Format: "hello %s %d",
		Data:   []interface{}{},
	}).WriteContentType(w)
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))

	err := (String{
		Format: "hola %s %d",
		Data:   []interface{}{"manu", 2},
	}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, "hola manu 2", w.Body.String())
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))

	rd := String{
		Format: "hola %s %d",
		Data:   []interface{}{"manu", 2},
	}
	hooks := make([]Hook, 0)
	hooks = append(hooks, func(body []byte) {
		assert.Equal(t, "hola manu 2", string(body))
	})
	rd.Hooks(hooks)
	err = rd.Render(w)
	assert.NoError(t, err)
}
