package main

import "math"

// 前置题目：leetcode53：最大子数组和
// 思路：动态规划
// 分成两类：
// 1. 最大子数组和不跨越数组首尾
// 2. 最大子数组和跨越数组首尾
// 对于第一类，直接使用 leetcode53 的解法即可
// 对于第二类，可以转换成第一类，即：最大子数组和 = 数组总和 - 最小子数组和
// 特殊情况：如果数组总和等于最小子数组和，说明数组中所有元素都是负数，此时最大子数组和就是最大元素
func maxSubarraySumCircular(nums []int) int {
	resMax, resMin := math.MinInt32, 0
	sum := 0
	for i, lastMax, lastMin := 0, 0, 0; i < len(nums); i++ {
		lastMax = nums[i] + max(lastMax, 0)
		resMax = max(resMax, lastMax)

		lastMin = nums[i] + min(lastMin, 0)
		resMin = min(resMin, lastMin)

		sum += nums[i]
	}
	if sum == resMin {
		return resMax
	}
	return max(resMax, sum-resMin)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
