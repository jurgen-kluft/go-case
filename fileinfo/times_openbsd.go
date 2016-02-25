package fileinfo

import (
	"os"
	"syscall"
	"time"
)

const (
	flagsOpenBSD = FlagHasCTime
)

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

func getTimespec(fi os.FileInfo) (t Times) {
	stat := fi.Sys().(*syscall.Stat_t)
	t.flags = flagsOpenBSD
	t.atime = timespecToTime(stat.Atim)
	t.mtime = timespecToTime(stat.Mtim)
	t.ctime = timespecToTime(stat.Ctim)
	return t
}
