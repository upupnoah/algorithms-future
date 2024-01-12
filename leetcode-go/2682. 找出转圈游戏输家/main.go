package main

func circularGameLosers(n, k int) []int {
	st := make([]bool, n+1)
	cur, cnt := 1, k
	for {
		if st[cur] {
			break
		}
		st[cur] = true
		cur = (cnt + cur) % n
		if cur == 0 {
			cur = n
		}
		cnt += k
	}
	var ans []int
	for i := 1; i < n+1; i++ {
		if !st[i] {
			ans = append(ans, i)
		}
	}
	return ans
}
