package main

import "math"

// dfs(i, c) 表示: 枚举前 i 个硬币, 凑成总金额所需的最小硬币个数
// 完全背包:
// dfs(i,c) = max(dfs(i-1, c), dfs(i, c-w[i])+v[i])
// 本题是完全背包的变形: 从 i 个物品中选, 可以重复选, 求价值[恰好]等于 capacity 时的最小价值和
// dfs(i,c) = min(dfs(i-1, c), dfs(i, c-w[i])+1)

func coinChange(coins []int, amount int) int {
	n := len(coins)
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, amount+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, c int) (res int) {
		if i < 0 {
			if c == 0 {
				return 0
			}
			return math.MaxInt / 2
		}
		C := &cache[i][c]
		if *C != -1 {
			return *C
		}
		defer func() { *C = res }()
		if c < coins[i] {
			return dfs(i-1, c)
		}
		return min(dfs(i-1, c), dfs(i, c-coins[i])+1)
	}
	ans := dfs(n-1, amount)
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}
