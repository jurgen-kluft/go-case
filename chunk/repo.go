package chunk

import (
	"github.com/jurgen-kluft/Case/hashing"
)

type repo struct {
	bins []Bin

	hasher hashing.Hasher
}
