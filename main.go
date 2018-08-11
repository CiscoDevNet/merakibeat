package main

import (
	"os"

	"github.com/CiscoDevNet/merakibeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
