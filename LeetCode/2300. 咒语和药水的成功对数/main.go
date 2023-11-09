package main

import "sort"

// 排序 + 二分
func successfulPairs(spells []int, potions []int, success int64) (ans []int) {
	sort.Ints(potions)
	n := len(potions)
	for _, s := range spells {
		l, r := 0, n // r = n 使 r 指向一个虚拟的位置，这样若是找不到，最后会停在 n 的位置
		for l < r {
			mid := l + r>>1
			if int64(s) > (success-1)/int64(potions[mid]) {
				r = mid
			} else {
				l = mid + 1
			}
		}
		ans = append(ans, n-l)
	}
	return
}
