package binidx

import (
	"encoding/binary"
	"errors"
	"github.com/jurgen-kluft/Case/hashing"
	"sort"
)

type BinIdx interface {
	Count() int
	IndexAt(int) Index
	HashAt(int) []byte
	Load() error
	Save() error
}

func New(reader Reader) BinIdx {
	return &StandardBinIdx{reader: reader, hdr: []byte{}, idx: []byte{}, block: []byte{}, sorted: []int{}}
}

type StandardBinIdx struct {
	reader Reader
	hdr    []byte // Slice of block
	idx    []byte // Slice of block
	block  []byte
	sorted []int // Sorted index
}

const (
	cDataAlignment     = 64
	cSizeOfIndexHash   = 16
	cSizeOfIndexOffset = 4
	cSizeOfIndex       = cSizeOfIndexHash + cSizeOfIndexOffset
)

type Index []byte

func (i Index) GetHash() []byte {
	return i[0:cSizeOfIndexHash]
}

func (i Index) GetOffset() uint64 {
	offset := uint64(binary.BigEndian.Uint32(i[cSizeOfIndexHash:cSizeOfIndex]))
	offset *= cDataAlignment
	return offset
}

func (idx *StandardBinIdx) Count() int {
	return len(idx.idx) / cSizeOfIndex
}

func (idx *StandardBinIdx) IndexAt(i int) (index Index) {
	index = Index(idx.idx[i*cSizeOfIndex:])
	return
}

func (idx *StandardBinIdx) HashAt(i int) []byte {
	index := Index(idx.idx[i*cSizeOfIndex:])
	return index.GetHash()
}

func (idx *StandardBinIdx) Load() error {
	size, err := idx.reader.Read([]byte{})
	if err != nil {
		return err
	}
	if size < 32 {
		return errors.New("Bin-Index header size == 32 but reader informs us that the total size < 32")
	}
	idx.block = make([]byte, size)
	size, err = idx.reader.Read(idx.block)
	if err == nil {
		idx.hdr = idx.block[:32]
		idx.idx = idx.block[32:]
		idx.sorted = make([]int)
	}
	return err
}

func (idx *StandardBinIdx) Save() error {
	return errors.New("Bin-Index save method not implemented yet")
}

// FindHash will return the index in IndexDB.db, -1 if entry doesn't exist
func (idx *StandardBinIdx) Find(hash []byte) (int, error) {
	// Binary-Search
	n := len(idx.idx) / cSizeOfIndex
	sort.Search(n, func(i int) bool {
		return hashing.CompareHashes(hash, idx.HashAt(i))
	})
	return -1, errors.New("BinIdx could not find Index by hash")
}
