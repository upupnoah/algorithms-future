package _21__数组中两个数的最大异或值

import "math"

type trie01Node struct {
	son [2]*trie01Node
	cnt int // 子树叶子数
	min int // 子树最小值
}

type trie01 struct{ root *trie01Node }

func newTrie01() *trie01 { return &trie01{&trie01Node{min: math.MaxInt32}} }

const trieBitLen = 31 // 30 for 1e9, 63 for int64, or bits.Len(MAX_VAL)

func (*trie01) bin(v int) []byte {
	s := make([]byte, trieBitLen)
	for i := range s {
		s[i] = byte(v >> (trieBitLen - 1 - i) & 1)
	}
	return s
}

func (t *trie01) put(v int) *trie01Node {
	o := t.root
	if v < o.min {
		o.min = v
	}
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			o.son[b] = &trie01Node{min: math.MaxInt32}
		}
		o = o.son[b]
		o.cnt++
		if v < o.min {
			o.min = v
		}
	}
	//o.val = v
	return o
}

// v 与 trie 上所有数的最大异或值，trie 不能是空的
func (t *trie01) maxXor(v int) (ans int) {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b^1] != nil {
			ans |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return
}
func findMaximumXOR(nums []int) (ans int) {
	trie := newTrie01()
	for _, x := range nums {
		trie.put(x)
		ans = max(ans, trie.maxXor(x)) // a^b = b^a, 因此不会错过最大值
	}
	return
}
