package is_acronym

func isAcronym(words []string, s string) bool {
	if len(s) != len(words) {
		return false
	}
	for i, v := range words {
		if v[0] != s[i] {
			return false
		}
	}
	return true
}