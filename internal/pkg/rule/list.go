package rule

import (
	"github.com/muniere/forgan/internal/pkg/pathname"
)

type List struct {
	Values []*Rule
}

func NewList(values ...*Rule) *List {
	return &List{
		Values: values,
	}
}

func (l *List) Get(i int) *Rule {
	return l.Values[i]
}

func (l *List) Equals(o *List) bool {
	if len(l.Values) != len(o.Values) {
		return false
	}

	for i := range l.Values {
		x := l.Get(i)
		y := o.Get(i)

		if !x.Equals(y) {
			return false
		}
	}

	return true
}

func (l *List) FindBySrc(path *pathname.Pathname) *Rule {
	for _, r := range l.Values {
		if r.Src.Equals(path) {
			return r
		}
	}
	return nil
}

func (l *List) FindByDst(path *pathname.Pathname) *Rule {
	for _, r := range l.Values {
		if r.Dst.Equals(path) {
			return r
		}
	}
	return nil
}

func (l *List) Append(rule *Rule) {
	l.Values = append(l.Values, rule)
}
