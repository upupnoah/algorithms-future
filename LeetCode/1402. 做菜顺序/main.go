package _402__做菜顺序

import (
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(satisfaction)))
	s := 0
	f := 0 // f(0) = 0
	for _, x := range satisfaction {
		s += x
		if s <= 0 { // 后面不可能找到更大的 f(k)
			break
		}
		f += s // f(k) = f(k-1) + s(k)
	}
	return f
}
