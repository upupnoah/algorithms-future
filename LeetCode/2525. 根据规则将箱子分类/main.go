package _525__根据规则将箱子分类

func categorizeBox(length int, width int, height int, mass int) string {
	x := length >= 1e4 || width >= 1e4 || height >= 1e4 || mass >= 1e4 || width*height >= 1e9/length || length*height >= 1e9/width
	y := mass >= 100
	switch {
	case x && y:
		return "Both"
	case x:
		return "Bulky"
	case y:
		return "Heavy"
	default:
		return "Neither"
	}
}
