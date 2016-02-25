package chunk

import (
	"encoding/binary"
)

const (
	cDataAlignment = 64
	cSizeOfHeader  = 64
	cSizeOfHash    = 32
	cOffsetOfHash  = 0
	cOffsetOfASize = 32
	cSizeOfASize   = 4
	cOffsetOfSSize = 36
	cSizeOfSSize   = 4
	cOffsetOfMagic = 40
	cSizeOfMagic   = 24
)

// Skein256 hash of the word 'Case'
var cMAGIC = []byte{0x0f, 0x99, 0x3a, 0x56, 0x25, 0xc7, 0xb1, 0xb1, 0x98, 0x01, 0x78, 0x89, 0xd3, 0x01, 0x57, 0xa5, 0x70, 0x94, 0x3e, 0xbf, 0x2d, 0x5a, 0xc6, 0xb0, 0xe2, 0xf4, 0x62, 0xca, 0x2c, 0x54, 0x42, 0xe0}

// Chunk:
// [32] Hash
// [ 4] ASize (actual size)
// [ 4] SSize (stored size)
// [24] Magic
type Chunk []byte

func NewChunk(size int) Chunk {
	c := Chunk(make([]byte, size+cSizeOfHeader))
	for i := 0; i < cSizeOfMagic; i++ {
		c[cOffsetOfMagic+i] = cMAGIC[i]
	}
	c.SetSize(0)
	c.SetStoredSize(0)
	return c
}

func (c Chunk) GetHash() []byte {
	return c[cOffsetOfHash : cOffsetOfHash+cSizeOfHash]
}
func (c Chunk) GetSize() uint32 {
	return binary.BigEndian.Uint32(c[cOffsetOfASize : cOffsetOfASize+cSizeOfASize])
}
func (c Chunk) SetSize(size uint32) {
	binary.BigEndian.PutUint32(c[cOffsetOfASize:cOffsetOfASize+cSizeOfASize], size)
}
func (c Chunk) GetStoredSize() uint32 {
	return binary.BigEndian.Uint32(c[cOffsetOfSSize : cOffsetOfSSize+cSizeOfSSize])
}
func (c Chunk) SetStoredSize(size uint32) {
	binary.BigEndian.PutUint32(c[cOffsetOfSSize:cOffsetOfSSize+cSizeOfSSize], size)
}
func (c Chunk) GetStoredBlock() []byte {
	ssize := c.GetStoredSize()
	return c[cSizeOfHeader : cSizeOfHeader+ssize]
}

func (c Chunk) GetMaxPossibleDataSize() uint32 {
	return uint32(len(c)) - cSizeOfHeader
}
func (c Chunk) GetMaxPossibleDataBlock() []byte {
	return c[cSizeOfHeader:]
}

func (c Chunk) IsValid() bool {
	valid := (c.GetSize() >= c.GetStoredSize())
	valid = valid && (c.GetSize() <= c.GetMaxPossibleDataSize())
	return valid
}
func (c Chunk) HasMagic() bool {
	for i, b := range c[cOffsetOfMagic : cOffsetOfMagic+cSizeOfMagic] {
		if cMAGIC[i] != b {
			return false
		}
	}
	return true
}
