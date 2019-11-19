package exe

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/muniere/forgan/internal/app/count/stat"
	"github.com/muniere/forgan/internal/pkg/math"
)

func Survey(root string, options *Options) (*stat.Report, error) {
	if err := validate(root); err != nil {
		return nil, err
	}

	fs, err := collect(root, options)
	if err != nil {
		return nil, err
	}

	return &stat.Report{Folders: fs}, nil
}

func Print(r *stat.Report) error {
	plen := 0
	slen := 0

	for _, f := range r.Folders {
		plen = math.Max(plen, len(f.Path))
		slen = math.Max(slen, math.Digits(len(f.Leafs)))
	}

	lines := []string{}

	for _, f := range r.Folders {
		p := f.Path + strings.Repeat(" ", plen-len(f.Path))
		t := fmt.Sprintf("%%s %%%sd", strconv.Itoa(slen))
		lines = append(lines, fmt.Sprintf(t, p, len(f.Leafs)))
	}

	fmt.Println(strings.Join(lines, "\n"))

	return nil
}

func validate(root string) error {
	info, err := os.Stat(root)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return errors.New(fmt.Sprintf("%v is not a directory.", root))
	}
	return nil
}

func collect(root string, options *Options) ([]*stat.Folder, error) {
	t := newTable()

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Errorf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			log.Tracef("walking directory: %s", filepath.Clean(path))
		} else {
			log.Tracef("walking file: %s", filepath.Clean(path))
		}

		return t.merge(path, info, options)
	})

	if err != nil {
		return []*stat.Folder{}, err
	}

	return t.slice(), nil
}
