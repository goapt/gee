package gee

import (
	"net/http"
	"net/http/pprof"
	"path/filepath"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var debugToken = "058759a8c4d920576348854616a58f3f"

func SetDebugToken(token string) {
	debugToken = token
}

func authDebug(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("token") != debugToken {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func DebugRoute(r Router, path ...string) {
	p := "/debug"
	// custom router path
	if len(path) > 0 && path[0] != "" {
		p = filepath.Join(path[0], p)
	}
	r.Use(authDebug)
	// pporf
	r.Group(p, func(r Router) {
		r.Get("/", pprof.Index)
		r.Get("/cmdline", pprof.Cmdline)
		r.Get("/profile", pprof.Profile)
		r.Post("/symbol", pprof.Symbol)
		r.Get("/symbol", pprof.Symbol)
		r.Get("/trace", pprof.Trace)
		r.Handle("/allocs", pprof.Handler("allocs"))
		r.Handle("/block", pprof.Handler("block"))
		r.Handle("/goroutine", pprof.Handler("goroutine"))
		r.Handle("/heap", pprof.Handler("heap"))
		r.Handle("/mutex", pprof.Handler("mutex"))
		r.Handle("/threadcreate", pprof.Handler("threadcreate"))
	})

	// prometheus
	r.Group(p, func(r Router) {
		r.Handle("/metrics", promhttp.Handler())
	})
}
