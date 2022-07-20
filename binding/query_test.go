package binding

import (
	"net/http"
	"testing"

	"github.com/goapt/gee/internal/testdata/binding"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_queryBindingBindProto(t *testing.T) {
	m := &binding.HelloRequest{}
	bind := queryBinding{}
	req, err := http.NewRequest("GET", "/?name=foo&sub.name=bar", nil)
	assert.NoError(t, err)
	err = bind.Bind(req, m)
	require.NoError(t, err)
	assert.Equal(t, "foo", m.Name)
	assert.Equal(t, "bar", m.Sub.Name)
}
