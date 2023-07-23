package main

func trap(height []int) int {
	n := len(height)
	left, right := make([]int, n), make([]int, n)
	left[0], right[n-1] = height[0], height[n-1] // 两侧不能存水
	for i, j := 1, n-2; i < n; i, j = i+1, j-1 {
		left[i] = max(left[i-1], height[i])
		right[j] = max(right[j+1], height[j])
	}
	res := 0
	for i := 0; i < n; i++ {
		res += min(left[i], right[i]) - height[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
