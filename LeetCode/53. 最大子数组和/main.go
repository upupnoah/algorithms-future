package main

// f[i] = max(f[i-1]+nums[i], nums[i])
// 因为要连续，因此以 i 结尾的最大自序和，要么是以 i-1 结尾的最大自序和加上 nums[i]，要么是 nums[i] 本身
func maxSubArray(nums []int) int {
	n := len(nums)
	f := make([]int, n+1)
	f[0] = nums[0]
	ans := f[0]
	for i := 1; i < n; i++ {
		f[i] = max(nums[i], f[i-1]+nums[i])
		ans = max(ans, f[i])
	}
	return ans
}
