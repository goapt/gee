// Copyright 2017 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"net/http"

	"github.com/goapt/gee/encoding"
	"github.com/goapt/gee/encoding/form"
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) Bind(req *http.Request, obj any) error {
	values := req.URL.Query()
	return encoding.GetCodec(form.Name).Unmarshal([]byte(values.Encode()), obj)
}

var Query = queryBinding{}
