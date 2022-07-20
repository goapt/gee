package binding

import (
	"net/url"
	"testing"

	"github.com/goapt/gee/internal/testdata/binding"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_uriBinding_BindUri(t *testing.T) {
	m := &binding.HelloRequest{}
	bind := uriBinding{}

	vars := url.Values{}
	vars.Add("name", "foo")
	vars.Add("sub.name", "bar")

	err := bind.BindUri(vars, m)
	require.NoError(t, err)
	assert.Equal(t, "foo", m.Name)
	assert.Equal(t, "bar", m.Sub.Name)
}
