package workdir

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	includeFilter := NewIncludeFilter([]string{"**\\*.jpg"})
	excludeFilter := NewIgnoreFilter([]string{"**\\.gitignore", "**\\*.txt", "**\\*.md", "**\\*.cs", "**\\*.cpp", "**\\*.h"})
	excludeFolderFilter := NewIgnoreFilter([]string{"**\\.git"})

	worktree := NewLocalWorkTree(includeFilter, excludeFilter, excludeFolderFilter)
	items, err := worktree.Scan("d:\\Jurgen\\private\\Vault")
	if err != nil {
		t.Fail()
	}

	for _, item := range items {
		fmt.Println(item.filepath)
	}
}
