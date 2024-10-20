package main

import "slices"

func smallestRangeI(nums []int, k int) int {
	return max(slices.Max(nums)-slices.Min(nums)-k*2, 0)
}
