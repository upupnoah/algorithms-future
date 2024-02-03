package predict_the_winner

// dfs
func predictTheWinner(nums []int) bool {
	n := len(nums)
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == j {
			return nums[i]
		}
		// 当前分数 - 下轮对手的分数
		cl := nums[i] - dfs(i+1, j)
		cr := nums[j] - dfs(i, j-1)
		if cl > cr {
			return cl
		}
		return cr
	}
	return dfs(0, n-1) >= 0
}

// memo dfs
func predictTheWinner2(nums []int) bool {
	n := len(nums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 << 63
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == j {
			return nums[i]
		}
		p := &memo[i][j]
		if *p != -1<<63 {
			return *p
		}
		cl := nums[i] - dfs(i+1, j)
		cr := nums[j] - dfs(i, j-1)
		if cl > cr {
			*p = cl
		} else {
			*p = cr
		}
		return *p
	}

	return dfs(0, n-1) >= 0
}

// dp
func predictTheWinner3(nums []int) bool {
	n := len(nums)
	// dp[i][j]: 表示 局面[i,j] 中先手, 赢过对方的分数
	// 参考记忆化搜索: dp[i][j] = max(nums[i] - dp[i+1][j], nums[j] - dp[i][j-1])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}
	// 参考状态转移方程, i要用到 i+1, 因此 i 倒着遍历, 并且从 n-2 开始(越界问题)
	// j 要用到 j-1, 因此 j 正着遍历, 并且从 i+1 开始(i=0时, 如果 j 从 i 开始, j-1 = -1 越界), dp[i][i] 已经计算过了
	for i := n - 2; i >= 0; i-- {
		// i <= j, 这里 dp[i][i] 已经计算过了, 因此 j 从 i+1 开始
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] >= 0
}
