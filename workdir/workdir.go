package workdir

import (
	"github.com/jurgen-kluft/Case/fileinfo"
	"time"
)

type WorkFile struct {
	filepath string
	size     int64
	ctime    time.Time // Creation-Time
	mtime    time.Time // Modification-Time
}

type WorkTree interface {
	Scan() []WorkFile
}

type LocalWorkTree struct {
	includeFileFilter   Filter
	excludeFileFilter   Filter
	excludeFolderFilter Filter
}

func (wd LocalWorkTree) Scan() (files []WorkFile, err error) {
}
