package main

import "math"

// 思路：
// 1. 如果正反一样，那么这个数字一定不会是答案，存入 Same 集合
// 2. 剩下不包含在 Same 集合中的数字，取最小值即可
// 为什么 2 是正确的？因为正反相等的被排除了，剩下的卡片，有几种可能：
//  1. 所有数字互不相等 -> 那显然是取最小值
//  2. 某张卡片的正反数字中的其中一个与另一张相等 -> 如果答案是相等的那个数字（最小），可以将两张卡片都翻转
//  2. 某张卡片的正反数字与另一张卡片相等 -> 如果答案是相等数字中较小的 -> 将较小数字翻转到下面即可
//
// 因此不管什么情况，答案为剩下数字中的最小值
func flipgame(fronts []int, backs []int) int {
	same := make(map[int]struct{})
	for i := 0; i < len(fronts); i++ {
		if fronts[i] == backs[i] {
			same[fronts[i]] = struct{}{}
		}
	}
	res := math.MaxInt32
	for i := 0; i < len(fronts); i++ {
		if _, ok := same[fronts[i]]; !ok {
			res = min(res, fronts[i])
		}
		if _, ok := same[backs[i]]; !ok {
			res = min(res, backs[i])
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
