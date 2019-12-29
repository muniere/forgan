package cli

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/muniere/forgan/internal/app/numberize/exe"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "numberize",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd, args)
		},
	}

	assemble(cmd)

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	ctx, err := parse(args, cmd.Flags())
	if err != nil {
		return err
	}

	if err := prepare(ctx.options); err != nil {
		return err
	}

	opts, err := translate(ctx.options)
	if err != nil {
		return err
	}

	if err := perform(ctx.paths, opts); err != nil {
		log.Warnf("%v", err)
	}

	return nil
}

func prepare(options *options) error {
	if options.verbose {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.SetOutput(os.Stderr)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    false,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "15:04:05.000",
	})

	return nil
}

func translate(options *options) (*exe.Options, error) {
	opts := &exe.Options{
		Length:    options.length,
		Start:     options.start,
		Prefix:    options.prefix,
		DryRun:    options.dryRun,
		Overwrite: options.overwrite,
		Verbose:   options.verbose,
	}
	return opts, nil
}

func perform(paths []string, options *exe.Options) error {
	return exe.Numberize(paths, options)
}
