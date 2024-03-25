package main

// dfs(i,c) 表示用前 i 种硬币, 剩下需要凑的金额 c 的方案数, 考虑选/不选
// 1. 不选第 i 种硬币: dfs(i-1, c)
// 2. 选第 i 种硬币: dfs(i, c-coins[i])
// 综上: dfs(i, c) = dfs(i-1, c) + dfs(i, c-coins[i])
// 递归边界: dfs(-1, 0)=1, dfs(-1, >0) = 0
// 递归入口: dfs(n-1, amount)
func change(amount int, coins []int) int {
	n := len(coins)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, amount+1)
		for j := range memo[i] {
			memo[i][j] = -1 // 表示没有访问过
		}
	}

	var dfs func(int, int) int
	dfs = func(i, c int) (res int) {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}
		p := &memo[i][c]
		if *p != -1 {
			return *p
		}
		defer func() {
			*p = res
		}()
		if c < coins[i] {
			return dfs(i-1, c)
		}
		return dfs(i-1, c) + dfs(i, c-coins[i])
	}
	return dfs(n-1, amount)
}

func change1(amount int, coins []int) int {
	n := len(coins)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	f[0][0] = 1
	for i, x := range coins {
		for c := 0; c <= amount; c++ {
			if c < x {
				f[i+1][c] =  f[i][c]
			} else {
				f[i+1][c] = f[i][c] + f[i+1][c-x]
			}
		}
	}
	return f[n][amount]
}

func change2(amount int, coins []int) int {
	f := make([]int, amount+1)
	f[0] = 1
	for _, x := range coins {
		for c := x; c <= amount; c++ {
				f[c] += f[c-x]
		}
	}
	return f[amount]
}
