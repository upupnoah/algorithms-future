package main

func maxArrayValue(nums []int) int64 {
	ans := nums[len(nums)-1]
	for i := len(nums)-1; i >= 1; i-- {
		if ans >= nums[i-1] {
			ans += nums[i-1]
		} else {
			ans = nums[i-1]
		}
	}
	return int64(ans)
}