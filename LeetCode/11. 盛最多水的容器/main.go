package main

import "math"

func maxArea(height []int) int {
	n := len(height)
	l, r := 0, n-1
	res := math.MinInt32
	for l < r {
		res = max(res, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}