package gee

import (
	"net/http"
	"path"
)

// Router defines all router handle interface.
type Router interface {
	Use(middlewares ...Middleware)
	Any(pattern string, h http.Handler)
	Get(pattern string, h http.Handler)
	Post(pattern string, h http.Handler)
	Delete(pattern string, h http.Handler)
	Patch(pattern string, h http.Handler)
	Put(pattern string, h http.Handler)
	Options(pattern string, h http.Handler)
	Head(pattern string, h http.Handler)
	Group(pattern string) Router
	Handle(pattern string, h http.Handler)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

var _ Router = (*Route)(nil)

type Route struct {
	mux             *http.ServeMux
	middlewares     []Middleware
	groupPath       string
	handler         http.Handler
	notFoundHandler http.Handler
}

func NewRouter() *Route {
	return &Route{
		mux: http.NewServeMux(),
	}
}

func (r *Route) Handle(pattern string, h http.Handler) {
	r.handler = chain(r.mux, r.middlewares)
	r.mux.Handle(pattern, h)
}

func (r *Route) Use(h ...Middleware) {
	r.middlewares = append(r.middlewares, h...)
}

func (r *Route) Any(pattern string, h http.Handler) {
	r.Get(pattern, h)
	r.Post(pattern, h)
	r.Delete(pattern, h)
	r.Patch(pattern, h)
	r.Put(pattern, h)
	r.Options(pattern, h)
	r.Head(pattern, h)
}

func (r *Route) Get(pattern string, h http.Handler) {
	r.Handle(r.getPattern("GET", pattern), h)
}

func (r *Route) Post(pattern string, h http.Handler) {
	r.Handle(r.getPattern("POST", pattern), h)
}

func (r *Route) Delete(pattern string, h http.Handler) {
	r.Handle(r.getPattern("DELETE", pattern), h)
}

func (r *Route) Patch(pattern string, h http.Handler) {
	r.Handle(r.getPattern("PATCH", pattern), h)
}

func (r *Route) Put(pattern string, h http.Handler) {
	r.mux.Handle(r.getPattern("PUT", pattern), h)
}

func (r *Route) Options(pattern string, h http.Handler) {
	r.mux.Handle(r.getPattern("OPTIONS", pattern), h)
}

func (r *Route) Head(pattern string, h http.Handler) {
	r.Handle(r.getPattern("HEAD", pattern), h)
}

func (r *Route) Group(pattern string) Router {
	return &Route{mux: r.mux, middlewares: r.middlewares, groupPath: pattern, handler: r.handler, notFoundHandler: r.notFoundHandler}
}

func (r *Route) NotFound(h http.Handler) {
	r.notFoundHandler = h
}

func (r *Route) getPattern(method string, pattern string) string {
	if r.groupPath == "" {
		return method + " " + pattern
	}
	return method + " " + path.Join(r.groupPath, pattern)
}

func (r *Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r.handler == nil {
		if r.notFoundHandler != nil {
			r.notFoundHandler.ServeHTTP(w, req)
			return
		}
		http.NotFound(w, req)
		return
	}

	r.handler.ServeHTTP(NewWrapResponseWriter(w), req)
}
