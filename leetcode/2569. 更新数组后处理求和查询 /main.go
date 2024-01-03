package main

type SegNode struct {
	l, r, sum int
	lazytag   bool
}

type SegTree struct {
	arr []SegNode
}

func NewSegTree(nums []int) *SegTree {
	st := &SegTree{
		arr: make([]SegNode, len(nums)*4+1),
	}
	st.build(1, 0, len(nums)-1, nums)
	return st
}

func (this *SegTree) build(id, l, r int, nums []int) {
	this.arr[id].l, this.arr[id].r, this.arr[id].lazytag = l, r, false
	if l == r {
		this.arr[id].sum = nums[l]
		return
	}
	mid := (l + r) >> 1
	this.build(2*id, l, mid, nums)
	this.build(2*id+1, mid+1, r, nums)
	this.arr[id].sum = this.arr[2*id].sum + this.arr[2*id+1].sum
}

func (this *SegTree) sumRange(left, right int) int {
	return this.query(1, left, right)
}

func (this *SegTree) reverseRange(left, right int) {
	this.modify(1, left, right)
}

func (this *SegTree) pushdown(x int) {
	if this.arr[x].lazytag {
		this.arr[2*x].lazytag = !this.arr[2*x].lazytag
		this.arr[2*x].sum = this.arr[2*x].r - this.arr[2*x].l + 1 - this.arr[2*x].sum
		this.arr[2*x+1].lazytag = !this.arr[2*x+1].lazytag
		this.arr[2*x+1].sum = this.arr[2*x+1].r - this.arr[2*x+1].l + 1 - this.arr[2*x+1].sum
		this.arr[x].lazytag = false
	}
}

func (this *SegTree) modify(id, l, r int) {
	if this.arr[id].l >= l && this.arr[id].r <= r {
		this.arr[id].sum = this.arr[id].r - this.arr[id].l + 1 - this.arr[id].sum
		this.arr[id].lazytag = !this.arr[id].lazytag
		return
	}
	this.pushdown(id)
	if this.arr[2*id].r >= l {
		this.modify(2*id, l, r)
	}
	if this.arr[2*id+1].l <= r {
		this.modify(2*id+1, l, r)
	}
	this.arr[id].sum = this.arr[2*id].sum + this.arr[2*id+1].sum
}

func (this *SegTree) query(id, l, r int) int {
	if this.arr[id].l >= l && this.arr[id].r <= r {
		return this.arr[id].sum
	}
	if this.arr[id].r < l || this.arr[id].l > r {
		return 0
	}
	this.pushdown(id)
	res := 0
	if this.arr[2*id].r >= l {
		res += this.query(2*id, l, r)
	}
	if this.arr[2*id+1].l <= r {
		res += this.query(2*id+1, l, r)
	}
	return res
}

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	n, m := len(nums1), len(queries)
	tree := NewSegTree(nums1)
	sum := int64(0)
	for _, x := range nums2 {
		sum = sum + int64(x)
	}
	res := []int64{}
	for i := 0; i < m; i++ {
		if queries[i][0] == 1 {
			l, r := queries[i][1], queries[i][2]
			tree.reverseRange(l, r)
		} else if queries[i][0] == 2 {
			sum += int64(tree.sumRange(0, n-1)) * int64(queries[i][1])
		} else if queries[i][0] == 3 {
			res = append(res, sum)
		}
	}
	return res
}
