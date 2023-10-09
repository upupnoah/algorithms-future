package _731__Movement_of_Robots

import "sort"

func sumDistance(nums []int, s string, d int) int {
	const mod = 1_000_000_007
	for i, c := range s {
		nums[i] += int(c&2-1) * d
	}
	sort.Ints(nums)
	res, sum := 0, 0
	for i, x := range nums {
		res = (res + x*i - sum) % mod
		sum += x
	}
	return res
}
