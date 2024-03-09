package main

import (
	"container/heap"
	"slices"
	"sort"
)

// 问题转化:
// 计算 nums 中非负数的和
// nums 中的任意一个子序列的元素和, 都等价于从 sum 中减去某些非负数/加上某些负数得到
// -> 减去非负数 和 加上负数 -> 等价于 减去 |nums[i]|
// sum 减去的数越小, nums 的子序列的和就越大
// 现在的问题:
// 把每个 nums[i] 取绝对值后, nums 的第 k 小的子序列和是多少?
// 最后的答案就是 sum - kth5

// 二分答案 + dfs
// 二分答案: sumLimit, 判断是否有至少 k 个子序列, 其元素和不超过 sumLimit (因为要前 k 大, 所有至少得有 k 个序列)
func kSum1(nums []int, k int) int64 {
	// sum: 所有原本的非负数元素和
	// total: 所有元素转化为正数之后的和
	sum, total := 0, 0
	for i, x := range nums {
		if x >= 0 {
			sum += x
			total += x
		} else {
			total -= x
			nums[i] = -x
		}
	}
	slices.Sort(nums)
	kthS := sort.Search(total, func(sumLimit int) bool {
		cnt := 1 // 空子序列算一个

		// 子集型回溯
		var dfs func(int, int)
		dfs = func(i, s int) {
			// 已经找到, 越界, 超过 sumLimit 了
			if cnt == k || i == len(nums) || s+nums[i] > sumLimit {
				return
			}
			cnt++
			dfs(i+1, s+nums[i]) // 选
			dfs(i+1, s)         // 不选
		}
		dfs(0, 0)
		return cnt == k // 找到 k 个元素和不超过 sumLimit 的子序列
	})
	return int64(sum - kthS)
}

// 方法二: 小根堆
func kSum(nums []int, k int) int64 {
	n := len(nums)
	sum := 0
	for i, x := range nums {
		if x >= 0 {
			sum += x
		} else {
			nums[i] = -x
		}
	}
	slices.Sort(nums)

	h := hp{{0, 0}} // 空子序列
	// 再找 k-1 个子序列(空子序列算一个)
	for ; k > 1; k-- {
		p := heap.Pop(&h).(pair)
		i := p.i
		if i < n {
			// 在子序列的末尾添加 nums[i]
			heap.Push(&h, pair{p.sum + nums[i], i + 1}) // 下一个添加/替换的元素下标为 i+1
			if i > 0 {
				// 替换子序列的末尾元素为 nums[i]
				heap.Push(&h, pair{p.sum + nums[i] - nums[i-1], i + 1})
			}
		}
	}
	return int64(sum - h[0].sum)
}

type pair struct{ sum, i int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
