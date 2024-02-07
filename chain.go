package gee

import "net/http"

type Middleware = func(next http.Handler) http.Handler

// chain builds a http.Handler composed of an inline middleware stack and endpoint
// handler in the order they are passed.
func chain(endpoint http.Handler, middlewares []Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		endpoint = middlewares[i](endpoint)
	}
	return endpoint
}
