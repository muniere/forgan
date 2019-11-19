package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Options struct {
	Length    int
	Start     int
	Prefix    string
	DryRun    bool
	Overwrite bool
	Verbose   bool
}

func Prepare(cmd *cobra.Command) {
	cmd.Flags().IntP("length", "l", 10, "Length of new names")
	cmd.Flags().IntP("start", "s", 1, "Start number of new names")
	cmd.Flags().StringP("prefix", "p", "", "Prefix for new names")
	cmd.Flags().BoolP("dry-run", "n", false, "Do not perform action actually")
	cmd.Flags().Bool("overwrite", false, "Overwrite existing files")
	cmd.Flags().BoolP("verbose", "v", false, "Show verbose messages")
}

func Decode(flags *pflag.FlagSet) (*Options, error) {
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

	opts := &Options{
		Length:    length,
		Start:     start,
		Prefix:    prefix,
		DryRun:    dryRun,
		Overwrite: overwrite,
		Verbose:   verbose,
	}
	return opts, nil
}
