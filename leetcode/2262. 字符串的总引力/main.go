package _262__字符串的总引力

func appealSum(s string) (ans int64) {
	sumG, last := 0, [26]int{}
	for i := range last {
		last[i] = -1
	}
	for i, c := range s {
		c -= 'a'
		sumG += i - last[c]
		ans += int64(sumG)
		last[c] = i
	}
	return
}
