package main

import "math"

// 思路：单调队列维护最大值（滑动窗口）
// 做法：求 yi + yj + |xi - xj| 的最大值
// 1. yi + yj + |xi - xj| = yi - xi + yj + xj
// 2. yj + xj 是固定的，因此只需要求 yi - xi 的最大值即可
// 3. yi - xi 的最大值可以用单调队列维护
// 4. 队首元素就是最大值
// 5. 队首元素的 x 值必须在窗口范围内
func findMaxValueOfEquation(points [][]int, k int) int {
	ans := math.MinInt32
	type pair struct{ x, yx int }
	q := []pair{}
	for _, p := range points {
		x, y := p[0], p[1]
		// 超出窗口范围的点出队
		for len(q) > 0 && q[0].x < x-k {
			q = q[1:]
		}
		if len(q) > 0 {
			ans = max(ans, x+y+q[0].yx) // 当前 x，y 加上队首元素的 yx（最大的 yi-yx)
		}
		// 老的必须比新的强，否则就淘汰
		for len(q) > 0 && q[len(q)-1].yx <= y-x {
			q = q[:len(q)-1]
		}
		// 最后加上当前的点
		q = append(q, pair{x, y - x})
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
