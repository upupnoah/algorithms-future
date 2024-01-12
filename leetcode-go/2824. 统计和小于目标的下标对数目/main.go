package _824__统计和小于目标的下标对数目

import "sort"

func countPairs(nums []int, target int) (ans int) {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum < target {
			ans += j - 1
			i++
		} else {
			j--
		}
	}
	return
}
