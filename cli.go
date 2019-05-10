package agent

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/fictionbase/fictionbase/type/fbprocess"
	"github.com/fictionbase/fictionbase/type/fbresource"
)

// RunCLI exec Cli
func RunCLI(i []interface{}) int {
	// @TODO Exec Config Set
	go fbresourceRun()
	go fbprocessRun()
	// signal Catch
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGTERM, os.Interrupt)
	select {
	case ch := <-signalCh:
		return catchSig(ch)
	}
}

func fbresourceRun() {
	// @TODO determine whether to execute from fictionbase.yml
	resource := fbresource.FictionBase{}
	resource.Run()
}

func fbprocessRun() {
	// @TODO determine whether to execute from fictionbase.yml
	processes := fbprocess.FictionBase{}
	processes.Run()
}
