package main

// 暴力 O(n^2)
func findTheLongestBalancedSubstring1(s string) (ans int) {
	n := len(s)
	check := func(str string) (res int) {
		i, j := 0, len(str)-1
		for i < j {
			if str[i] == '0' && str[j] == '1' {
				res += 2
			} else {
				return 0
			}
			i++
			j--
		}
		return
	}
	for i := 1; i <= n; i++ {
		for j := 0; j+i-1 < n; j++ {
			sub := s[j : j+i]
			ans = max(ans, check(sub))
		}
	}
	return
}

// O(N)
func findTheLongestBalancedSubstring(s string) (ans int) {
	c0, c1 := 0, 0
	for i, ch := range s {
		if ch == '0' {
			if i == 0 || s[i-1] == '1' {
				c0, c1 = 1, 0
			} else {
				c0++
			}
		} else {
			c1++
			ans = max(ans, min(c0, c1)*2)
		}
	}
	return
}
