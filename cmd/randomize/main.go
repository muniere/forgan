package main

import (
	"github.com/muniere/forgan/internal/app/randomize/cli"
	"github.com/muniere/forgan/internal/pkg/sys"
)

func main() {
	cmd := cli.NewCommand()
	err := cmd.Execute()
	sys.CheckError(err)
}
