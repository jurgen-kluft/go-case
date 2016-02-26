package fileinfo

import (
	"os"
	"syscall"
	"time"
)

// HasChangeTime and HasBirthTime are true if and only if
// the target OS supports them.
const (
	flagsSolaris = flagHasCTime
)

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

func getTimespec(fi os.FileInfo) (t Times) {
	stat := fi.Sys().(*syscall.Stat_t)
	t.flags = flagsSolaris
	t.atime = timespecToTime(stat.Atim)
	t.mtime = timespecToTime(stat.Mtim)
	t.ctime = timespecToTime(stat.Ctim)
	return t
}
