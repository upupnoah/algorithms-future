package main

import "strings"

func maximumOddBinaryNumber(s string) string {
	cnt1 := 0
	for i := range s {
		if s[i] == '1' {
			cnt1++
		}
	}
	ans := ""
	for i := 0; i < len(s)-1; i++ {
		if cnt1 != 1 {
			ans += "1"
			cnt1--
			continue
		}
		ans += "0"
	}
	ans += "1"
	return ans
}

func maximumOddBinaryNumber0x3f(s string) string {
	cnt1 := strings.Count(s, "1")
	return strings.Repeat("1", cnt1-1) + strings.Repeat("0", len(s)-cnt1) + "1"
}
