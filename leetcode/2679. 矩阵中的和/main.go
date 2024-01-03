package main

import (
	"container/heap"
	"sort"
)

// 自己的想法：排序，然后找每列的最大值求和
// 时间复杂度：O(nmlogn)
// 空间复杂度：O(nlogm)
func matrixSum(nums [][]int) int {
	n, m := len(nums), len((nums[0]))
	for i := 0; i < n; i++ {
		sort.Ints(nums[i])
	}
	ans := 0
	for i := 0; i < m; i++ {
		t := -1
		for j := 0; j < n; j++ {
			t = max(t, nums[j][i])
		}
		ans += t
	}
	return ans
}

// 另一种参考的方法：给每一行都 build 一个大根堆
// 然后每次取堆顶的元素，然后更新堆顶元素（m 次）
// 时间复杂度：O(nmlogm)
// 空间复杂度：O(mn)
// 建小根堆，然后插入负数，这样就是大根堆了
func matrixSum2(nums [][]int) int {
	res := 0
	n, m := len(nums), len(nums[0])
	pq := make([]IntHeap, n)

	// 为每一行建堆
	for i := 0; i < n; i++ {
		pq[i] = make(IntHeap, 0)
		heap.Init(&pq[i])
		for j := 0; j < m; j++ {
			heap.Push(&pq[i], -nums[i][j])
		}
	}

	for i := 0; i < m; i++ {
		maxVal := -1
		for j := 0; j < n; j++ {
			t := -heap.Pop(&pq[j]).(int)
			maxVal = max(maxVal, t)
		}
		res += maxVal
	}

	return res
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
