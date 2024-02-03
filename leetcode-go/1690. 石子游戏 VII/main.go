package stone_game_vii

func stoneGameVII(stones []int) int {
	n := len(stones)
	dp := make([][]int, n)
	//dp[i][j]表示在区间ij玩家先手能取得的最大分差，dp[i][i]=0,i<=j有意义
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 2; i >= 0; i-- {
		//sum表示区间总和，起始为stones[i]
		sum := stones[i]
		for j := i + 1; j < n; j++ {
			//枚举区间，结束为stones[j]
			sum += stones[j]
			dp[i][j] = max(sum-stones[i]-dp[i+1][j], sum-stones[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1]
}

// 一维优化
func stoneGameVII2(stones []int) int {
	n := len(stones)
	dp := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		sum := stones[i]
		for j := i + 1; j < n; j++ {
			sum += stones[j]
			dp[j] = max(sum-stones[i]-dp[j], sum-stones[j]-dp[j-1])
		}
	}
	return dp[n-1]
}