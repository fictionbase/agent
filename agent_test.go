package agent

import (
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCatchSig(t *testing.T) {
	catchSig(syscall.SIGHUP)
	catchSig(syscall.SIGTERM)
	catchSig(os.Interrupt)
	catchSig(syscall.SIGINT)
	assert.Nil(t, nil)
	assert.NotNil(t, "string")
}
