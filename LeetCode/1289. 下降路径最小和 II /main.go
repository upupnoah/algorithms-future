package main

// 思路：数字三角形问题
func minFallingPathSum(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, m)
	}
	// initialize
	for i := 0; i < m; i++ {
		f[0][i] = grid[0][i] // 第一行的 非零偏移下降路径 数字和的最小值为本身
	}
	INF := int(1e8)
	for i := 1; i < n; i++ {
		// 求出最小数和次小数
		d1, d2 := INF, INF
		for j := 0; j < m; j++ {
			x := f[i-1][j]
			if x < d1 {
				d1, d2 = x, d1
			} else if x < d2 {
				d2 = x
			}
		}
		// 更新 f[i][j]
		for j := 0; j < m; j++ {
			if f[i-1][j] == d1 {
				f[i][j] = d2 + grid[i][j]
			} else {
				f[i][j] = d1 + grid[i][j]
			}
		}
	}
	res := INF
	for i := 0; i < m; i++ {
		res = min(res, f[n-1][i])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
