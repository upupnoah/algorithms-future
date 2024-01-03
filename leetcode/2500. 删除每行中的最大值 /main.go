package main

import "sort"

// 思路：对每一行进行排序，取每一列的最大值
func deleteGreatestValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		sort.Ints(grid[i])
	}
	res, t := 0, 0
	for i := 0; i < len(grid[0]); i++ {
		t = 0
		for j := 0; j < len(grid); j++ {
			t = max(t, grid[j][i])
		}
		res += t
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
