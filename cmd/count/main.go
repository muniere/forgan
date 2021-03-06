package main

import (
	"github.com/muniere/forgan/internal/app/count/cli"
	"github.com/muniere/forgan/internal/pkg/sys"
)

func main() {
	cmd := cli.NewCommand()
	err := cmd.Execute()
	sys.CheckError(err)
}
