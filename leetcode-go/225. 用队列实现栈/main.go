package main

type MyStack struct {
	q1 []int
	q2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q1 = append(this.q1, x)
}

func (this *MyStack) Pop() int {
	for len(this.q1) > 1 {
		this.q2 = append(this.q2, this.q1[0])
		this.q1 = this.q1[1:len(this.q1)]
	}
	v := this.q1[0]
	this.q1 = this.q2
	this.q2 = []int{}
	return v
}

func (this *MyStack) Top() int {
	for len(this.q1) > 1 {
		this.q2 = append(this.q2, this.q1[0])
		this.q1 = this.q1[1:len(this.q1)]
	}
	v := this.q1[0]
	this.q2 = append(this.q2, this.q1[0])
	this.q1 = this.q2
	this.q2 = []int{}
	return v
}

func (this *MyStack) Empty() bool {
	return len(this.q1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
