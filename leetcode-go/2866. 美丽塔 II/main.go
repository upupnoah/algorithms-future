package maximum_sum_of_heights

// 要点
// 1. 相同的数, 只需要保留最靠左的那一个(的下标)
// 2. 用单调栈维护需要保留的数
// func maximumSumOfHeights(a []int) int64 {
// 	ans := 0
// 	n := len(a)
// 	suf := make([]int, n+1)
// 	stk := []int{n} // 哨兵, 表示第 n 个位置 (方便计算下标)
// 	sum := 0
// 	for i := n - 1; i >= 0; i-- {
// 		x := a[i]
// 		for len(stk) > 1 && x <= a[stk[len(stk)-1]] {
// 			j := stk[len(stk)-1]
// 			stk = stk[:len(stk)-1]
// 			sum -= a[j] * (stk[len(stk)-1] - j)
// 		}
// 		sum += x * (stk[len(stk)-1] - i) // 从 i 到 stk[len(st)-1]-1 都是 x
// 		suf[i] = sum
// 		stk = append(stk, i)
// 	}
// 	ans = sum

// 	stk = []int{-1} // 哨兵, 表示第 -1 个位置
// 	pre := 0
// 	for i, x := range a {
// 		for len(stk) > 1 && x <= a[stk[len(stk)-1]] {
// 			j := stk[len(stk)-1]
// 			stk = stk[:len(stk)-1]
// 			pre -= a[j] * (j - stk[len(stk)-1]) // 撤销之前加到 pre 中的
// 		}
// 		pre += x * (i - stk[len(stk)-1]) // 从 st[len(st)-1] 到 i 都是 x
// 		ans = max(ans, pre+suf[i+1])
// 		stk = append(stk, i)
// 	}
// 	return int64(ans)
// }

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
