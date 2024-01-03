package main

import "strings"

func reverseWords(s string) string {
	res := ""
	i, j := len(s)-1, len(s)-1
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
		for i >= 0 && s[i] != ' ' {
			i--
		}
		res += s[i+1:j+1] + " "
	}
	return strings.TrimRight(res, " ")
}
