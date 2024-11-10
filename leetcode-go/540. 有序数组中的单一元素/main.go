package main

func singleNonDuplicate(nums []int) int {
	left, right := -1, len(nums)/2
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid*2] != nums[mid*2+1] {
			right = mid
		} else {
			left = mid
		}
	}
	return nums[right*2]
}
