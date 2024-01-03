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

// 算法 3：dfs (超时）
func maxAlternatingSum3(nums []int) int64 {
	n := len(nums)
	var dfs func(int, int) int
	dfs = func(i, odd int) int {
		if i < 0 {
			return 0
		}
		if odd == 1 {
			return max(dfs(i-1, 1), dfs(i-1, 0)+nums[i])
		}
		return max(dfs(i-1, 0), dfs(i-1, 1)-nums[i])
	}
	return int64(dfs(n-1, 1))
}

// 算法 4：dfs + 记忆化
func maxAlternatingSum4(nums []int) int64 {
	n := len(nums)
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1}
	}
	var dfs func(int, int) int
	dfs = func(i, odd int) (res int) {
		if i < 0 {
			return 0
		}
		m := &memo[i][odd]
		if *m != -1 {
			return *m
		}
		defer func() { *m = res }()
		if odd == 1 {
			return max(dfs(i-1, 1), dfs(i-1, 0)+nums[i])
		}
		return max(dfs(i-1, 0), dfs(i-1, 1)-nums[i])
	}
	return int64(dfs(n-1, 1))
}
