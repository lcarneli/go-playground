package main

import (
	"github.com/lcarneli/go-playground/vault/cmd/cli"
)

func main() {
	if err := cli.RootCommand.Execute(); err != nil {
		return
	}
}
