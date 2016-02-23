package hashing

import (
	"errors"
	"github.com/jurgen-kluft/Case/hashing/skein"
)

// CompareHashes will compare 2 byte arrays and return true if they
// contain identical bytes. The user needs to make sure to pass 2
// byte arrays of the same length.
func CompareHashes(left []byte, right []byte) bool {
	for i, b := range left {
		if b != right[i] {
			return false
		}
	}
	return true
}

// Hasher is the hash engine that requires an instance to be used to compute
// a hash of a block of data. Call NewHasher() to obtain an instance of Hasher.
type Hasher interface {
	Hash([]byte, []byte)
}

type SkeinHasher struct {
	engine *skein.Skein
}

type HasherType int

const (
	HasherTypeSkein HasherType = 1
)

func NewHasher(hasherType HasherType) (Hasher, error) {
	if hasherType == HasherTypeSkein {
		hasher := &SkeinHasher{}
		hasher.engine, _ = skein.New(skein.Skein512, 256)
		return hasher, nil
	}
	return nil, errors.New("unsupported hasher type")
}

func (h *SkeinHasher) Hash(data []byte, hash []byte) {
	h.engine.Update(data)
	h.engine.DoFinal(hash)
	return
}
