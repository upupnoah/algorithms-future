package main

// 脑瘫了，直接写了一个高精度加法
func addStrings(num1 string, num2 string) string {
	n, m := len(num1), len(num2)
	if n < m {
		return addStrings(num2, num1)
	}
	num1, num2 = reverse(num1), reverse(num2)
	return add(num1, num2)
}

func add(a, b string) string {
	t := 0
	var c []byte
	for i := 0; i < len(a) || i < len(b); i++ {
		if i < len(a) {
			t += int(a[i] - '0')
		}
		if i < len(b) {
			t += int(b[i] - '0')
		}
		c = append(c, byte(t%10+'0'))
		t /= 10
	}
	if t != 0 {
		c = append(c, byte('1'))
	}
	return reverse(string(c))
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
