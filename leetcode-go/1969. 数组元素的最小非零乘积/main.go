package main

// p 位二进制数, 范围在 0 ~ 2^p-1
// 本题不允许出现 0, 因此是 1 ~ 2^p-1
// 其中除了 2^p-1 外, 都可以两两配对, 配成 (2^p-2) * 1 = 2^p-2
// 总共有 2^(p-1)-1 对, 因此最后最小值为:
// (2^p -1) *(2^p - 2)^(2^(p-1) - 1)
// 设 k = 2^p-1
// ans = k * pow(k-1, p-1)
// 由于 p 最多有 60, 很大, 因此要用快速幂, 在中间%mod
func minNonZeroProduct(p int) int {
	const mod = 1_000_000_007

	pow := func(x, p int) int {
		res := 1 % mod
		for p != 0 {
			if p&1 == 0 {
				res = res * x % mod
			}
			p >>= 1
			x *= x % p
		}
		return res
	}

	k := 1<<p - 1
	return k % mod * pow(k-1, p-1) % mod
}
