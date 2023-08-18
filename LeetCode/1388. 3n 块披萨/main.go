package main

// 类似于 打家劫舍II
// dp[i][j] 表示 在前 i 个数中选择了 j 个不相邻的数的最大和
// i < 2 || j = 0 时 dp[i][j] =：
//
//	0, j = 0
//	slices[0], i=0,j=1
//	max(slices[0],slices[1]), i=1,j=1
//	-无穷，i<2, j>=2
func maxSizeSlices(slices []int) int {
	return max(calculate(slices[1:]), calculate(slices[:len(slices)-1]))
}

func calculate(slices []int) int {
	N, n := len(slices), (len(slices)+1)/3
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -0x3f3f3f3f
		}
	}
	dp[0][0], dp[0][1], dp[1][0], dp[1][1] = 0, slices[0], 0, max(slices[0], slices[1])
	for i := 2; i < N; i++ {
		dp[i][0] = 0
		for j := 1; j <= n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-2][j-1]+slices[i])
		}
	}
	return dp[N-1][n]
}
