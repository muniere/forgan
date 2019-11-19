package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Options struct {
	IncludesHidden bool
	Pattern        string
	Verbose        bool
}

func Prepare(cmd *cobra.Command) {
	cmd.Flags().BoolP("all", "a", false, "Include hidden files or directories")
	cmd.Flags().BoolP("verbose", "v", false, "Show verbose log messages")
	cmd.Flags().StringP("pattern", "p", "", "Pattern to filter files")
}

func Decode(flags *pflag.FlagSet) (*Options, error) {
	includesHidden, err := flags.GetBool("all")
	if err != nil {
		return nil, err
	}

	pattern, err := flags.GetString("pattern")
	if err != nil {
		return nil, err
	}

	verbose, err := flags.GetBool("verbose")
	if err != nil {
		return nil, err
	}

	opts := &Options{
		IncludesHidden: includesHidden,
		Pattern:        pattern,
		Verbose:        verbose,
	}
	return opts, nil
}
