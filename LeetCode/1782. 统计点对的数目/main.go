package main

import "sort"

func countPairs(n int, edges [][]int, queries []int) []int {
	// deg[i] 表示与点 i 相连的边的数目
	deg := make([]int, n+1) // 节点编号从 1 到 n
	type edge struct{ x, y int }
	cntE := map[edge]int{}
	for _, e := range edges {
		x, y := e[0], e[1]
		if x > y {
			x, y = y, x
		}
		deg[x]++
		deg[y]++
		// 统计每条边的出现次数，注意 1-2 和 2-1 算同一条边
		cntE[edge{x, y}]++
	}

	ans := make([]int, len(queries))
	sortedDeg := append([]int(nil), deg...)
	sort.Ints(sortedDeg) // 排序，为了双指针
	for j, q := range queries {
		left, right := 1, n // 相向双指针
		for left < right {
			if sortedDeg[left]+sortedDeg[right] <= q {
				left++
			} else {
				ans[j] += right - left
				right--
			}
		}
		for e, c := range cntE {
			s := deg[e.x] + deg[e.y]
			if s > q && s-c <= q {
				ans[j]--
			}
		}
	}
	return ans
}
