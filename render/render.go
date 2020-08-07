package render

import "net/http"

type Hook func(body []byte)

// Render interface is to be implemented by JSON, XML, HTML, YAML and so on.
type Render interface {
	// Render writes data with custom ContentType.
	Render(http.ResponseWriter) error
	// WriteContentType writes custom ContentType.
	WriteContentType(w http.ResponseWriter)
	// Hooks set rendner hooks
	Hooks(hooks []Hook)
}

var (
	_ Render = (*JSON)(nil)
	_ Render = (*XML)(nil)
	_ Render = (*String)(nil)
)

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
