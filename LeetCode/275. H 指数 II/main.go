package _75__H_指数_II

// 思路:log(n)算法, 升序排列 ->  二分
// I 是降序排序,然后找 citations[i] < i+1, return i
// II 是升序排序,然后找到最后一个 citations[n-i] >= i

func hIndex(citations []int) int {
	n := len(citations)
	l, r := 0, n
	for l < r {
		mid := l + r + 1>>1
		// if check(), 将整个看为 check
		// 满足要求: l = mid, l左边的肯定都满足要求(包括 mid)
		// 否则: r = mid-1, r右边的肯定都不满足要求(mid 也不满足)
		if citations[n-mid] >= mid {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return r
}
