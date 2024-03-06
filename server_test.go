package gee

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	}
	r := NewRouter()
	r.Get("/", http.HandlerFunc(fn))
	t.Run("cancel", func(t *testing.T) {
		srv := New(
			Address(":8888"),
			Handler(r),
			StopTimeout(3*time.Second),
		)
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(1 * time.Second)
			cancel()
		}()
		err := srv.Start(ctx)
		assert.NoError(t, err)
	})

	t.Run("kill", func(t *testing.T) {
		srv := New(
			Address(":8888"),
			Handler(r),
			StopTimeout(1*time.Second),
		)
		go func() {
			time.Sleep(1 * time.Second)
			_ = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
		}()
		err := srv.Start(context.Background())
		assert.NoError(t, err)
	})
}

func TestNew2(t *testing.T) {
	md1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Log(1)
			defer func() {
				t.Log(4)
				body := Response(w).Body()
				t.Log(string(body))
			}()
			next.ServeHTTP(w, r)
			t.Log(3)
		})
	}

	r := NewRouter()
	r.Use(md1)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t.Log(2)
		w.Write([]byte("123"))
	})

	r2 := r.Group("/api")

	r2.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Log("group")
			next.ServeHTTP(w, r)
		})
	})

	r2.Get("/test", func(writer http.ResponseWriter, request *http.Request) {
		t.Log(5)
		writer.Write([]byte("test"))
	})

	s := httptest.NewServer(r)
	http.Get(s.URL)
	s2 := httptest.NewServer(r2)
	http.Get(s2.URL + "/api/test")
}
