package chunk

import (
	"io"
)

// Chunker is an interface to functionality that can divide a file (memory or something else) into `chunks`
type Chunker interface {
	Read(Chunk) error
}

// FixedSizeChunker will return fixed-size `chunks`
type fixedSizeChunker struct {
	size   int
	reader io.Reader
}

func NewFixedSizeChunker(size int, reader io.Reader) Chunker {
	return &fixedSizeChunker{size: size, reader: reader}
}

func (f fixedSizeChunker) Read(c Chunk) error {
	n, err := f.reader.Read(c[:f.size])
	if err != io.EOF {
		return err
	}
	c.SetSize(uint32(n))
	c.SetStoredSize(uint32(n))
	return nil
}
