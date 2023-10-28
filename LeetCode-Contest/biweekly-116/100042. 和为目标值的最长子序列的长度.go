package biweekly_116

// dp[i][j] 表示数组 nums的前 i 个元素,且和为 j 的最长子序列长度
// 初始化 dp 数组为:-1, 表示无法通过选择前'i'个元素得到和为'j'的子序列
// dp[0][0] 为 0, 没有选择任何元素时 和 为[0] 的子序列的长度为 0
// 状态转移:
//
//	包含 nums[i]: dp[i][j] = max(dp[i][j], d[i-1][j-nums[i]]+1)
//	不包含 nums[i]: dp[i][j] = max(dp[i][j], dp[i-1][j])
//
// 返回:dp[n][target]

func lengthOfLongestSubsequence(nums []int, target int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j] // 不包含 nums[i-1]
			if j >= nums[i-1] && dp[i-1][j-nums[i-1]] != -1 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-nums[i-1]]+1) // 包含 nums[i-1]
			}
		}
	}
	return dp[n][target]
}
