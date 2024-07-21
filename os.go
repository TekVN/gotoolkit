package gotoolkit

import (
	"os"
	"os/signal"
)

func RegisterOSSignalHandler(f func(), signals ...os.Signal) {
	if len(signals) == 0 {
		return
	}

	go WaitOSSignalHandler(f, signals...)
}

func WaitOSSignalHandler(f func(), signals ...os.Signal) {
	if len(signals) == 0 {
		return
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)
	<-ch

	f()
}
