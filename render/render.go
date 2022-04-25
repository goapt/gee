package render

import (
	"net/http"

	"github.com/gin-gonic/gin/render"
)

var (
	_ render.Render = (*JSON)(nil)
)

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
