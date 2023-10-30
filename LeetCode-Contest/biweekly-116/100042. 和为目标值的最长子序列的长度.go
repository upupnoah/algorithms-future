package biweekly_116

import "math"

// 01背包, 每个数选或不选
// 有 n 个数, 和为 target, 求最长子序列的长度
// 每个数只能使用一次

func lengthOfLongestSubsequence(nums []int, target int) int {
	f := make([]int, target+1)
	for i := 1; i <= target; i++ {
		f[i] = math.MinInt
	}
	s := 0
	for _, x := range nums {
		// 如果前两个数的和为 5,那么是不可能组成大于 5 的,后面的 dp 一定是 -inf, 无需计算
		// s 为前 i 个数的和
		s = min(s+x, target)
		for j := s; j >= x; j-- {
			f[j] = max(f[j], f[j-x]+1)
		}
	}
	if f[target] > 0 {
		return f[target]
	}
	return -1
}
