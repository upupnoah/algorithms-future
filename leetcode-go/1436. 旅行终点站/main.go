package main

func destCity(paths [][]string) string {
	setA := make(map[string]struct{}, len(paths)) // 预分配空间
	for _, p := range paths {
		setA[p[0]] = struct{}{}
	}

	for _, p := range paths {
		if _, ok := setA[p[1]]; !ok {
			return p[1]
		}
	}
	return ""
}
