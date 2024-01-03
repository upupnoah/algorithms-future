package main

func vowelStrings(words []string, left int, right int) (ans int) {
	vowel := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	for _, v := range words[left : right+1] {
		if vowel[v[0]] && vowel[v[len(v)-1]] {
			ans++
		}
	}
	return
}
