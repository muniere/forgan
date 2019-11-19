package exe

import (
	"os"
	"path/filepath"
	"sort"

	log "github.com/sirupsen/logrus"

	"github.com/muniere/forgan/internal/app/count/stat"
	"github.com/muniere/forgan/internal/pkg/pathname"
)

type table struct {
	values map[string]*stat.Folder
}

func newTable() *table {
	return &table{
		values: map[string]*stat.Folder{},
	}
}

func (t *table) merge(path string, info os.FileInfo, options *Options) error {
	if info.IsDir() {
		return t.mergeDir(path, info, options)
	} else {
		return t.mergeFile(path, info, options)
	}
}

func (t *table) mergeDir(path string, info os.FileInfo, options *Options) error {
	p := pathname.New(path).Clean()

	if !options.testByVisibility(p) {
		log.Debugf("skip hidden directory: %v", p.Value)
		return filepath.SkipDir
	}

	// do not match patterns for directories

	t.values[p.Value] = &stat.Folder{
		Path:  p.Value,
		Leafs: []os.FileInfo{},
	}
	return nil
}

func (t *table) mergeFile(path string, info os.FileInfo, options *Options) error {
	p := pathname.New(path).Clean()

	if !options.testByVisibility(p) {
		log.Debugf("skip hidden file: %s", p.Value)
		return nil
	}
	if !options.testByPattern(p) {
		log.Debugf("skip file: %s", p.Value)
		return nil
	}

	k := p.Parent().Value
	f := t.values[k]

	if f == nil {
		t.values[k] = &stat.Folder{
			Path:  k,
			Leafs: []os.FileInfo{info},
		}
		return nil
	}

	f.Append(info)
	return nil
}

func (t *table) slice() []*stat.Folder {
	fs := []*stat.Folder{}

	for _, f := range t.values {
		fs = append(fs, f)
	}

	sort.Slice(fs, func(i, j int) bool {
		return fs[i].Path < fs[j].Path
	})

	return fs
}
