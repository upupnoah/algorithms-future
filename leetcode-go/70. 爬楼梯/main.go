package climb_stairs

// 动态规划
func climbStairs1(n int) int {
	f := make([]int, n+5)
	f[1], f[2] = 1, 2
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// 动态规划 + 空间优化
func climbStairs(n int) int {
	p, q, r := 0, 0, 1 // p, q, r 分别表示 f[i-2], f[i-1], f[i]
	for i := 1; i <= n; i++ {
		p, q = q, r
		r = p + q
	}
	return r
}
