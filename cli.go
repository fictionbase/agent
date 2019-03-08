package agent

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/fictionbase/fictionbase/type/fbresource"
)

type fictionBase interface {
	Run()
}

// RunCLI exec Cli
func RunCLI(i []interface{}) int {
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
	return 0
}
