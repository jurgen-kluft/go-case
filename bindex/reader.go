package bindex

import (
	"fmt"
	"os"
)

type Reader interface {
	Count() int
	Read(int) (Dex, error)
}

type InMemoryReader struct {
	filepath string
	dex      []byte
}

// NewInMemoryReader is a dex reader that reads the whole content into memory
func NewInMemoryReader(filepath string) (Reader, error) {
	fhnd, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer fhnd.Close()

	finfo, err := fhnd.Stat()
	if err != nil {
		return nil, err
	}

	size := finfo.Size()
	block := make([]byte, size)
	n, err := fhnd.Read(block)
	if err != nil {
		return nil, err
	}
	if int64(n) != size {
		return nil, err
	}

	return &InMemoryReader{filepath: filepath, dex: block}, nil
}

func (r *InMemoryReader) Count() int {
	return len(r.dex) / cSizeOfDex
}

// Load
func (r *InMemoryReader) Read(i int) (dex Dex, err error) {
	offset := i * cSizeOfDex
	if offset <= (len(r.dex) - cSizeOfDex) {
		return Dex(r.dex[offset : offset+cSizeOfDex]), nil
	}
	return Dex([]byte{}), fmt.Errorf("FileReader::Read was asked to get a Dex entry at index %d while the maximum index is %d", i, r.Count()-1)
}
