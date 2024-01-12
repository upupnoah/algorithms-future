package LCP_06__拿硬币

// 求 sum(coin/2) 上取整即可
func minCount(coins []int) int {
	ans := 0
	for _, x := range coins {
		//if x%2 == 0 {
		//	ans += x / 2
		//} else {
		//	ans += x/2 + 1
		//}
		// 可以扩展到 (x + k-1) / k
		// 优先级： +-*/ 的优先级比 >> 低，因此需要加括号
		ans += (x + 1) >> 1 // 忽略奇偶性 求/2 上取整的技巧
	}
	return ans
}
