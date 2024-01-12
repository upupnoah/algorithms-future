package count_words

func countWords(words1, words2 []string) int {
	cnt1, cnt2 := make(map[string]int, 0), make(map[string]int, 0)
	for _, word := range words1 {
		cnt1[word]++
	}
	for _, word := range words2 {
		cnt2[word]++
	}
	ans := 0
	for k, v := range cnt1 {
		if u, ok := cnt2[k]; ok && v == 1 && u == 1 {
			ans++
		}
	}
	return ans
}
