package gee

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Middleware = func(http.Handler) http.Handler

// Router defines all router handle interface.
type Router interface {
	Use(middlewares ...Middleware) Router
	Any(pattern string, h http.HandlerFunc)
	Get(pattern string, h http.HandlerFunc)
	Post(pattern string, h http.HandlerFunc)
	Delete(pattern string, h http.HandlerFunc)
	Patch(pattern string, h http.HandlerFunc)
	Put(pattern string, h http.HandlerFunc)
	Options(pattern string, h http.HandlerFunc)
	Head(pattern string, h http.HandlerFunc)
	Group(pattern string) Router
	Handle(pattern string, h http.Handler)
}

var _ Router = (*Route)(nil)

type Route struct {
	chi chi.Router
}

func NewRouter() *Route {
	return &Route{chi: chi.NewRouter()}
}

func (r *Route) Handle(pattern string, h http.Handler) {
	r.chi.Handle(pattern, h)
}

func (r *Route) Use(middlewares ...Middleware) Router {
	rt := r.chi.With(middlewares...)
	return &Route{chi: rt}
}

func (r *Route) Any(pattern string, h http.HandlerFunc) {
	r.chi.HandleFunc(pattern, h)
}

func (r *Route) Get(pattern string, h http.HandlerFunc) {
	r.chi.Get(pattern, h)
}

func (r *Route) Post(pattern string, h http.HandlerFunc) {
	r.chi.Post(pattern, h)
}

func (r *Route) Delete(pattern string, h http.HandlerFunc) {
	r.chi.Delete(pattern, h)
}

func (r *Route) Patch(pattern string, h http.HandlerFunc) {
	r.chi.Patch(pattern, h)
}

func (r *Route) Put(pattern string, h http.HandlerFunc) {
	r.chi.Put(pattern, h)
}

func (r *Route) Options(pattern string, h http.HandlerFunc) {
	r.chi.Options(pattern, h)
}

func (r *Route) Head(pattern string, h http.HandlerFunc) {
	r.chi.Head(pattern, h)
}

func (r *Route) Group(pattern string) Router {
	rt := r.chi.Route(pattern, func(rt chi.Router) {})
	return &Route{chi: rt}
}

func (r *Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.chi.ServeHTTP(w, req)
}
