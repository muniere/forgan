package cli

import (
	"github.com/spf13/pflag"
)

type context struct {
	paths   []string
	options *options
}

func parse(args []string, flags *pflag.FlagSet) (*context, error) {
	paths, err := normalize(args)
	if err != nil {
		return nil, err
	}

	options, err := decode(flags)
	if err != nil {
		return nil, err
	}

	ctx := &context{
		paths:   paths,
		options: options,
	}
	return ctx, nil
}
