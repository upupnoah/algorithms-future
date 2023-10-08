package _578__Split_With_Minimum_Sum

import "sort"

// 贪心 + 排序

func splitNum(num int) int {
	var nums []int
	for num != 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	sort.Ints(nums)
	num1, num2 := 0, 0
	for i := 0; i < len(nums); i += 2 {
		num1 = num1*10 + nums[i]
		if i+1 < n {
			num2 = num2*10 + nums[i+1]
		}
	}
	return num1 + num2
}
