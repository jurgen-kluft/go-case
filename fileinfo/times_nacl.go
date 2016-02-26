package fileinfo

import (
	"os"
	"syscall"
	"time"
)

const (
	flagsNaCL = flagHasCTime
)

func timespecToTime(sec, nsec int64) time.Time {
	return time.Unix(sec, nsec)
}

func getTimespec(fi os.FileInfo) (t Times) {
	stat := fi.Sys().(*syscall.Stat_t)
	t.flags = flagsNaCL
	t.atime = timespecToTime(stat.Atime, stat.AtimeNsec)
	t.mtime = timespecToTime(stat.Mtime, stat.MtimeNsec)
	t.ctime = timespecToTime(stat.Ctime, stat.CtimeNsec)
	return t
}
