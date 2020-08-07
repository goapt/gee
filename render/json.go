package render

import (
	"net/http"

	"github.com/goapt/gee/internal/json"
)

type JSON struct {
	Data     interface{}
	hookFunc []Hook
}

var jsonContentType = []string{"application/json; charset=utf-8"}

func (r JSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

func (r JSON) WriteJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	jsonBytes, err := json.Marshal(obj)
	// 写write之前先调用hook
	r.runHooks(jsonBytes)

	_, err = w.Write(jsonBytes)
	return err
}

func (r JSON) Render(w http.ResponseWriter) (err error) {
	r.WriteContentType(w)

	if err = r.WriteJSON(w, r.Data); err != nil {
		panic(err)
	}

	return
}

func (r *JSON) Hooks(hooks []Hook) {
	r.hookFunc = hooks
}

func (r *JSON) runHooks(body []byte) {
	for _, fn := range r.hookFunc {
		fn(body)
	}
}
