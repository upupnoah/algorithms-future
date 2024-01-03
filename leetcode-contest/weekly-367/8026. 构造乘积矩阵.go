package main

// 类似于 leetcode 238. 除自身以外数组的乘积

func constructProductMatrix(grid [][]int) [][]int {
	const mod = 12345
	n, m := len(grid), len(grid[0])
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	suf := 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			ans[i][j] = suf
			suf = suf * grid[i][j] % mod
		}
	}
	pre := 1
	for i, row := range grid {
		for j, x := range row {
			ans[i][j] = ans[i][j] * pre % mod
			pre = pre * x % mod
		}
	}
	return ans
}
