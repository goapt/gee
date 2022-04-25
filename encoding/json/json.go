package json

import (
	"encoding/json"

	"github.com/goapt/gee/encoding"
	"github.com/ilibs/jsontime"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Name is the name registered for the json codec.
const Name = "json"

var (
	// MarshalOptions is a configurable JSON format marshaller.
	MarshalOptions = protojson.MarshalOptions{
		UseProtoNames: true,
	}

	// UnmarshalOptions is a configurable JSON format unmarshaller.
	UnmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	// EnableJsonTime is used to call the jsontime format method on the JSON
	EnableJsonTime = true
)

// codec is a Codec implementation with json.
type codec struct{}

func (codec) Marshal(obj interface{}) ([]byte, error) {
	var (
		jsonBytes []byte
		err       error
	)

	switch m := obj.(type) {
	case json.Marshaler:
		jsonBytes, err = m.MarshalJSON()
	case proto.Message:
		jsonBytes, err = MarshalOptions.Marshal(m)
	default:
		if EnableJsonTime {
			jsonBytes, err = jsontime.ConfigWithCustomTimeFormat.Marshal(m)
		} else {
			jsonBytes, err = json.Marshal(m)
		}
	}
	if err != nil {
		return nil, err
	}
	return jsonBytes, err
}

func (codec) Unmarshal(data []byte, obj interface{}) error {
	switch m := obj.(type) {
	case json.Unmarshaler:
		return m.UnmarshalJSON(data)
	case proto.Message:
		return UnmarshalOptions.Unmarshal(data, m)
	default:
		if EnableJsonTime {
			return jsontime.ConfigWithCustomTimeFormat.Unmarshal(data, obj)
		}
		return json.Unmarshal(data, obj)
	}
}

func (codec) Name() string {
	return Name
}

func init() {
	encoding.RegisterCodec(codec{})
}
