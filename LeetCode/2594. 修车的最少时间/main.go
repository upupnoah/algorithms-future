package _594__修车的最少时间

import "math"

// 思路： 二分可能得时间，从 1 — 正无穷（或者任意一个工人修完所有车的时间）
// 二分时间t，r*n*n <= t, 因此 n <= sqrt(t/r), 所有在 t 时间内，该工人最多修 下取整（sqrt(t/r)） 辆车
func repairCars(ranks []int, cars int) int64 {
	minR := ranks[0]
	for _, v := range ranks {
		if v < minR {
			minR = v
		}
	}
	l, r := 0, minR*cars*cars
	for l < r {
		mid := (l + r) >> 1
		sum := 0
		for _, v := range ranks {
			sum += int(math.Sqrt(float64(mid / v)))
		}
		// 要找到满足条件的第一个（最少的时间）
		if sum >= cars {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return int64(r)
}
