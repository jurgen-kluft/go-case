package workdir

import (
	"github.com/jurgen-kluft/Case/glob"
)

type Filter interface {
	Match(string) (bool, error)
}

type includeFilter struct {
	patterns []string
}

func NewIncludeFilter(patterns []string) Filter {
	return &includeFilter{patterns: patterns}
}

func filterMatch(patterns []string, s string) (bool, error) {
	for _, p := range patterns {
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

func (filter includeFilter) Match(s string) (bool, error) {
	return filterMatch(filter.patterns, s)
}

type ignoreFilter struct {
	patterns []string
}

func NewIgnoreFilter(patterns []string) Filter {
	return &ignoreFilter{patterns: patterns}
}

func (filter ignoreFilter) Match(s string) (bool, error) {
	return filterMatch(filter.patterns, s)
}
