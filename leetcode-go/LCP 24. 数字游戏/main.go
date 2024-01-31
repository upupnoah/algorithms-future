package main

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() any {
	n := len(h.IntSlice)
	x := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}

func (h *hp) top() int {
	return h.IntSlice[0]
}

type minHp struct {
	*hp
}

func (h *minHp) Less(i, j int) bool {
	return h.hp.Less(i, j)
}

type maxHp struct {
	*hp
}

func (h *maxHp) Less(i, j int) bool {
	return h.hp.Less(j, i)
}

func numsGame(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	lower, upper := &maxHp{&hp{}}, &minHp{&hp{}}
	lowerSum, upperSum := 0, 0
	mod := int(1e9 + 7)
	for i := 0; i < n; i++ {
		x := nums[i] - i
		if lower.Len() == 0 || lower.top() >= x {
			lowerSum += x
			heap.Push(lower, x)
			if lower.Len() > upper.Len()+1 {
				upperSum += lower.top()
				heap.Push(upper, lower.top())
				lowerSum -= heap.Pop(lower).(int)
			}
		} else {
			upperSum += x
			heap.Push(upper, x)
			if lower.Len() < upper.Len() {
				lowerSum += upper.top()
				heap.Push(lower, upper.top())
				upperSum -= heap.Pop(upper).(int)
			}
		}
		if (i+1)%2 == 0 {
			res[i] = int((upperSum - lowerSum) % mod)
		} else {
			res[i] = (upperSum - lowerSum + lower.top()) % mod
		}
	}
	return res
}
