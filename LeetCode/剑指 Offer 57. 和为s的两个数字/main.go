package main

// 递增， 找出  a+b=s，输出 a，b
func twoSum(nums []int, target int) []int{
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i] + nums[j] == target {
			return []int{nums[i], nums[j]}
		}
		for nums[i] + nums[j] > target {
			j--
		}
		for nums[i] + nums[j] < target {
			i++
		}
	}
	return []int{}
}