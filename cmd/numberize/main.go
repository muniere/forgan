package main

import (
	"github.com/muniere/forgan/internal/app/numberize"
	"github.com/muniere/forgan/internal/pkg/sys"
)

func main() {
	cmd := numberize.NewCommand()
	err := cmd.Execute()
	sys.CheckError(err)
}
