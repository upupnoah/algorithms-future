package main

type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack1) == 0 {
		return -1
	}
	for len(this.stack1) > 0 {
		this.stack2 = append(this.stack2, this.stack1[len(this.stack1) -1])
		this.stack1 = this.stack1[:len(this.stack1) - 1]
	}
	res := this.stack2[len(this.stack2) - 1]
	this.stack2 = this.stack2[:len(this.stack2) - 1]
	for len(this.stack2) > 0 {
		this.stack1 = append(this.stack1, this.stack2[len(this.stack2)-1])
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
	return res
}