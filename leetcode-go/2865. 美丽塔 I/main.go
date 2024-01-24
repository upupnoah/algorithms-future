package maximum_sum_of_heights

// 直接使用美丽塔 II 的代码(单调栈)
func maximumSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	stk := []int{n}
	sum := 0
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		x := maxHeights[i]
		for len(stk) > 1 && x <= maxHeights[stk[len(stk)-1]] {
			j := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			sum -= maxHeights[j] * (stk[len(stk)-1] - j)
		}
		sum += x * (stk[len(stk)-1] - i)
		stk = append(stk, i)
		suf[i] = sum
	}
	ans := sum

	pre := 0
	stk = []int{-1}
	for i, v := range maxHeights {
		for len(stk) > 1 && v <= maxHeights[stk[len(stk)-1]] {
			j := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			pre -= maxHeights[j] * (j - stk[len(stk)-1])
		}
		pre += v * (i - stk[len(stk)-1])
		stk = append(stk, i)
		ans = max(ans, pre + suf[i+1])
	}
	return int64(ans)
}
