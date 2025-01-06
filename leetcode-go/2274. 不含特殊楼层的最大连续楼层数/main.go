package main

import "slices"

// 1. sort
// 2. iterate

func maxConsecutive(bottom int, top int, special []int) int {
	slices.Sort(special)
	n := len(special)
	ans := max(special[0]-bottom, top-special[n-1])
	for i := 1; i < n; i++ {
		ans = max(ans, special[i]-special[i-1]-1)
	}
	return ans
}
