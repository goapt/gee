package gee

import (
	"context"
	"fmt"
	"net/http"
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
	r.Get("/", fn)
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
