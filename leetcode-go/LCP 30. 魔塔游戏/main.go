package main

import (
	"container/heap"
	"sort"
)

type hp struct{ sort.IntSlice }

// func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 加上这行变成最大堆
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) push(v int) { heap.Push(h, v) }
func (h *hp) pop() int   { return heap.Pop(h).(int) } // 稍微封装一下，方便使用
func (h *hp) top() int {
	return h.IntSlice[0]
}
func magicTower(nums []int) int {
	h := &hp{}
	cur := 1
	ans := 0
	delay := 0
	for i := range nums {
		if nums[i] < 0 {
			h.push(nums[i])
		}
		cur += nums[i]
		if cur <= 0 {
			ans++
			// 将前面遇到的最大的负数加回来(相当于放到最后)
			delay += h.IntSlice[0] // 放到最后
			cur -= h.pop()         // 先加回来
		}
	}
	if cur+delay <= 0 {
		return -1
	}
	return ans
}
