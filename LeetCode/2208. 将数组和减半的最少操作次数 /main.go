package main

import "container/heap"

// 思路：贪心，堆（每次减半最大值），然后重新塞入堆
type PriorityQueue []float64

func (p PriorityQueue) Len() int {
	return len(p)
}

// 大根堆
func (p PriorityQueue) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.(float64))
}

func (p *PriorityQueue) Pop() any {
	old, n := *p, len(*p)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func halveArray(nums []int) int {
	var sum float64
	p := &PriorityQueue{}
	for _, v := range nums {
		sum += float64(v)
		heap.Push(p, float64(v))
	}
	res, half := 0, sum/2
	for sum > half {
		x := heap.Pop(p).(float64)
		sum -= x / 2
		res++
		heap.Push(p, x/2)
	}
	return res
}
