package fileinfo

import (
	"os"
	"testing"
	"time"
)

func TestStat(t *testing.T) {
	fileTest(t, func(f *os.File) {
		ts, err := Stat(f.Name())
		if err != nil {
			t.Error(err.Error())
		}
		timespecTest(ts, newInterval(time.Now(), 50*time.Millisecond), t)
	})
}

func TestGet(t *testing.T) {
	fileTest(t, func(f *os.File) {
		fi, err := os.Stat(f.Name())
		if err != nil {
			t.Error(err.Error())
		}
		timespecTest(Get(fi), newInterval(time.Now(), 50*time.Millisecond), t)
	})
}

func TestStatErr(t *testing.T) {
	_, err := Stat("badfile?")
	if err == nil {
		t.Error("expected an error")
	}
}
