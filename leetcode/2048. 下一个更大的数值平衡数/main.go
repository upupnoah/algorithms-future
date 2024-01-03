package next_beautiful_number

func solve(n int) bool {
	cnt := make([]int, 10)
	for n != 0 {
		cnt[n%10]++
		n /= 10
	}
	for i := 0; i < 10; i++ {
		if cnt[i] != i && cnt[i] != 0 {
			return false
		}
	}
	return true
}

func nextBeautifulNumber(n int) int {
	for i := n + 1; i <= 1224444; i++ {
		if solve(i) {
			return i
		}
	}
	return 0
}
