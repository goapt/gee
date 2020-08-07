package gee

import (
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterShutDown(t *testing.T) {
	RegisterShutDown(func(signal os.Signal) {
		assert.Equal(t, signal, syscall.SIGINT)
	})

	registerLastShutDown(func(signal os.Signal) {
		assert.Equal(t, signal, syscall.SIGINT)
	})

	waitCh := make(chan struct{})
	stopSignals := make(chan os.Signal, 1)
	go WaitSignal(waitCh, stopSignals)
	stopSignals <- syscall.SIGINT
	<-waitCh
}
