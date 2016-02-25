package fileinfo

import (
	"os"
	"syscall"
	"time"
)

const (
	flagsPlan9 = 0
)

func getTimespec(fi os.FileInfo) (t Times) {
	stat := fi.Sys().(*syscall.Dir)
	t.flags = flagsPlan9
	t.atime.v = time.Unix(int64(stat.Atime), 0)
	t.mtime.v = time.Unix(int64(stat.Mtime), 0)
	return t
}
