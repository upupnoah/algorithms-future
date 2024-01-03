package main

import (
	"math"
)

func strToInt(str string) int {
	k := 0
	for k < len(str) && str[k] == ' ' {
		k++
	}
	var res int64 = 0

	minus := 1
	if k < len(str) {
		if str[k] == '-' {
			minus = -1
			k++
		} else if str[k] == '+' {
			k++
		}
	}
	for k < len(str) && str[k] >= '0' && str[k] <= '9' {
		res = res*10 + int64(str[k]-'0')
		if res > 1e11 {
			break
		}
		k++
	}

	res *= int64(minus)
	if res > math.MaxInt32 {
		res = math.MaxInt32
	}
	if res < math.MinInt32 {
		res = math.MinInt32
	}
	return int(res)
}
