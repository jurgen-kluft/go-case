package bindex

import (
	"encoding/binary"
	"fmt"
	"path"
)

type BinDex interface {
	Count() int
	Read(int) (Dex, error)
	Write(Dex) error
}

const (
	cDataAlignment  = 64
	cOffsetOfHash   = 0
	cSizeOfHash     = 16
	cOffsetOfOffset = 16
	cSizeOfOffset   = 4
	cSizeOfDex      = cSizeOfHash + cSizeOfOffset
)

type Dex []byte

func (i Dex) GetHash() []byte {
	return i[0:cSizeOfHash]
}
func (i Dex) SetHash(hash []byte) {
	for o, b := range hash {
		i[cOffsetOfHash+o] = b
	}
}
func (i Dex) GetOffset() uint64 {
	offset := uint64(binary.BigEndian.Uint32(i[cOffsetOfOffset : cOffsetOfOffset+cSizeOfOffset]))
	offset *= cDataAlignment
	return offset
}
func (i Dex) SetOffset(offset uint64) {
	offset32 := uint32(offset / cDataAlignment)
	binary.BigEndian.PutUint32(i[cOffsetOfOffset:cOffsetOfOffset+cSizeOfOffset], offset32)
}

type StandardBinDex struct {
	reader Reader
	writer Writer
}

func GetBinDexPath(repopath string, index int) string {
	return path.Join(repopath, "dex", fmt.Sprintf("%04X.dex", index))
}

func NewReader(dexpath string) (Reader, error) {
	return NewInMemoryReader(dexpath)
}
func NewWriter(dexpath string) (Writer, error) {
	return NewFileAppendWriter(dexpath)
}
func NewBinDex(reader Reader, writer Writer) BinDex {
	dex := &StandardBinDex{reader: reader, writer: writer}
	return dex
}

func (bd *StandardBinDex) Count() int {
	return bd.reader.Count()
}

func (bd *StandardBinDex) Read(i int) (dex Dex, err error) {
	return bd.reader.Read(i)
}

func (bd *StandardBinDex) Write(dex Dex) error {
	return bd.writer.Write(dex)
}
