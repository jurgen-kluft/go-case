package bindat

import (
	"github.com/jurgen-kluft/Case/chunk"
	"os"
)

type Reader interface {
	Read(uint64, Chunk) error
}

type Closer interface {
	Close() error
}

func NewFileReader(filepath string) (reader Reader, err error) {
	fhnd, err := os.OpenFile(filepath, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	reader = &FileReader{fhnd: fhnd}
	return
}

type FileReader struct {
	fhnd *os.File
}

// Read, buffers are already pre-allocated to reduce large-allocations since
// these buffers are re-used.
func (reader *FileReader) Read(offset uint64, c Chunk) error {
	return nil
}

func (reader *FileReader) Close() (err error) {
	err = reader.fhnd.Close()
	reader.fhnd = nil
	return err
}
