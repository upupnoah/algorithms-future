package main

import "math"

var memo [1001]int

func twoEggDrop(n int) int {
	if n == 0 {
		return 0
	}
	p := &memo[n]
	if *p > 0 { // 之前计算过
		return *p
	}
	res := math.MaxInt
	for j := 1; j <= n; j++ {
		res = min(res, max(j, twoEggDrop(n-j)+1))
	}
	*p = res // 记忆化
	return res
}
