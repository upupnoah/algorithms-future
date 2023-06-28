package main

import (
	"math"
	"math/bits"
)

// 题意
//  1. 给定整数数组 nums，以及 k，要求将 nums 划分为 k 个子集，每个子集中不包含重复数字
//     如果有则直接返回-1
//  2. 在 1 条件成立的情况下，求子集的最大不兼容性（最大值-最小值）的和的最小值
// 1 <= k <= nums.length <= 16
// nums.length 能被 k 整除。
// 1 <= nums[i] <= nums.length

func minimumIncompatibility(nums []int, k int) int {
	n := len(nums)
	groupLen := n / k
	// 鸽巢原理，判断是否满足条件 1
	// n+1 个相同数字要放入 n 个容器，则肯定会有 1 个容器超过 1 个数字（重复）
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
		if counter[num] > k {
			return -1
		}
	}
	// init
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	values := make(map[int]int)

	// 遍历每个子集
	for mask := 1; mask < (1 << n); mask++ {
		if bits.OnesCount(uint(mask)) != groupLen {
			continue
		}
		minVal, maxVal := 20, 0
		cur := make(map[int]bool)
		// 判断当前子集的每一位
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				if cur[nums[i]] {
					break
				}
				// 更新当前集合最大值和最小值
				cur[nums[i]] = true
				minVal, maxVal = min(minVal, nums[i]), max(maxVal, nums[i])
			}
		}
		// 如果有 groupLen 位 1，那么就表示有 groupLen 个数字（满足划分）
		if len(cur) == groupLen {
			values[mask] = maxVal - minVal // 当前子集的最大不兼容性
		}
	}

	for mask := 0; mask < (1 << n); mask++ {
		if dp[mask] == math.MaxInt32 {
			continue
		}
		seen := make(map[int]int)
		for i := 0; i < n; i++ {
			if (mask & (1 << i)) == 0 {
				seen[nums[i]] = i
			}
		}
		if len(seen) < groupLen {
			continue
		}
		sub := 0
		for _, v := range seen {
			sub |= (1 << v)
		}
		nxt := sub
		for nxt > 0 {
			if val, ok := values[nxt]; ok {
				dp[mask|nxt] = min(dp[mask|nxt], dp[mask]+val)
			}
			nxt = (nxt - 1) & sub
		}
	}
	if dp[(1<<n)-1] < math.MaxInt32 {
		return dp[(1<<n)-1]
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
