package _530__执行_K_次操作后的最大分数

import (
	"container/heap"
	"sort"
)

// 每次需要一个最大的分数 -> 堆

func maxKelements(nums []int, k int) int64 {
	h := hp{}
	for _, v := range nums {
		h.push(v)
	}
	var ans int64
	for k != 0 {
		t := h.pop()
		ans += int64(t)
		h.push((t + 2) / 3)
		k--
	}
	return ans
}

// 下面这些都是最小堆

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 加上这行变成最大堆
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) push(v int)        { heap.Push(h, v) }
func (h *hp) pop() int          { return heap.Pop(h).(int) } // 稍微封装一下，方便使用

// EXTRA: 参考 Python，引入下面两个效率更高的方法（相比调用 push + pop）
// replace 弹出并返回堆顶，同时将 v 入堆
// 需保证 h 非空
func (h *hp) replace(v int) int {
	top := h.IntSlice[0]
	h.IntSlice[0] = v
	heap.Fix(h, 0)
	return top
}

// pushPop 将 v 入堆，然后弹出并返回堆顶
// 使用见下面的 dynamicMedians
func (h *hp) pushPop(v int) int {
	if h.Len() > 0 && v < h.IntSlice[0] { // 最大堆改成 v < h.IntSlice[0]
		v, h.IntSlice[0] = h.IntSlice[0], v
		heap.Fix(h, 0)
	}
	return v
}
