package main

import (
	"os"

	"github.com/orange-lightsaber/psb-rotatorc/cmd"
)

var (
	// VERSION is set during build
	VERSION = "0.1.0"
)

func main() {
	os.Exit(int(cmd.Exec(VERSION)))
}
