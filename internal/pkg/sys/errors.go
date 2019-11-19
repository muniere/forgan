package sys

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, "%s: %v", os.Args[0], err)
	os.Exit(1)
}