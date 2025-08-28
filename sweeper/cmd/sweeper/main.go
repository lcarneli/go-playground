package main

import (
	"github.com/lcarneli/go-playground/sweeper/cmd/cli"
	"os"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
