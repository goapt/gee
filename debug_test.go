package gee

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebugRoute(t *testing.T) {

	router := Default()
	DebugRoute(router.Engine)
	w := performRequest(router, "GET", "/debug/info?debug_token=4821726c1947cdf3eebacade98173939")
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, "GET", "/debug/pprof/?debug_token=4821726c1947cdf3eebacade98173939")
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, "GET", "/debug/pprof/heap?debug_token=4821726c1947cdf3eebacade98173939")
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, "GET", "/debug/pprof/goroutine?debug_token=4821726c1947cdf3eebacade98173939")
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, "GET", "/debug/pprof/block?debug_token=4821726c1947cdf3eebacade98173939")
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, "GET", "/debug/pprof/threadcreate?debug_token=4821726c1947cdf3eebacade98173939")
	assert.Equal(t, http.StatusOK, w.Code)
}
