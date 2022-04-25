package binding

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin/binding"
	"github.com/goapt/gee/encoding"
	"github.com/goapt/gee/encoding/json"
)

type JsonBindingError struct {
	Err error
}

func (j *JsonBindingError) Error() string {
	return j.Err.Error()
}

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (jsonBinding) Bind(req *http.Request, obj interface{}) error {
	if req == nil || req.Body == nil {
		return fmt.Errorf("invalid request")
	}
	return decodeJSON(req.Body, obj)
}

func (jsonBinding) BindBody(body []byte, obj interface{}) error {
	return decodeJSON(bytes.NewReader(body), obj)
}

func decodeJSON(r io.Reader, obj interface{}) error {
	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	if err := encoding.GetCodec(json.Name).Unmarshal(body, obj); err != nil {
		return &JsonBindingError{
			Err: err,
		}
	}
	return validate(obj)
}

func validate(obj interface{}) error {
	if binding.Validator == nil {
		return nil
	}
	return binding.Validator.ValidateStruct(obj)
}

var JSON = jsonBinding{}
