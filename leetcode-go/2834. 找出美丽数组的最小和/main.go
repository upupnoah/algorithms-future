package main

// 5 6
// 1 2 3 4 6|
func minimumPossibleSum(n int, target int) int {
	const mod = int(1e9 + 7)
	m := min(target/2, n)
	return (m*(m+1) + (target*2+n-m-1)*(n-m)) / 2 % mod
}
