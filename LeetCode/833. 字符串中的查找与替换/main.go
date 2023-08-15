package main

import (
	"sort"
	"strings"
)

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	m := len(indices)
	id := make([]int, m)
	for i := 0; i < m; i++ {
		id[i] = i
	}
	// 每次替换只会影响后面字符的下标，因此可以从后往前做
	// 映射关系， id[i] -> indices[i]
	// 通过 id[i] 找到对应的 indices[i]
	// id = [0, 1, 2]
	// indices = [100, 2, 200] -> id = [1, 0, 2]
	sort.Slice(id, func(i, j int) bool {
		return indices[id[i]] < indices[id[j]]
	})
	for i := m - 1; i >= 0; i-- {
		k := id[i]
		if strings.HasPrefix(s[indices[k]:], sources[k]) {
			s = s[:indices[k]] + targets[k] + s[indices[k]+len(sources[k]):]
		}
	}
	return s
}
