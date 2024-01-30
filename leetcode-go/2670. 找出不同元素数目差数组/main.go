package distinct_difference_array

func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	suf := make([]int, n)
	diff := make([]int, n)
	st := map[int]int{}
	for i := n - 1; i >= 0; i-- {
		suf[i] = len(st)
		st[nums[i]]++
	}

	st = map[int]int{}
	pre := 0
	for i := 0; i < n; i++ {
		st[nums[i]]++
		pre = len(st)
		diff[i] = pre - suf[i]
	}
	return diff
}
