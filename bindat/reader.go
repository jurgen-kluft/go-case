package bindat

import (
	"fmt"
	"github.com/jurgen-kluft/Case/chunk"
	"os"
)



type Reader interface {
	Read(int64, chunk.Chunk) error
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
// these buffers are re-used. The constraint here is that the maximum size
// of Chunk should be enough to load any chunk.
func (reader *FileReader) Read(offset int64, c chunk.Chunk) error {
	n, err := reader.fhnd.ReadAt(c[0:SizeOfHeader], offset)
	if n != SizeOfHeader {
		return fmt.Errorf("FileReader tried to read %d bytes but only got %d bytes", SizeOfHeader, n)
	}
	if err != nil {
		return err
	}
	clen := c.GetLength()
	n, err = reader.fhnd.ReadAt(c[SizeOfHeader:SizeOfHeader+clen], offset+SizeOfHeader)
	if uint32(n) != clen {
		return fmt.Errorf("FileReader tried to read %d bytes but only got %d bytes", clen, n)
	}
	if err == nil {
		return err
	}
	return nil
}

func (reader *FileReader) Close() (err error) {
	err = reader.fhnd.Close()
	reader.fhnd = nil
	return err
}
