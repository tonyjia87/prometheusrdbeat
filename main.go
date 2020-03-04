package main

import (
	"os"

	"github.com/tonyjia87/prometheusrdbeat/cmd"

	_ "github.com/tonyjia87/prometheusrdbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
