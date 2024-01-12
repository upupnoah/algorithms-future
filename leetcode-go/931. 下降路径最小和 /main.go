package main

import "math"

// 做法：动态规划 -> 类似于数字三角形
// matrix[i][j] 这个位置只能从 3 个位置走过来：
// 1. matrix[i-1][j-1]
// 2. matrix[i-1][j]
// 3. matrix[i-1][j+1]
// 因此，我们可以得到状态转移方程：
// dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
func minFallingPathSum(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = math.MaxInt // 因为要取最小值，初始化为 int 的最大值
		}
	}
	for i := 0; i < m; i++ {
		dp[0][i] = matrix[0][i] // 第一行的每个位置到当前位置的最短路线都是自己的值
	}
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			// 遍历 3 个位置
			for k := max(0, j-1); k <= min(m-1, j+1); k++ {
				dp[i][j] = min(dp[i][j], dp[i-1][k]+matrix[i][j])
			}
		}
	}
	res := math.MaxInt
	for i := 0; i < m; i++ {
		res = min(res, dp[n-1][i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
