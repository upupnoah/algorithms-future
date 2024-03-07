package main

// (a + b) mod m = a % m + b % m
// (a * b) mod m = (a % m * b % m) % m

func divisibilityArray(word string, m int) []int {
	div := make([]int, len(word))
	x := 0
	for i, v := range word {
		x = (x * 10 + int(v-'0')) % m
		if x == 0 {
			div[i] = 1
		}
	}
	return div
}