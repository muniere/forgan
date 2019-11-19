package rule

import (
	"github.com/muniere/forgan/internal/pkg/pathname"
)

type Rule struct {
	Src *pathname.Pathname
	Dst *pathname.Pathname
}

func New(src string, dst string) *Rule {
	return &Rule{
		Src: pathname.New(src),
		Dst: pathname.New(dst),
	}
}

func (r *Rule) Equals(o *Rule) bool {
	return r.Src.Equals(o.Src) && r.Dst.Equals(o.Dst)
}
