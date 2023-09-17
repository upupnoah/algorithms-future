package _13__打家劫舍_II

// 思路：考虑是否偷 nums[0]
// 偷 nums[0] -> nums[n-1] 与 nums[1] 不能偷，变成了 nums[2] -> nums[n-2]的问题
// 不偷 nums[0] -> nums[n-1] 与 nums[1] 能偷，变成了 nums[1] -> nums[n-1]的问题
func rob(nums []int) int {
	n := len(nums)
	return max(_rob(nums, 2, n-1)+nums[0], _rob(nums, 1, n))
}

func _rob(nums []int, start, end int) int {
	f0, f1 := 0, 0
	for i := start; i < end; i++ {
		f0, f1 = f1, max(f1, f0+nums[i])
	}
	return f1
}
