package find_rotate_steps

import "math"

func findRotateSteps(ring string, key string) int {
	// dp[i][j]: 在匹配 key 中第 i 个字母的时候, 使用 ring 中第 j 个位置, 所需要的最小步数
	n, m := len(ring), len(key)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		// 初始化为最大值
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	// 记录 ring 中每个字母出现的位置
	pos := [26][]int{}
	for i, ch := range ring {
		pos[ch-'a'] = append(pos[ch-'a'], i)
	}

	// 第一个字符, 使用 ring 中 第 p 个位置, 所需要的最小步数
	for _, p := range pos[key[0]-'a'] {
		dp[0][p] = min(p, n-p) + 1 // 1 是 按下按钮
	}
	for i := 1; i < m; i++ {
		for _, j := range pos[key[i]-'a'] { // 当前字符在 ring 中的位置
			for _, k := range pos[key[i-1]-'a'] { // 上一个字符在 ring 中的位置
				dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), n-abs(j-k))+1)
			}
		}
	}
	return min(dp[m-1]...)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}
