package _465__切割后面积最大的蛋糕

import "sort"

func getMaxSize(size int, cuts []int) int {
	sort.Ints(cuts)
	res := max(cuts[0], size-cuts[len(cuts)-1]) // 两侧
	for i := 1; i < len(cuts); i++ {
		res = max(res, cuts[i]-cuts[i-1])
	}
	return res
}

func maxArea(h, w int, horizontalCuts, verticalCus []int) int {
	return getMaxSize(h, horizontalCuts) * getMaxSize(w, verticalCus) % 1_000_000_007
}
