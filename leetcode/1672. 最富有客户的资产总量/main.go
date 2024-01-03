package main

// 题意：返回 n 行矩阵中的最大值（对每一行求和）
func maximumWealth(accounts [][]int) int {
	ans := 0
	for _, v := range accounts {
		sum := 0
		for _, vv := range v {
			sum += vv
		}
		ans = max(ans, sum)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}