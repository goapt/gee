// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package render

import (
	"fmt"
	"io"
	"net/http"
)

// String contains the given interface object slice and its format.
type String struct {
	Format   string
	Data     []interface{}
	hookFunc []Hook
}

var plainContentType = []string{"text/plain; charset=utf-8"}

// Render (String) writes data with custom ContentType.
func (r String) Render(w http.ResponseWriter) error {
	return r.WriteString(w, r.Format, r.Data)
}

// WriteContentType (String) writes Plain ContentType.
func (r String) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, plainContentType)
}

// WriteString writes data according to its format and write custom ContentType.
func (r String) WriteString(w http.ResponseWriter, format string, data []interface{}) (err error) {
	writeContentType(w, plainContentType)
	if len(data) > 0 {
		format = fmt.Sprintf(format, data...)
	}

	// 写write之前先调用hook
	r.runHooks([]byte(format))

	_, err = io.WriteString(w, format)
	return
}

func (r *String) Hooks(hooks []Hook) {
	r.hookFunc = hooks
}

func (r *String) runHooks(body []byte) {
	for _, fn := range r.hookFunc {
		fn(body)
	}
}
