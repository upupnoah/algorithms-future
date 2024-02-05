package max_sliding_window

import (
	"container/heap"
	"sort"
)

var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return a[h.IntSlice[i]] > a[h.IntSlice[j]] } // 加上这行变成最大堆
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

// func (h *hp) push(v int)        { heap.Push(h, v) }
// func (h *hp) pop() int          { return heap.Pop(h).(int) } // 稍微封装一下，方便使用

// 优先队列
// heap 中存索引, less 根据索引对应的值大小排序
// 取出 out of window 的元素只需要关心堆顶元素是否在滑动窗口中!!!
func maxSlidingWindow1(nums []int, k int) []int {
	a = nums // 因为 Less 方法要用到 nums 数组，所以把它设为全局传递出去
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	// underlying array = n-k+1, size = 1, 方便后面直接 append, 而不需要(索引 && 扩容)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k { // 如果堆顶元素不在滑动窗口中，就弹出
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

// 单调队列
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, 0, n-k+1)
	q := []int{}
	for i, x := range nums {
		// 维护单调性 (从左到右, 严格递减)
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1]
		}
		// 入队
		q = append(q, i)
		// out of window
		if i-q[0] >= k {
			q = q[1:] // slice O(1)
		}
		// save ans
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
	}
	return ans
}

// 分块 + 预处理(前后缀最大值)
// 划分为 n/k 个块, 每个块的大小为 k, 最后一个块可能小于 k
func maxSlidingWindow3(nums []int, k int) []int {
	n := len(nums)
	pre, suf := make([]int, n), make([]int, n)
	for i, v := range nums {
		if i%k == 0 {
			pre[i] = v
		} else {
			pre[i] = max(pre[i-1], v)
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || (i+1)%k == 0 {
			suf[i] = nums[i]
		} else {
			suf[i] = max(suf[i+1], nums[i])
		}
	}

	ans := make([]int, n-k+1)
	for i := range ans {
		ans[i] = max(suf[i], pre[i+k-1])
	}
	return ans
}
