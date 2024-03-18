package main

func maximumScore(nums []int, k int) int {
	i, j := k, k
	minVal, ans := nums[k], nums[k]
	n := len(nums)
	for i > 0 || j < n-1 {
		if i == 0 {
			j++
		} else if j == n-1 {
			i--
		} else if nums[i-1] > nums[j+1] {
			i--
		} else {
			j++
		}
		minVal = min(minVal, min(nums[i], nums[j]))
		ans = max(ans, minVal*(j-i+1))
	}
	return ans
}