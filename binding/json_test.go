package binding

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONBindingBindBody(t *testing.T) {
	var s struct {
		Foo string `json:"foo"`
	}
	bind := jsonBinding{}
	err := bind.BindBody([]byte(`{"foo": "FOO"}`), &s)
	require.NoError(t, err)
	assert.Equal(t, "FOO", s.Foo)
	assert.Equal(t, "json", bind.Name())
}

func TestJsonBinding_Bind(t *testing.T) {
	t.Run("bind request", func(t *testing.T) {
		var s struct {
			Foo string `json:"foo"`
		}
		req, _ := http.NewRequest("POST", "/", strings.NewReader(`{"foo": "FOO"}`))
		bind := jsonBinding{}
		err := bind.Bind(req, &s)
		require.NoError(t, err)
		assert.Equal(t, "FOO", s.Foo)
	})

	t.Run("bind request error", func(t *testing.T) {
		var s struct {
			Foo string `json:"foo"`
		}
		bind := jsonBinding{}
		err := bind.Bind(nil, &s)
		require.Error(t, err)
	})
}

func TestJsonBindingError_Error(t *testing.T) {
	var s struct {
		Foo string `json:"foo"`
	}
	err := jsonBinding{}.BindBody([]byte(`foo,bar`), &s)
	require.Error(t, err)
	require.NotEmpty(t, err.Error())
	require.IsType(t, &JsonBindingError{}, err)
}
