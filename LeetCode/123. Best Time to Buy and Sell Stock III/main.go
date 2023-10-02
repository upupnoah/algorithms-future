package _23__Best_Time_to_Buy_and_Sell_Stock_III

import "math"

// 题意：最多完成 2 次交易的股票出售题目，k = 2
// 188 IV 题的简化版，可以修改 IV 题的代码即可

// 算法 1：dfs（超时）
// 题意：在股票买卖的基础上，限制了至多 k 次交易
// 思路：在 dfs 的时候增加一个参数 k，表示交易数
// 递归边界： k < 0， 返回负无穷，表示不合法
// 			其他的边界和股票买卖一样

func maxProfit(prices []int) int {
	k := 2
	n := len(prices)
	var dfs func(int, int, int) int
	dfs = func(i, k, hold int) int {
		if k < 0 {
			return math.MinInt / 2
		}
		if i < 0 {
			if hold == 1 {
				return math.MinInt / 2
			}
			return 0
		}
		if hold == 1 {
			return max(dfs(i-1, k, 1), dfs(i-1, k, 0)-prices[i])
		}
		return max(dfs(i-1, k, 0), dfs(i-1, k-1, 1)+prices[i])
	}
	return dfs(n-1, k, 0)
}

// 算法 2：dfs + 记忆化
func maxProfit1(prices []int) int {
	k := 2
	n := len(prices)
	memo := make([][][2]int, n)
	for i := range memo {
		memo[i] = make([][2]int, k+1)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, k, hold int) (res int) {
		if k < 0 {
			return math.MinInt / 2 // 避免 -的时候溢出
		}
		if i < 0 {
			if hold == 1 {
				return math.MinInt / 2
			}
			return 0
		}
		p := &memo[i][k][hold]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if hold == 1 {
			return max(dfs(i-1, k, 1), dfs(i-1, k, 0)-prices[i])
		}
		return max(dfs(i-1, k, 0), dfs(i-1, k-1, 1)+prices[i])
	}
	return dfs(n-1, k, 0)
}

// 算法 3：DP
func maxProfit2(prices []int) int {
	k := 2
	n := len(prices)
	f := make([][][2]int, n+1)
	for i := range f {
		f[i] = make([][2]int, k+2)
		for j := range f[i] {
			f[i][j] = [2]int{math.MinInt / 2, math.MinInt / 2}
		}
	}
	for j := 1; j <= k+1; j++ {
		f[0][j][0] = 0
	}
	for i, p := range prices {
		for j := 1; j <= k+1; j++ {
			f[i+1][j][0] = max(f[i][j][0], f[i][j][1]+p)
			f[i+1][j][1] = max(f[i][j][1], f[i][j-1][0]-p)
		}
	}
	return f[n][k+1][0]
}

// 算法 4：DP + 空间优化
// 思路：因为 f[i+1] 只会用到 f[i], 因此可以去掉一维
func maxProfit3(prices []int) int {
	k := 2
	f := make([][2]int, k+2)
	for j := 1; j <= k+1; j++ {
		f[j][1] = math.MinInt / 2
	}
	f[0][0] = math.MinInt32 / 2
	for _, p := range prices {
		// 类似于 0/1 背包，倒序枚举
		for j := k + 1; j >= 1; j-- {
			//f[i+1][j][0] = max(f[i][j][0], f[i][j][1]+p)
			//f[i+1][j][1] = max(f[i][j][1], f[i][j-1][0]-p)
			// 计算 f[i+1]的时候， 这实际上是 f[i]这一层的
			// 又因为 j-1 < j，因此会被先计算
			// 如果正序枚举的话，这里的 f[j-1][0] 是 f[i+1]这一层的，但是我们要的是f[i]这一层的
			// 因此可以倒序枚举，这样这里的 f[j-1][0]是我们需要的f[i]这一层的
			f[j][1] = max(f[j][1], f[j-1][0]-p)
			f[j][0] = max(f[j][0], f[j][1]+p)
		}
	}
	return f[k+1][0]
}
