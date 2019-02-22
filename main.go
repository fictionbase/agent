package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fictionbase/fictionbase/type/fbresource"
)

type fictionBase interface {
	Run()
}

func main() {
	// @TODO fictionBase return by config want Run
	var fictionBase []fbresource.FictionBase
	resource := fbresource.FictionBase{}
	fictionBase = append(fictionBase, resource)
	// loop goroutine on background
	for _, fiction := range fictionBase {
		go fiction.Run()
	}

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGTERM, os.Interrupt)
	select {
	case ch := <-signalCh:
		catchSig(ch)
	}
}

func catchSig(sig os.Signal) {
	// @TODO change print zap
	switch sig {
	case syscall.SIGHUP:
		fmt.Println("SIGHUP Happend! fictionbase stop", sig)
	case syscall.SIGTERM:
		fmt.Println("SIGTERM Happend! fictionbase stop", sig)
	case os.Interrupt:
		fmt.Println("os.Interrupt Happend! fictionbase stop", sig)
	default:
		fmt.Println("singal Happend! fictionbase stop", sig)
	}
}
