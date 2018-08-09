package main

import (
	"os"

	"github.com/npateriya/merakibeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
