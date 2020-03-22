package gee

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var showdownFns []func(os.Signal)
var l = sync.Mutex{}

func WaitSignal(waitCh chan struct{}, stopSignals chan os.Signal, sigs ...os.Signal) {
	// stopSignals := make(chan os.Signal, 1)
	if len(sigs) == 0 {
		signal.Notify(stopSignals, syscall.SIGINT, syscall.SIGTERM)
	} else {
		signal.Notify(stopSignals, sigs...)
	}
	// 正常关闭和非正常关闭
	sig := <-stopSignals
	l.Lock()
	defer l.Unlock()
	appCancel()
	for i := len(showdownFns) - 1; i >= 0; i-- {
		showdownFns[i](sig)
	}
	waitCh <- struct{}{}
}

func RegisterShutDown(fn func(os.Signal)) {
	l.Lock()
	defer l.Unlock()
	showdownFns = append(showdownFns, fn)
}
