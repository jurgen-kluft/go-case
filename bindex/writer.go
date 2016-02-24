package bindex

import (
	"fmt"
	"os"
)

type Writer interface {
	Write(Dex) error
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
func (writer *FileAppendWriter) Write(c Dex) error {
	n, err := writer.fhnd.Write(c)
	if n != len(c) {
		return fmt.Errorf("FileAppendWriter tried to write %d bytes but only did %d bytes", len(c), n)
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
