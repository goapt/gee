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

var authDebug = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("token") != debugToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
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
	g := r.Group(p)
	g.Get("/", http.HandlerFunc(pprof.Index))
	g.Get("/cmdline", http.HandlerFunc(pprof.Cmdline))
	g.Get("/profile", http.HandlerFunc(pprof.Profile))
	g.Post("/symbol", http.HandlerFunc(pprof.Symbol))
	g.Get("/symbol", http.HandlerFunc(pprof.Symbol))
	g.Get("/trace", http.HandlerFunc(pprof.Trace))
	g.Handle("/allocs", pprof.Handler("allocs"))
	g.Handle("/block", pprof.Handler("block"))
	g.Handle("/goroutine", pprof.Handler("goroutine"))
	g.Handle("/heap", pprof.Handler("heap"))
	g.Handle("/mutex", pprof.Handler("mutex"))
	g.Handle("/threadcreate", pprof.Handler("threadcreate"))

	// prometheus
	g.Handle("/metrics", promhttp.Handler())
}
