package main

import "math"

// 找规律 | dp | 贪心
func maxAlternatingSum(nums []int) int64 {
	res := int64(nums[0])
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 {
			res += int64(diff)
		}
	}
	return res
}

// 算法 2：动态规划
// 定义f[i][0]表示前i个数中长为偶数的子序列的最大交替和，f[i][1]表示前个数中长为奇数的子序列的最大交替和
// 初始时有 f[0][0] = 0, f[0][1] = -inf
//对于第 i 个数，有选或不选两种决策。
//
//对于f[i+1][0]，若不选第 i 个数，则从 f[i][0] 转移过来，否则从f[i][1]-nums[i]转移过来，取二者最大值。
//
//对于f[i+1][1]，若不选第 i 个数，则从 f[i][1] 转移过来，否则从f[i][0]+nums[i]转移过来，取二者最大值。

func maxAlternatingSum2(nums []int) int64 {
	f0, f1 := 0, math.MinInt64/2
	for _, v := range nums {
		//newF0 := max(f0, f1+p)
		//f1 = max(f1, f0-p)
		//f0 = newF0
		f0, f1 = max(f0, f1-v), max(f1, f0+v)
	}
	return int64(f1)
}
