package cleanpath

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type longestFirst []string

func RemoveGoPath(path string) string {
	dirs := filepath.SplitList(os.Getenv("GOPATH"))
	// Sort in decreasing order by length so the longest matching prefix is removed
	sort.Stable(longestFirst(dirs))
	for _, dir := range dirs {
		srcdir := filepath.Join(dir, "src")
		rel, err := filepath.Rel(srcdir, path)
		// filepath.Rel can traverse parent directories, don't want those
		if err == nil && !strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			return rel
		}
	}

	return path
}

func (strs longestFirst) Len() int {
	return len(strs)
}

func (strs longestFirst) Less(i, j int) bool {
	return len(strs[i]) > len(strs[j])
}

func (strs longestFirst) Swap(i, j int) {
	strs[i], strs[j] = strs[j], strs[i]
}
