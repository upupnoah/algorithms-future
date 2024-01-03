package main

// 题意：奇数在前，偶数在后
// 做法：双指针左右存
func exchange(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if nums[i] % 2 == 0 {
			ans[r] = nums[i]
			r--
		} else {
			ans[l] = nums[i]
			l++
		}
	}
	return ans
}