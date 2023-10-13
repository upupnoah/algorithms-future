package _488__Avoid_Flood_in_The_City

import "sort"

func avoidFlood(rains []int) []int {
	n := len(rains)
	var zero []int
	ans := make([]int, n)
	lakeIdx := make(map[int]int)

	// 初始化 0 的情况，并且没有满的 lake
	for i := 0; i < n; i++ {
		ans[i] = 1
	}

	for i, lake := range rains {
		if lake == 0 {
			zero = append(zero, i)
		} else {
			ans[i] = -1
			if id, ok := lakeIdx[lake]; ok {
				idx := sort.SearchInts(zero, id)
				if idx == len(zero) {
					return []int{}
				}
				ans[zero[idx]] = lake
				copy(zero[idx:len(zero)-1], zero[idx+1:])
				zero = zero[:len(zero)-1]
			}
			lakeIdx[lake] = i
		}
	}

	return ans
}
