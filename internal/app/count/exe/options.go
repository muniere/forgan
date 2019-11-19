package exe

import (
	"regexp"

	"github.com/muniere/forgan/internal/pkg/pathname"
)

type Options struct {
	IncludesHidden bool
	Pattern        string
}

func (o *Options) testByVisibility(path *pathname.Pathname) bool {
	if o.IncludesHidden {
		return true
	}

	return !path.IsHiddenR()
}

func (o *Options) testByPattern(path *pathname.Pathname) bool {
	if len(o.Pattern) == 0 {
		return true
	}

	r := regexp.MustCompilePOSIX(o.Pattern)
	return r.MatchString(path.Base())
}
