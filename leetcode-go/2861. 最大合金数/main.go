package max_number_of_all_oys

import (
	"slices"
	"sort"
)

// 最好的情况, 最多可以造 min(stock) + budget/n 个合金
// 二分答案
func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	mx := slices.Min(stock) + budget/n

	check := func(num int) bool {
		// 每台机器(依次判断是否有机器符合要求)
		for _, comp := range composition {
			money := budget
			// 制作一块合金: 所需要的数量的每种类型的金属
			for i, v := range comp {
				if v*num <= stock[i] {
					continue
				}
				money -= (num*v - stock[i]) * cost[i]
			}
			if money >= 0 {
				return true
			}
		}
		return false
	}
	l, r := 0, mx
	for l < r {
		mid := (l + r + 1) >> 1
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func maxNumberOfAlloys2(n, _, budget int, composition [][]int, stock, cost []int) (ans int) {
	mx := slices.Min(stock) + budget/n
	for _, comp := range composition {
		ans += sort.Search(mx-ans, func(num int) bool {
			num += ans + 1
			money := 0
			for i, s := range stock {
				if s < comp[i]*num {
					money += (comp[i]*num - s) * cost[i]
					if money > budget {
						return true
					}
				}
			}
			return false
		})
	}
	return
}
