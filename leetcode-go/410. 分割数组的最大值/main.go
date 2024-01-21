package split_array

import "math"

// f[i][j]: 表示将数组的前 i 个数分割为 j 段所能得到的最大连续子数组和的最小值
func splitArray(nums []int, m int) int {
	n := len(nums)
	f := make([][]int, n+1)
	sub := make([]int, n+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, m+1)
		for j := 0; j < len(f[i]); j++ {
			f[i][j] = math.MaxInt32
		}
	}
	// 经典前缀和(求连续的子数组和)
	for i := 0; i < n; i++ {
		sub[i+1] = sub[i] + nums[i]
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, m); j++ {
			// 枚举 k，其中前 k 个数被分割为 j−1 段，而第 k+1 到第 i 个数为第 j 段
			for k := 0; k < i; k++ {
				f[i][j] = min(f[i][j], max(f[k][j-1], sub[i]-sub[k]))
			}
		}
	}
	return f[n][m]
}
