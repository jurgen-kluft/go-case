package workdir

import (
	"fmt"
	"github.com/jurgen-kluft/Case/fileinfo"
	"os"
	"path/filepath"
)

type WorkItem struct {
	filepath string
	size     int64
	btime    uint64 // Creation-Time (milli-seconds since January 1 1970)
	mtime    uint64 // Modification-Time (milli-seconds since January 1 1970)
}

type WorkTree interface {
	Scan(dirpath string) (items []WorkItem, err error)
}

type localWorkTree struct {
	includeFileFilter   Filter
	excludeFileFilter   Filter
	excludeFolderFilter Filter
}

func NewLocalWorkTree(includeFilter Filter, excludeFilter Filter, excludeFolderFilter Filter) WorkTree {
	return &localWorkTree{includeFileFilter: includeFilter, excludeFileFilter: excludeFilter, excludeFolderFilter: excludeFolderFilter}
}

func (worktree localWorkTree) Scan(dirpath string) (items []WorkItem, err error) {
	err = filepath.Walk(dirpath, func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() == false {
			match, merr := worktree.excludeFileFilter.Match(path)
			if merr == nil && !match {
				fmt.Println(path)
				match, merr = worktree.includeFileFilter.Match(path)
				if merr == nil && match {
					size := fi.Size()
					btime, mtime := fileinfo.GetCreationAndModificationTime64(fi)
					wi := WorkItem{filepath: path, size: size, btime: btime, mtime: mtime}
					items = append(items, wi)
				}
			}
		} else {
			match, merr := worktree.excludeFolderFilter.Match(path)
			if merr == nil && match {
				return filepath.SkipDir
			}
		}
		return nil
	})
	return
}
