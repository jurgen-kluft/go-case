// Package fileinfo provides a platform-independent way to get atime, mtime, ctime and btime for files.
package fileinfo

import (
	"os"
	"time"
)

const (
	flagHasCTime = 1
	flagHasBTime = 2
)

func HasFlag(flags int, flag int) bool {
	return (flags & flag) == flag
}

type Times struct {
	flags int
	atime time.Time
	mtime time.Time
	ctime time.Time
	btime time.Time
}

func (t Times) AccessTime() time.Time {
	return t.atime
}
func (t Times) ModTime() time.Time {
	return t.mtime
}
func (t Times) ChangeTime() time.Time {
	return t.ctime
}
func (t Times) HasChangeTime() bool {
	return (t.flags & flagHasCTime) == flagHasCTime
}
func (t Times) BirthTime() time.Time {
	return t.btime
}
func (t Times) HasBirthTime() bool {
	return (t.flags & flagHasBTime) == flagHasBTime
}

func Stat(filepath string) (t Times, err error) {
	fhdn, err := os.Open(filepath)
	if err != nil {
		return
	}
	fi, err := fhdn.Stat()
	if err != nil {
		return
	}
	t = Get(fi)
	return t, err
}

func Get(fi os.FileInfo) (times Times) {
	times = getTimespec(fi)
	return
}

var epochTime = time.Unix(0, 0)

func TimeToTime64(t time.Time) uint64 {
	d := t.Sub(epochTime)
	return uint64(d.Nanoseconds())
}

func GetCreationAndModificationTime64(fi os.FileInfo) (btime uint64, mtime uint64) {
	times := getTimespec(fi)
	return TimeToTime64(times.btime), TimeToTime64(times.ctime)
}
