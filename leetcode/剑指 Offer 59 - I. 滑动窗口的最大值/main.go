package main

// 暴力
func maxSlidingWindow(nums []int, k int ) []int {
	if len(nums) == 0 {
		return []int{}
	}
	var res []int
	for i := 0; i < len(nums) - k + 1; i++ {
		max := nums[i]
		for j := i; j < i + k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		res = append(res, max)
	}
	return res
}

// 滑动窗口
func maxSlidingWindow2(nums []int, k int ) []int {
	var res []int
	var q []int
	for i := 0; i < len(nums); i++ {
		if i >= k && q[0] <= i-k {
			q = q[1:]
		}
		for len(q) > 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i >= k -1 {
			res = append(res, nums[q[0]])
		}
	}
	return res
}