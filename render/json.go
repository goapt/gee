package render

import (
	"net/http"

	"github.com/goapt/gee/encoding"
	"github.com/goapt/gee/encoding/json"
)

type JSON struct {
	Data interface{}
}

var jsonContentType = []string{"application/json; charset=utf-8"}

func (r JSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

func (r JSON) WriteJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	jsonBytes, err := encoding.GetCodec(json.Name).Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

func (r JSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	if err := r.WriteJSON(w, r.Data); err != nil {
		return err
	}
	return nil
}
