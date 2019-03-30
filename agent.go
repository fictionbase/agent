package agent

import (
	"os"
	"syscall"

	"go.uber.org/zap"
)

func catchSig(sig os.Signal) int {

	// @TODO set config from fictionbase.yml
	logger, _ := zap.NewProduction()
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
