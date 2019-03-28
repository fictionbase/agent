package agent

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fictionbase/fictionbase/type/fbprocess"
	"github.com/fictionbase/fictionbase/type/fbresource"
)

type MessageBase struct {
	TypeKey    string    `json:"type_key"`
	StorageKey string    `json:"storage_key"`
	TimeKey    time.Time `json:"time_key"`
}

// RunCLI exec Cli
func RunCLI(i []interface{}) int {
	resource := fbresource.FictionBase{}
	resource.Run()
	processes := fbprocess.FictionBase{}
	processes.Run()
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGTERM, os.Interrupt)
	select {
	case ch := <-signalCh:
		catchSig(ch)
	}
	return 0
}
