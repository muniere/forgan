package stat

import (
	"os"
)

type Folder struct {
	Path  string
	Leafs []os.FileInfo
}

func (f *Folder) Append(info os.FileInfo) {
	f.Leafs = append(f.Leafs, info)
}
