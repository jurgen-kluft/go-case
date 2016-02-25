package workdir

import (
	"github.com/jurgen-kluft/Case/glob"
)

type Filter interface {
	Match(string) bool
}

type IgnoreFilter struct {
	patterns []string
}

func (filter IgnoreFilter) Match(s string) (bool, error) {
	for _, p := range filter.patterns {
		match, err := glob.PathMatch(p, s)
		if err != nil {
			return false, err
		}
		if match {
			return true, nil
		}
	}
	return false, nil
}
