package main

import "fmt"

func validStrings(n int) (ans []string) {
	mask := 1<<n - 1
	for x := range 1 << n {
		if x>>1&x == 0 {
			ans = append(ans, fmt.Sprintf("%0*b", n, x^mask))
		}
	}
	return
}
