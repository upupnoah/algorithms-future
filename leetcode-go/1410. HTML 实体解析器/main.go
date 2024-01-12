package main

func entityParser(text string) string {
	st := make(map[string]string)
	st["&quot;"] = "\""
	st["&apos;"] = "'"
	st["&amp;"] = "&"
	st["&gt;"] = ">"
	st["&lt;"] = "<"
	st["&frasl;"] = "/"
	ans := ""
	for i := 0; i < len(text); i++ {
		if text[i] != '&' {
			ans += string(text[i])
			continue
		}
		flag := false
		for j := i; j < len(text); j++ {
			if text[j] == ';' {
				if v, ok := st[text[i: j+1]]; ok {
					ans += v
					i = j
					flag = true
					break
				}
			}
		}
		if !flag {
			ans += string(text[i])
		}

	}
	return ans
}