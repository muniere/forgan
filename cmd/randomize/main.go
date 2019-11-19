package main

import (
	"github.com/muniere/forgan/internal/app/randomize"
	"github.com/muniere/forgan/internal/pkg/sys"
)

func main() {
	cmd := randomize.NewCommand()
	err := cmd.Execute()
	sys.CheckError(err)
}
