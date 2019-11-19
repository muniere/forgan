package main

import (
	"github.com/muniere/forgan/internal/app/count"
	"github.com/muniere/forgan/internal/pkg/sys"
)

func main() {
	cmd := count.NewCommand()
	err := cmd.Execute()
	sys.CheckError(err)
}
