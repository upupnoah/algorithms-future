package main

// solution1：双指针
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {

			i++
		}
	}
	return []int{}
}

// solution2：二分
// 对于每一个数字 numbers[i],都找一下是否存在一个数字相加为 target
// nlogn
func twoSum2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		low, high := i+1, len(numbers)-1
		for low <= high {
			mid := (high-low)/2 + low
			if numbers[mid] == target-numbers[i] {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > target-numbers[i] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return []int{-1, -1}
}
