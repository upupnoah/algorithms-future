package main

// 前后缀分解
func trap(height []int) int {
	n := len(height)
	pre, suf := make([]int, n), make([]int, n)
	pre[0], suf[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		pre[i] = max(height[i], pre[i-1])
		suf[n-1-i] = max(height[n-1-i], suf[n-i])
	}
	res := 0
	for i := 0; i < n; i++ {
		res += min(pre[i], suf[i]) - height[i]
	}
	return res
}

// 双指针
func trap1(height []int) int {
	n := len(height)
	l, r := 0, n -1
	preMax, sufMax := 0, 0
	res := 0
	for l < r {
		preMax, sufMax = max(preMax, height[l]), max(preMax, height[r])
		if preMax < sufMax {
			res += preMax - height[l]
			l++
		} else {
			res += sufMax - height[r]
			r--
		}
	}
	return res
}

// 单调栈
func trap2(height []int) int {
	// todo
	return 0
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
