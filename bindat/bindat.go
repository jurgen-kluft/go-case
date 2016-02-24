package bindat

import (
	"fmt"
	"path"
)

const (
	DataAlignment = 64
	SizeOfHeader  = 64
)

func GetBinDatPath(repopath string, index int) string {
	return path.Join(repopath, "dat", fmt.Sprintf("%04X.dat", index))
}

func NewReader(datpath string) (reader Reader, err error) {
	return NewFileReader(datpath)
}

func NewWriter(datpath string) (writer Writer, err error) {
	return NewFileAppendWriter(datpath)
}
