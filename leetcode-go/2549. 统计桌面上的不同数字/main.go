package main

// x % i == 1
// x mod (x-1) == 1
// 因此最后有 n-1 个
// 例如 5 -> 5, 4, 3, 2
// 1 很特殊, 因为任何数%1 都为 0, 因此不可能有 1
func distinctIntegers(n int) int {
	return max(n-1, 1)
}