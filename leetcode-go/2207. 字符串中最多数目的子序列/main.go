package main

func maximumSubsequenceCount(text, pattern string) (ans int64) {
	x, y := pattern[0], pattern[1]
	cntX, cntY := 0, 0
	for i := range text {
		c := text[i]
		if c == y {
			ans += int64(cntX)
			cntY++
		}
		if c == x {
			cntX++
		}
	}
	return ans + int64(max(cntX, cntY))
}
