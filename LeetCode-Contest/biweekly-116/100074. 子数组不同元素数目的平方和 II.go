package biweekly_116

func sumCounts(nums []int) int {
	ans := 0
	n := len(nums)
	root := &segTree{0, n - 1, 0, nil, nil, 0}
	occur := make(map[int]int)
	sum := 0
	for i := range nums {
		l := -1
		if v, ok := occur[nums[i]]; ok {
			l = v
		}
		q := root.query(l+1, i)
		sum = (sum + 2*q + (i - l)) % MOD
		ans = (ans + sum) % MOD
		root.insert(l+1, i, 1)
		occur[nums[i]] = i
	}
	return ans
}

const MOD = int(1e9 + 7)

type segTree struct {
	left, right, sum int
	lChild, rChild   *segTree
	lazyAdd          int
}

func (this *segTree) calcLazySum() int {
	return this.lazyAdd * (this.right - this.left + 1) % MOD
}
func (this *segTree) calcSum() int {
	return (this.sum + this.calcLazySum()) % MOD
}
func (this *segTree) markDown() {
	if this.left == this.right {
		return
	}
	mid := (this.left + this.right) / 2
	if this.lChild == nil {
		this.lChild = &segTree{left: this.left, right: mid}
	}
	this.lChild.lazyAdd = (this.lChild.lazyAdd + this.lazyAdd) % MOD
	if this.rChild == nil {
		this.rChild = &segTree{left: mid + 1, right: this.right}
	}
	this.rChild.lazyAdd = (this.rChild.lazyAdd + this.lazyAdd) % MOD
	this.sum = this.calcSum()
	this.lazyAdd = 0
}
func (this *segTree) insert(l, r, val int) {
	if l <= this.left && r >= this.right {
		this.lazyAdd = (this.lazyAdd + val) % MOD
		return
	}
	this.markDown()
	mid := (this.left + this.right) / 2
	if l <= mid {
		this.lChild.insert(l, r, val)
	}
	if r > mid {
		this.rChild.insert(l, r, val)
	}
	this.sum = (this.lChild.calcSum() + this.rChild.calcSum()) % MOD
}
func (this *segTree) query(l, r int) int {
	if l <= this.left && r >= this.right {
		return this.calcSum()
	}
	this.markDown()
	mid := (this.left + this.right) / 2
	ans := 0
	if l <= mid {
		ans = (ans + this.lChild.query(l, r)) % MOD
	}
	if r > mid {
		ans = (ans + this.rChild.query(l, r)) % MOD
	}
	return ans
}
