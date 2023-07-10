package main

// 找规律 | dp | 贪心
func maxAlternatingSum(nums []int) int64 {
	res := int64(nums[0])
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 {
			res += int64(diff)
		}
	}
	return res
}
