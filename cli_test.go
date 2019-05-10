package agent

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	i := []interface{}{}
	go RunCLI(i)
	go fbresourceRun()
	go fbprocessRun()
	time.Sleep(2 * time.Second)
}
