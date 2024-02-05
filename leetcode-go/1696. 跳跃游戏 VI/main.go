package main

import (
	"container/heap"
	"slices"
	"sort"
)

var dp []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return dp[h.IntSlice[i]] > dp[h.IntSlice[j]] } // 加上这行变成最大堆
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

// 方法一：动态规划 + 大根堆(优先队列) O(nlogn)
// dp[i]: 从 0 跳到第 i 个位置的最大得分
func maxResult1(nums []int, k int) int {
	n := len(nums)
	dp = make([]int, n)

	h := &hp{sort.IntSlice{}}
	heap.Push(h, 0)

	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		for h.Len() > 0 && h.IntSlice[0] < i-k {
			heap.Pop(h)
		}
		dp[i] = nums[i] + dp[h.IntSlice[0]]
		heap.Push(h, i)
		// for j := max(0, i-k); j < i; j++ {
		// 	dp[i] = max(dp[i], dp[j]+nums[i])
		// }
	}
	return dp[n-1]
}

// 方法二：动态规划 + 单调队列 O(n)
func maxResult2(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	q := []int{0}
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i], dp[q[0]]+nums[i])
		for len(q) > 0 && dp[q[len(q)-1]] <= dp[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i-q[0] >= k {
			q = q[1:]
		}
	}
	return dp[n-1]
}

// 超时，需要进一步优化
func maxResultTimeOut(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n)
	f[0] = nums[0]
	for i := 1; i < n; i++ {
		f[i] = slices.Max(f[max(i-k, 0):i]) + nums[i]
	}
	return f[n-1]
}
