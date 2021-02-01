package gee

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDebugRoute(t *testing.T) {
	router := Default()
	DebugRoute(router)

	{
		w := performRequest(router, "GET", "/debug/pprof/?debug_token=4821726c1947cdf3eebacade98173939")
		assert.Equal(t, http.StatusOK, w.Code)
		body, _ := ioutil.ReadAll(w.Body)
		assert.Contains(t, string(body), "<body>")
	}

	{
		w := performRequest(router, "GET", "/debug/metrics?debug_token=4821726c1947cdf3eebacade98173939")
		assert.Equal(t, http.StatusOK, w.Code)
		body, _ := ioutil.ReadAll(w.Body)
		fmt.Println(string(body))
		assert.Contains(t, string(body), "go_info")
	}
}

func TestDebugRouteWithPath(t *testing.T) {
	router := Default()
	gin.SetMode(gin.DebugMode)

	api := router.Group("/example")
	DebugRoute(router, api.BasePath())

	{
		w := performRequest(router, "GET", "/example/debug/pprof/?debug_token=4821726c1947cdf3eebacade98173939")
		assert.Equal(t, http.StatusOK, w.Code)
		body, _ := ioutil.ReadAll(w.Body)
		assert.Contains(t, string(body), "<body>")
	}

	{
		w := performRequest(router, "GET", "/example/debug/metrics?debug_token=4821726c1947cdf3eebacade98173939")
		assert.Equal(t, http.StatusOK, w.Code)
		body, _ := ioutil.ReadAll(w.Body)
		fmt.Println(string(body))
		assert.Contains(t, string(body), "go_info")
	}
}

func TestDebugWithRouteGroup(t *testing.T) {
	router := Default()
	gin.SetMode(gin.DebugMode)

	api := router.Group("/example")

	DebugWithRouteGroup(api)

	{
		w := performRequest(router, "GET", "/example/debug/pprof/?debug_token=4821726c1947cdf3eebacade98173939")
		assert.Equal(t, http.StatusOK, w.Code)
		body, _ := ioutil.ReadAll(w.Body)
		assert.Contains(t, string(body), "<body>")
	}

	{
		w := performRequest(router, "GET", "/example/debug/metrics?debug_token=4821726c1947cdf3eebacade98173939")
		assert.Equal(t, http.StatusOK, w.Code)
		body, _ := ioutil.ReadAll(w.Body)
		fmt.Println(string(body))
		assert.Contains(t, string(body), "go_info")
	}
}
