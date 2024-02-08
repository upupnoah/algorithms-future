package heap

const N = 1e6 + 10

var n, m int
var cnt int

// 如果下标从 1 开始, 那么父节点和子节点的关系是:
// 父节点 i 的左儿子是 2*i, 右儿子是 2*i+1

// 如果下标从 0 开始, 那么父节点和子节点的关系是:
// 父节点 i 的左儿子是 2*i+1, 右儿子是 2*i+2
var h = make([]int, N)

// 递归写法
func down(u int) {
	t := u
	c := u * 2
	// 找到左右儿子中较大的那个, 这是一个大根堆
	if c <= cnt && h[c] < h[t] {
		t = c
	}
	if c+1 <= cnt && h[c+1] < h[t] {
		t = c + 1
	}
	if u != t {
		h[u], h[t] = h[t], h[u]
		down(t) // 让 t 节点继续下沉, 直到无法下沉(满足堆的性质)
	}
}

// 非递归写法
// func down(u int) {
// 	t := u
// 	for {
// 		c := u * 2
// 		if c <= cnt && h[c] < h[t] {
// 			t = c
// 		}
// 		if c+1 <= cnt && h[c+1] < h[t] {
// 			t = c + 1
// 		}
// 		if u != t {
// 			h[u], h[t] = h[t], h[u]
// 			u = t
// 		} else {
// 			break
// 		}
// 	}
// }

// 递归写法
func up(u int) {
	if u > 1 && h[u] > h[u/2] {
		h[u], h[u/2] = h[u/2], h[u]
		up(u / 2)
	}
}

// 非递归写法
// func up(u int) {
// 	for u > 1 {
// 		if h[u] > h[u/2] {
// 			h[u], h[u/2] = h[u/2], h[u]
// 			u /= 2
// 		} else {
// 			break
// 		}
// 	}
// }

func build() {
	// 叶子节点开始于 cnt/2+1, 叶子不需要 down (所有的叶子节点都是合法的堆)
	// 小根堆: 任意父节点节点 <= 子节点
	// 将所有非叶子节点 down 一遍
	for i := cnt / 2; i > 0; i-- {
		down(i)
	}

	// 如果下标从 0 开始
	// for i := cnt/2 - 1; i >= 0; i-- {
	// 	down(i)
	// }
}

// 这里下标从 1 开始
func top() int {
	return h[1]
}

func pop() {
	h[1] = h[cnt]
	cnt--
	down(1)
}
