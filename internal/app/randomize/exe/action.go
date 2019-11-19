package exe

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/muniere/forgan/internal/pkg/pathname"
	"github.com/muniere/forgan/internal/pkg/rule"
)

const retry = 5

func Randomize(paths []string, options *Options) error {
	rand.Seed(time.Now().UnixNano())

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

func Print(rules *rule.List) error {
	for _, r := range rules.Values {
		fmt.Printf("%s/{%s => %s}\n", r.Src.Parent().Value, r.Src.Base(), r.Dst.Base())
	}
	return nil
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

	for _, src := range sources.Values {
		for i := 0; ; i++ {
			if i >= retry {
				return nil, errors.New("Cannot compose rename rules")
			}

			pre := prefix(src, options)
			seq := sequence(src, options)
			dst := src.Parent().Join(pre + seq + src.Ext())

			if found := sources.Find(dst); found != nil {
				log.Debugf("New pathname conflicted: %v", found.Value)
				continue
			}
			if found := l.FindByDst(dst); found != nil {
				log.Debugf("New pathname conflicted: %v", found.Dst.Value)
				continue
			}

			r := &rule.Rule{
				Src: src,
				Dst: dst,
			}
			l.Append(r)
			break
		}
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

func sequence(path *pathname.Pathname, options *Options) string {
	max := int(math.Pow10(options.Length)) - options.Start
	number := rand.Intn(max) + options.Start
	format := fmt.Sprintf("%%0%dd", options.Length)
	return fmt.Sprintf(format, number)
}
