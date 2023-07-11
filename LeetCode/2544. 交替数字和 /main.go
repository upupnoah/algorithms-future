package main

// 我的做法
func alternateDigitSum(n int) int {
	t := n
	cnt := 0
	for t != 0 {
		cnt++
		t /= 10
	}
	res := 0
	flag := 1
	if cnt%2 == 0 {
		flag = -1
	}
	for n != 0 {
		res += flag * (n % 10)
		n /= 10
		flag *= -1
	}
	return res
}

// 题解做法
func alternateDigitSum2(n int) int {
	res, sign := 0, 1
	for n > 0 {
		res += n % 10 * sign
		sign *= -1
		n /= 10
	}
	return -sign * res
}
