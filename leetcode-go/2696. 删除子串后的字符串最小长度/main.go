package min_length

func minLength(s string) int {
	stk := []byte{}
	for i := range s {
		stk = append(stk, s[i])
		m := len(stk)
		if m >= 2 && (string(stk[m-2:]) == "AB" || string(stk[m-2:]) == "CD") {
			stk = stk[:m-2]
		}
	}
	return len(stk)
}
