package cli

import (
	"github.com/spf13/pflag"
)

type Context struct {
	Paths   []string
	Options *Options
}

func Parse(args []string, flags *pflag.FlagSet) (*Context, error) {
	paths, err := Normalize(args)
	if err != nil {
		return nil, err
	}

	options, err := Decode(flags)
	if err != nil {
		return nil, err
	}

	ctx := &Context{
		Paths:   paths,
		Options: options,
	}
	return ctx, nil
}
