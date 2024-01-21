package maximum_swap

import "container/heap"

func maximumSwap(num int) int {
	nums := []int{}
	hp := hp64{}
	for num > 0 {
		hp.push(num % 10)
		nums = append(nums, num%10)
		num /= 10
	}
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		t := hp.pop()
		if nums[i] == t {
			continue
		}
		hp.push(t)
		last := 0
		for j := i - 1; j >= 0; j-- {
			if nums[j] == t {
				last = j
			}
		}
		nums[i], nums[last] = nums[last], nums[i]
		break
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans = ans*10 + nums[i]
	}
	return ans
}

// 自定义类型（int64 可以替换成其余类型）
type hp64 []int

func (h hp64) Len() int           { return len(h) }
func (h hp64) Less(i, j int) bool { return h[i] > h[j] } // > 为最大堆
func (h hp64) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp64) Push(v any)        { *h = append(*h, v.(int)) }
func (h *hp64) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp64) push(v int)        { heap.Push(h, v) }
func (h *hp64) pop() int          { return heap.Pop(h).(int) } // 稍微封装一下，方便使用
