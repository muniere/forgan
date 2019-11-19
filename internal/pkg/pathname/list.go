package pathname

import (
	"sort"
)

type List struct {
	Values []*Pathname
}

func NewList(values ...*Pathname) *List {
	return &List{
		Values: values,
	}
}

func (l *List) Get(i int) *Pathname {
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

func (l *List) Find(path *Pathname) *Pathname {
	for _, p := range l.Values {
		if p.Value == path.Value {
			return p
		}
	}
	return nil
}

func (l *List) Append(path *Pathname) {
	l.Values = append(l.Values, path)
}

func (l *List) SortStable() {
	sort.SliceStable(l.Values, func(i, j int) bool {
		return l.Values[i].Base() < l.Values[j].Base()
	})
}
