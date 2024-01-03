package main

func maxProduct(words []string) (ans int) {
	mask := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			mask[i] |= 1 << (ch-'a')
		}
	}
	for i :=0; i < len(words)-1; i++ {
		for j := i+1; j < len(words); j++ {
			if mask[i] & mask[j] == 0{
				ans = max(ans, len(words[i]) * len(words[j]))
			}
		}
	}
	return
}
