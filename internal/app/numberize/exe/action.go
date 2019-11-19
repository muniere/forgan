package exe

import (
	"errors"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/muniere/forgan/internal/pkg/pathname"
	"github.com/muniere/forgan/internal/pkg/rule"
)

func Numberize(paths []string, options *Options) error {
	sources, err := prepare(paths)
	if err != nil {
		return err
	}

	rules, err := compose(sources, options)
	if err != nil {
		return err
	}

	if options.DryRun {
		return Print(rules)
	} else {
		return Apply(rules)
	}
}

func Apply(rules *rule.List) error {
	for _, r := range rules.Values {
		log.Debugf("%s/{%s => %s}\n", r.Src.Parent().Value, r.Src.Base(), r.Dst.Base())
		if err := os.Rename(r.Src.Value, r.Dst.Value); err != nil {
			return err
		}
	}
	return nil
}

func Print(rules *rule.List) error {
	for _, r := range rules.Values {
		fmt.Printf("%s/{%s => %s}\n", r.Src.Parent().Value, r.Src.Base(), r.Dst.Base())
	}
	return nil
}

func prepare(sources []string) (*pathname.List, error) {
	l := pathname.NewList()

	for _, s := range sources {
		p := &pathname.Pathname{
			Value: s,
		}
		if p.IsHidden() {
			continue
		}
		l.Append(p)
	}

	l.SortStable()

	return l, nil
}

func compose(sources *pathname.List, options *Options) (*rule.List, error) {
	l := rule.NewList()

	for i, src := range sources.Values {
		pre := prefix(src, options)
		seq := sequence(i, options)
		dst := src.Parent().Join(pre + seq + src.Ext())

		if found := sources.Find(dst); found != nil {
			return nil, errors.New(fmt.Sprintf("New pathname conflicted: %v", found.Value))
		}
		if found := l.FindByDst(dst); found != nil {
			return nil, errors.New(fmt.Sprintf("New pathname conflicted: %v", found.Dst.Value))
		}

		r := &rule.Rule{
			Src: src,
			Dst: dst,
		}
		l.Append(r)
	}

	return l, nil
}

func prefix(path *pathname.Pathname, options *Options) string {
	if len(options.Prefix) > 0 {
		return options.Prefix
	}

	prefix := path.Prefix()

	if len(prefix) == 0 {
		return prefix
	}

	separator := "_"

	if strings.HasSuffix(prefix, separator) {
		return prefix
	} else {
		return prefix + separator
	}
}

func sequence(n int, options *Options) string {
	number := n + options.Start
	format := fmt.Sprintf("%%0%dd", options.Length)
	return fmt.Sprintf(format, number)
}
