package main

import (
	"os"

	"github.com/fictionbase/agent"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	i := []interface{}{}
	os.Exit(agent.RunCLI(i))
}
