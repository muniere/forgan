package pathname

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const tilde = "~"

type Pathname struct {
	Value string
}

func New(value string) *Pathname {
	return &Pathname{value}
}

func (p *Pathname) Equals(o *Pathname) bool {
	return p.Value == o.Value
}

func (p *Pathname) IsHidden() bool {
	base := p.Base()

	if base == "." {
		return false
	}

	return strings.HasPrefix(base, ".")
}

func (p *Pathname) IsHiddenR() bool {
	sep := string(os.PathSeparator)

	for x := p; x.Value != "." && x.Value != sep; x = x.Parent() {
		if x.IsHidden() {
			return true
		}
	}

	return false
}

func (p *Pathname) Stat() (os.FileInfo, error) {
	return os.Stat(p.Value)
}

func (p *Pathname) StatIsDir() bool {
	stat, err := p.Stat()

	if err != nil {
		return false
	}

	return stat.IsDir()
}

func (p *Pathname) StatIsFile() bool {
	return !p.StatIsDir()
}

func (p *Pathname) Prefix() string {
	pure := p.Pure()

	if regexp.MustCompile("^\\d+$").MatchString(pure) {
		return ""
	}

	return regexp.MustCompile("\\d+$").ReplaceAllString(pure, "")
}

func (p *Pathname) Abs() *Pathname {
	s, _ := filepath.Abs(p.Expand().Value)
	return New(s)
}

func (p *Pathname) Expand() *Pathname {
	if !strings.HasPrefix(p.Value, tilde+string(os.PathSeparator)) {
		return p
	}

	u, err := os.UserHomeDir()
	if err != nil {
		return p
	}

	s := strings.Replace(p.Value, tilde, u, 1)
	return New(s)
}

func (p *Pathname) Parent() *Pathname {
	s := filepath.Dir(p.Value)
	return &Pathname{s}
}

func (p *Pathname) Clean() *Pathname {
	return &Pathname{filepath.Clean(p.Value)}
}

func (p *Pathname) Join(path ...string) *Pathname {
	base := []string{p.Value}
	plus := append(base, path...)
	return New(filepath.Join(plus...))
}

func (p *Pathname) Base() string {
	return filepath.Base(p.Value)
}

func (p *Pathname) Ext() string {
	return filepath.Ext(p.Value)
}

func (p *Pathname) Pure() string {
	return strings.Replace(p.Base(), p.Ext(), "", -1)
}
