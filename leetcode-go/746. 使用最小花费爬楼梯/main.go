package min_cost_climbing_stairs


// dp[i] 表示爬到第 i 级台阶的最小花费
// dp[i] = min(dp[i-1] + cost[i-1], dp[i-2]+cost[i-2])
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2])
	}
	return dp[n]
}