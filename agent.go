package agent

import (
	"fmt"
	"os"
	"syscall"
)

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
