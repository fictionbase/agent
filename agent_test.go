package agent

import (
	"os"
	"syscall"
	"testing"
)

func TestCatchSig(t *testing.T) {
	catchSig(syscall.SIGHUP)
	catchSig(syscall.SIGTERM)
	catchSig(os.Interrupt)
	catchSig(syscall.SIGINT)
}
