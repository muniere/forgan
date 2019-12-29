package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type options struct {
	includesHidden bool
	pattern        string
	verbose        bool
}

func assemble(cmd *cobra.Command) {
	cmd.Flags().BoolP("all", "a", false, "Include hidden files or directories")
	cmd.Flags().BoolP("verbose", "v", false, "Show verbose log messages")
	cmd.Flags().StringP("pattern", "p", "", "pattern to filter files")
}

func decode(flags *pflag.FlagSet) (*options, error) {
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

	opts := &options{
		includesHidden: includesHidden,
		pattern:        pattern,
		verbose:        verbose,
	}
	return opts, nil
}
