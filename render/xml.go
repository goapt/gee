package render

import (
	"encoding/xml"
	"net/http"
)

// XML contains the given interface object.
type XML struct {
	Data     interface{}
	hookFunc []Hook
}

var xmlContentType = []string{"application/xml; charset=utf-8"}

// Render (XML) encodes the given interface object and writes data with custom ContentType.
func (r XML) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	xmlBytes, err := xml.Marshal(r.Data)
	if err != nil {
		return err
	}
	// 写write之前先调用hook
	r.runHooks(xmlBytes)

	_, err = w.Write(xmlBytes)
	return err
}

// WriteContentType (XML) writes XML ContentType for response.
func (r XML) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, xmlContentType)
}

func (r *XML) Hooks(hooks []Hook) {
	r.hookFunc = hooks
}

func (r *XML) runHooks(body []byte) {
	for _, fn := range r.hookFunc {
		fn(body)
	}
}
