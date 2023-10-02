package _22__Best_Time_to_Buy_and_Sell_Stock_II

import "math"

// 算法1：贪心
// 思路：交易的拆解
//
//	在 pi 天买入，pj 天卖出可以看成在 pi天买入，pi+1 天卖出，在 pi+1 天买入，在 pi+2 天卖出...
func maxProfit(prices []int) (ans int) {
	for i := 0; i+1 < len(prices); i++ {
		ans += max(0, prices[i+1]-prices[i])
	}
	return
}

// 算法 2：dfs （超时）
// 子问题：
//   - 第 i 天结束时的利润 等价于 第 i-1 天结束时的利润 + 第 i 天的收益（啥也不干，买入，卖出）
//   - 可以使用状态机来表示
//
// 递归边界：
//   - dfs(-1, 0)  第 0 天开始未持有股票，利润为 0
//   - dfs(-1, 1)  第 0 天开始不可能持有股票，初始化为负无穷（取 max 的时候一定取不到这个状态）
//
// 递归入口：
//   - max(dfs(n-1,0), dfs(n-1,1)
//     = dfs(n-1, 0) // 如果最后一天还持有股票，那一定是不划算的
func maxProfit1(prices []int) int {
	n := len(prices)
	var dfs func(int, bool) int
	dfs = func(i int, hold bool) int {
		if i < 0 {
			if hold {
				return math.MinInt / 2
			}
			return 0
		}
		if hold { // 如果要计算持有股票的状态：max(啥也不干，买入股票）
			return max(dfs(i-1, true), dfs(i-1, false)-prices[i])
		}
		// 未持有股票的状态：max(啥也不干，卖出股票）
		return max(dfs(i-1, false), dfs(i-1, true)+prices[i])
	}
	return dfs(n-1, false)
}

// 算法 3：dfs+记忆化搜索
func maxProfit2(prices []int) int {
	n := len(prices)
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1}
	}
	var dfs func(int, int) int
	dfs = func(i, hold int) (res int) {
		if i < 0 {
			if hold == 1 {
				return math.MinInt / 2
			}
			return 0
		}
		p := &memo[i][hold]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if hold == 1 { // 如果要计算持有股票的状态：max(啥也不干，买入股票）
			return max(dfs(i-1, 1), dfs(i-1, 0)-prices[i])
		}
		// 未持有股票的状态：max(啥也不干，卖出股票）
		return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])
	}
	return dfs(n-1, 0)
}

// 算法 4：dp
func maxProfit3(prices []int) int {
	n := len(prices)
	f := make([][2]int, len(prices))
	// 初始化边界，整体右移一位
	f[0][1] = math.MinInt / 2
	for i, p := range prices {
		f[i+1][0] = max(f[i][0], f[i][1]+p)
		f[i+1][1] = max(f[i][1], f[i][0]-p)
	}
	return f[n][0]
}

// 算法 5：dp + 优化空间
// - 考虑到每次只用到了 f[i+1] 和 f[i], 因此可以用两个变量来表示
func maxProfit4(prices []int) int {
	f0, f1 := 0, math.MinInt/2
	for _, p := range prices {
		//newF0 := max(f0, f1+p)
		//f1 = max(f1, f0-p)
		//f0 = newF0
		f0, f1 = max(f0, f1+p), max(f1, f0-p)
	}
	return f0
}
