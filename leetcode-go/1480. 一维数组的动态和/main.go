package main

// 前缀和
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1]+nums[i]
	}
	return nums
}