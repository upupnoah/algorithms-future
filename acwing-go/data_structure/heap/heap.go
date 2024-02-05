package heap

const N = 1e6 + 10

var n, m int
var cnt int
var h [N]int

func down(u int) {
	t := u
	c := u * 2
	if c <= cnt && h[c] < h[t] {
		t = c
	}
	if c+1 <= cnt && h[c+1] < h[t] {
		t = c + 1
	}
	if u != t {
		h[u], h[t] = h[t], h[u]
		down(t)
	}
}

func up(u int) {
	for u > 1 {
		if h[u] > h[u/2] {
			h[u], h[u/2] = h[u/2], h[u]
			u /= 2
		} else {
			break
		}
	}
}

func build() {
	for i := cnt / 2; i > 0; i-- {
		down(i)
	}
}

func top() int {
	return h[1]
}

func pop() {
	h[1] = h[cnt]
	cnt--
	down(1)
}
