package binidx

import (
	"os"
)

type Reader interface {
	Read([]byte) (int64, error)
}

type FileReader struct {
	filepath  string
	fhnd      *os.File
	remainder int64
}

func NewFileReader(filepath string) Reader {
	return &FileReader{filepath: filepath, remainder: int64(0)}
}

// Load
func (reader *FileReader) Read(block []byte) (int64, error) {
	var err error
	if reader.fhnd == nil {
		reader.fhnd, err = os.Open(reader.filepath)
		if err != nil {
			return int64(0), err
		}

		finfo, err := reader.fhnd.Stat()
		if err != nil {
			return int64(0), err
		}

		reader.remainder = finfo.Size()
	}
	if len(block) == 0 {
		return reader.remainder, nil
	}

	n, err := reader.fhnd.Read(block)
	if err != nil {
		return 0, err
	}
	reader.remainder -= int64(n)

	if reader.remainder <= 0 {
		err = reader.fhnd.Close()
		reader.fhnd = nil
		reader.remainder = 0
		return 0, err
	}

	return reader.remainder, nil
}
