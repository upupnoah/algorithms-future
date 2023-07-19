package main

import (
	"container/heap"
	"sort"
)

type IntervalHeap [][]int

func (h *IntervalHeap) Len() int           { return len(*h) }
func (h *IntervalHeap) Less(i, j int) bool { return (*h)[i][0] < (*h)[j][0] }
func (h *IntervalHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntervalHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *IntervalHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	sortedQueries := make([][2]int, len(queries))
	for i, q := range queries {
		sortedQueries[i] = [2]int{q, i}
	}
	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i][0] < sortedQueries[j][0]
	})

	intervalsHeap := &IntervalHeap{}
	heap.Init(intervalsHeap)

	results := make([]int, len(queries))
	i := 0
	for _, q := range sortedQueries {
		queryValue, queryIdx := q[0], q[1]

		for i < len(intervals) && intervals[i][0] <= queryValue {
			intervalStart, intervalEnd := intervals[i][0], intervals[i][1]
			intervalLength := intervalEnd - intervalStart + 1
			heap.Push(intervalsHeap, []int{intervalLength, intervalEnd})
			i++
		}

		for intervalsHeap.Len() > 0 && (*intervalsHeap)[0][1] < queryValue {
			heap.Pop(intervalsHeap)
		}

		if intervalsHeap.Len() > 0 {
			results[queryIdx] = (*intervalsHeap)[0][0]
		} else {
			results[queryIdx] = -1
		}
	}

	return results
}
