package main

// 思路：记忆化搜索

func minimumtime(n int, relations [][]int, time []int) int {
	res := 0
	prev := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		prev[i] = make([]int, 0)
	}
	for _, relation := range relations {
		x := relation[0]
		y := relation[1]
		prev[y] = append(prev[y], x)
	}
	memo := make(map[int]int)
	for i := 1; i <= n; i++ {
		res = max(res, dp(i, time, prev, memo))
	}
	return res
}

func dp(i int, time []int, prev [][]int, memo map[int]int) int {
	if _, ok := memo[i]; !ok {
		cur := 0
		for _, p := range prev[i] {
			cur = max(cur, dp(p, time, prev, memo))
		}
		cur += time[i-1]
		memo[i] = cur
	}
	return memo[i]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
