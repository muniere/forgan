package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type options struct {
	length    int
	start     int
	prefix    string
	dryRun    bool
	overwrite bool
	verbose   bool
}

func assemble(cmd *cobra.Command) {
	cmd.Flags().IntP("length", "l", 10, "length of new names")
	cmd.Flags().IntP("start", "s", 1, "start number of new names")
	cmd.Flags().StringP("prefix", "p", "", "prefix for new names")
	cmd.Flags().BoolP("dry-run", "n", false, "Do not perform action actually")
	cmd.Flags().Bool("overwrite", false, "overwrite existing files")
	cmd.Flags().BoolP("verbose", "v", false, "Show verbose messages")
}

func decode(flags *pflag.FlagSet) (*options, error) {
	length, err := flags.GetInt("length")
	if err != nil {
		return nil, err
	}

	start, err := flags.GetInt("start")
	if err != nil {
		return nil, err
	}

	prefix, err := flags.GetString("prefix")
	if err != nil {
		return nil, err
	}

	dryRun, err := flags.GetBool("dry-run")
	if err != nil {
		return nil, err
	}

	overwrite, err := flags.GetBool("overwrite")
	if err != nil {
		return nil, err
	}

	verbose, err := flags.GetBool("verbose")
	if err != nil {
		return nil, err
	}

	opts := &options{
		length:    length,
		start:     start,
		prefix:    prefix,
		dryRun:    dryRun,
		overwrite: overwrite,
		verbose:   verbose,
	}
	return opts, nil
}
