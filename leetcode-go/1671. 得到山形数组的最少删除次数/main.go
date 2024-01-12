package minimum_mountain_removals

// 山形数组(一段上升, 一段下降)
// 题意: 通过删除元素(最少), 来构造山行数组
// -> 转化为 求以 nums[i] 为结尾的最长上升子序列 和 以 nums[i]为开始的最长下降子序列

// dp[i] = max(dp[i], dp[0..i-1]+1)
func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	lis := make([]int, n+1)
	lds := make([]int, n+1)
	for i := 0; i < n; i++ {
		lis[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				lis[i] = max(lis[i], lis[j]+1)
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		lds[i] = 1
		for j := n - 1; j > i; j-- {
			if nums[i] > nums[j] {
				lds[i] = max(lds[i], lds[j]+1)
			}
		}

	}

	ans := 0
	for i := 1; i < n-1; i++ {
		if lds[i] > 1 && lis[i] > 1 { // 保证是山顶
			ans = max(lds[i]+lis[i]-1, ans)
		}

	}
	return n - ans
}
