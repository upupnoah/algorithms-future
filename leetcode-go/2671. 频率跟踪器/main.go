package main

type FrequencyTracker struct {
	cnt       map[int]int
	frequency map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		map[int]int{},
		map[int]int{},
	}
}

func (this *FrequencyTracker) Update(number, delta int) {
	this.frequency[this.cnt[number]]--
	this.cnt[number] += delta
	this.frequency[this.cnt[number]]++
}

func (this *FrequencyTracker) Add(number int) {
	this.Update(number, 1)
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if this.cnt[number] > 0 {
		this.Update(number, -1)
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.frequency[frequency] > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
