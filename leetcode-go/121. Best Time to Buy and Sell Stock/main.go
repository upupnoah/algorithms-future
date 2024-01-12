package _21__Best_Time_to_Buy_and_Sell_Stock

// 算法：枚举
// 思路：从左到右枚举，维护一个 minPrice（从 0 - i-1之间的最小值）
//
//	在遍历的过程中，更新 ans（cur - minPrice
func maxProfit(prices []int) (ans int) {
	minPrice := prices[0]
	for _, p := range prices {
		ans = max(ans, p-minPrice)
		minPrice = min(minPrice, p)
	}
	return
}
