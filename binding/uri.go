// Copyright 2017 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"net/url"

	"github.com/goapt/gee/encoding"
	"github.com/goapt/gee/encoding/form"
)

type uriBinding struct{}

func (uriBinding) Name() string {
	return "query"
}

func (uriBinding) BindUri(values url.Values, obj interface{}) error {
	return encoding.GetCodec(form.Name).Unmarshal([]byte(values.Encode()), obj)
}

var Uri = uriBinding{}
