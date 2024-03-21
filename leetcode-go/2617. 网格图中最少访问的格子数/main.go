package main

import (
	"container/heap"
	"math"
)

func minimumVisitedCells(grid [][]int) int {
	colHeaps := make([]hp, len(grid[0])) // 每一列的最小堆
	rowH := hp{}                         // 行最小堆
	f := 1                               // 起点算 1 个格子
	for i, row := range grid {
		rowH = rowH[:0]
		for j, g := range row {
			for len(rowH) > 0 && rowH[0].idx < j { // 无法到达第 j 列
				heap.Pop(&rowH) // 弹出无用数据
			}
			colH := colHeaps[j]
			for len(colH) > 0 && colH[0].idx < i { // 无法到达第 i 行
				heap.Pop(&colH) // 弹出无用数据
			}
			if i > 0 || j > 0 {
				f = math.MaxInt
			}
			if len(rowH) > 0 {
				f = rowH[0].f + 1 // 从左边跳过来
			}
			if len(colH) > 0 {
				f = min(f, colH[0].f+1) // 从上边跳过来
			}
			if g > 0 && f < math.MaxInt {
				heap.Push(&rowH, pair{f, g + j}) // 经过的格子数，向右最远能到达的列号
				heap.Push(&colH, pair{f, g + i}) // 经过的格子数，向下最远能到达的行号
			}
			colHeaps[j] = colH // 注意上面的入堆出堆操作只改了 colH，没有改 colHeaps[j]
		}
	}
	// 此时的 f 是在 (m-1, n-1) 处算出来的
	if f < math.MaxInt {
		return f
	}
	return -1
}

type pair struct{ f, idx int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].f < h[j].f }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
