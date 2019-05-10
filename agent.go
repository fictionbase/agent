package agent

import (
	"os"
	"syscall"

	"github.com/fictionbase/fictionbase"
)

func catchSig(sig os.Signal) int {

	logger := fictionbase.GetLogger()

	switch sig {
	case syscall.SIGHUP:
		logger.Info("SIGHUP Happend! fictionbase stop")
	case syscall.SIGTERM:
		logger.Info("SIGTERM Happend! fictionbase stop")
	case os.Interrupt:
		logger.Info("os.Interrupt Happend! fictionbase stop")
	default:
		logger.Info("signal Happend! fictionbase stop")
		return 1
	}
	return 0
}
