package main

type MaxQueue struct {
	in []int
	out []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.out) == 0 {
		return -1
	}
	return this.out[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.in = append(this.in, value)
	for len(this.out) > 0 && this.out[len(this.out)-1] < value {
		this.out = this.out[:len(this.out)-1]
	}
	this.out = append(this.out, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.in) == 0 {
		return -1
	}
	res := this.in[0]
	this.in = this.in[1:]
	if res == this.out[0] {
		this.out = this.out[1:]
	}
	return res
}