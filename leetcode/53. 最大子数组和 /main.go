package main

import "math"

// 思路：动态规划
// dp[i] 表示以 nums[i] 为结尾的最大子数组和
// dp[i] = nums[i] + max(dp[i-1], 0)
// 由于每次只用到 dp[i-1], 因此可以用一个变量记录 dp[i-1]的值
func maxSubArray(nums []int) int {
	res := math.MinInt32
	for i, last := 0, 0; i < len(nums); i++ {
		last = nums[i] + max(last, 0)
		res = max(res, last) // 在遍历的过程中记录最大值
	}
	return res
}
