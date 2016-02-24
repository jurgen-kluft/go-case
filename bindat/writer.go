package bindat

import (
	"fmt"
	"github.com/jurgen-kluft/Case/chunk"
	"os"
)

type Writer interface {
	Write(int64, chunk.Chunk) error
}

func NewFileAppendWriter(filepath string) (writer Writer, err error) {
	fhnd, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}
	return &FileAppendWriter{fhnd: fhnd}, nil
}

type FileAppendWriter struct {
	fhnd *os.File
}

func align(v uint32, a uint32) uint32 {
	v = (v + (a - 1)) & ^(a - 1)
	return v
}

// Write,
func (writer *FileAppendWriter) Write(offset int64, c chunk.Chunk) error {
	n, err := writer.fhnd.WriteAt(c[0:SizeOfHeader], offset)
	if n != SizeOfHeader {
		return fmt.Errorf("FileAppendWriter tried to write %d bytes but only did %d bytes", SizeOfHeader, n)
	}
	if err != nil {
		return err
	}
	clen := c.GetLength()
	clen = align(clen, SizeOfHeader)

	n, err = writer.fhnd.WriteAt(c[SizeOfHeader:SizeOfHeader+clen], offset+SizeOfHeader)
	if uint32(n) != clen {
		return fmt.Errorf("FileAppendWriter tried to write %d bytes but only did %d bytes", clen, n)
	}
	if err != nil {
		return err
	}
	return nil
}

// Close, closes the file-handle owned by the file-append-writer
func (writer *FileAppendWriter) Close() (err error) {
	err = writer.fhnd.Close()
	writer.fhnd = nil
	return err
}
