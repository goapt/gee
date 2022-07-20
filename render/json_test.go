package render

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRenderJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := &struct {
		Id      int
		Name    string
		Created time.Time `time_format:"2006-01-02 15:04:05"`
	}{
		1,
		"test",
		time.Date(2020, 1, 1, 11, 11, 11, 0, time.UTC),
	}

	(JSON{Data: data}).WriteContentType(w)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	err := (JSON{Data: data}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, `{"Id":1,"Name":"test","Created":"2020-01-01T11:11:11Z"}`, w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}
