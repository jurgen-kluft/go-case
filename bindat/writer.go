package bindat

import (
	"github.com/jurgen-kluft/Case/chunk"
	"os"
)

type Writer interface {
	Write(uint64, Chunk) error
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

// Write,
func (writer *FileAppendWriter) Write(offset uint64, c Chunk) error {
	return nil
}

// Close, closes the file-append writer
func (writer *FileAppendWriter) Close() (err error) {
	err = writer.fhnd.Close()
	writer.fhnd = nil
	return err
}
