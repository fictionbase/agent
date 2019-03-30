package main

import (
	"os"

	"github.com/fictionbase/agent"
)

//@TODO Add Run Info

func main() {
	i := []interface{}{}
	os.Exit(agent.RunCLI(i))
}
