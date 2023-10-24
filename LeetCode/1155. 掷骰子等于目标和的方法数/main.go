package _155__掷骰子等于目标和的方法数

// 阶段，状态，决策，
// 从 n 组物品中选，每组有 k 个物品每组物品选一个，使得选出的物品的和为 target
// 物品的值为 1~k
// dp[i][j]: 从前 i 个物品中选，使得和为 j 的方案数
// dp[0][0] = 1
// dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+

func numRollsToTarget(n int, m int, target int) int {
	const mod = 1_000_000_007
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, target+1)
	}
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= target; j++ {
			for k := 1; k <= m; k++ {
				if j >= k {
					f[i][j] = (f[i][j] + f[i-1][j-k]) % mod
				}
			}
		}
	}
	return f[n][target]
}
