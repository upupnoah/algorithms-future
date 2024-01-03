package _652__倍数求和

// 方法一： 枚举

func sumOfMultiples(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			ans += i
		}
	}
	return ans
}

// 方法二： 容斥原理

func s(n, m int) int {
	return (1 + n/m) * (n / m) / 2 * m
}

func sumOfMultiples2(n int) int {
	return s(n, 3) + s(n, 5) + s(n, 7) - s(n, 15) - s(n, 21) - s(n, 35) + s(n, 105)
}
