package workdir

import (
	"github.com/jurgen-kluft/Case/fileinfo"
	"os"
	"path/filepath"
)

type WorkItem struct {
	filepath string
	size     int64
	ctime    uint64 // Creation-Time (milli-seconds since January 1 1970)
	mtime    uint64 // Modification-Time (milli-seconds since January 1 1970)
}

type WorkTree interface {
	Scan() []WorkItem
}

type LocalWorkTree struct {
	includeFileFilter   Filter
	excludeFileFilter   Filter
	excludeFolderFilter Filter
}

func (wd LocalWorkTree) Scan(dir string) (files []WorkItem, err error) {
	err = filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
		size := fi.Size()
        fileinfo.
		wi := WorkItem{filepath: path, size: size}
		return nil
	})
}
