package count

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"github.com/muniere/forgan/internal/app/count/cli"
	"github.com/muniere/forgan/internal/app/count/exe"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "count",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd, args)
		},
	}

	cli.Prepare(cmd)

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	ctx, err := cli.Parse(args, cmd.Flags())
	if err != nil {
		return err
	}

	if err := prepare(ctx.Options); err != nil {
		return err
	}

	opts, err := translate(ctx.Options)
	if err != nil {
		return err
	}

	for _, p := range ctx.Paths {
		if err := perform(p, opts); err != nil {
			log.Warnf("%v", err)
		}
	}

	return nil
}

func prepare(options *cli.Options) error {
	if options.Verbose {
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

func translate(options *cli.Options) (*exe.Options, error) {
	opts := &exe.Options{
		IncludesHidden: options.IncludesHidden,
		Pattern:        options.Pattern,
	}
	return opts, nil
}

func perform(path string, options *exe.Options) error {
	report, err := exe.Survey(path, options)
	if err != nil {
		return err
	}

	return exe.Print(report)
}
