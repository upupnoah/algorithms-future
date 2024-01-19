package split_words_by_separator

func splitWordsBySeparator(words []string, separator byte) []string {
	ans := make([]string, 0)

	for _, word := range words {
		ans = append(ans, splitWord(word, separator)...)
	}
	return ans
}

func splitWord(word string, separator byte) []string {
	res := make([]string, 0)
	start := 0
	for i := 0; i < len(word); i++ {
		if word[i] == separator {
			if start == i {
				start++
				continue
			}
			res = append(res, word[start:i])
			start = i + 1
		}
	}
	if start < len(word) {
		res = append(res, word[start:])
	}
	return res
}
