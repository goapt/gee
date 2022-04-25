package gee

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithTimeout(t *testing.T) {
	apply := WithTimeout(time.Second)
	c := &clientOptions{}
	apply(c)
	assert.Equal(t, c.timeout, time.Second)
}
