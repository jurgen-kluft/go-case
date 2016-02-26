package fileinfo

import (
	"os"
	"syscall"
	"time"
)

const (
	flagsWindows = flagHasBTime
)

func getTimespec(fi os.FileInfo) (t Times) {
	stat := fi.Sys().(*syscall.Win32FileAttributeData)
	t.flags = flagsWindows
	t.atime = time.Unix(0, stat.LastAccessTime.Nanoseconds())
	t.mtime = time.Unix(0, stat.LastWriteTime.Nanoseconds())
	t.btime = time.Unix(0, stat.CreationTime.Nanoseconds())
	return t
}
