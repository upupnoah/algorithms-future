package _560__打家劫舍_IV

import "sort"

// 题意：
//
//	给定数组 nums，从中选择一个长度至少为 k 的子序列 A，满足 A 中没有任何元素在 nums 中是相邻的
//	min(max(A), 找一个最小的 max(A)
//
// 套路：最小化最大值，最大化最小值 -> 二分答案

// 本题：
//
//	「偷走的最大金额」越小，能偷的房子越少，反之越多
//	例如 nums = 【1，4，2，3】，在最大金额为 2 时，只能偷 1 和 2
//
// 一般地，二分的值越小，越不能/能满足要求；二分的值越大，越能/不能满足要求。有单调性的保证，就可以二分答案了
func minCapability(nums []int, k int) int {
	// 在[0~1e9]中找一个答案
	return sort.Search(1e9+1, func(mx int) bool {
		f0, f1 := 0, 0
		// 定义 f[i] 表示从 nums[0]到 nums[i]中偷金额不超过 mx的房屋，最多能偷多少间房屋
		// 如果 f[n−1]≥k则表示答案至多为 mx，否则表示答案必须超过 mx
		for _, x := range nums {
			if x <= mx {
				f0, f1 = f1, max(f1, f0+1)
			} else {
				f0 = f1
			}
		}
		return f1 >= k
	})
}
