package _726__同积元组

// 每一对元组都可以 转换成 8 种选法
// [a*b] = [c*d]
// [c*d] = [a*b]
// ...
// 思路：
//  1. 使用 hash 表统计两两之间乘积出现的次数
//  2. 对于每一个乘积，从中任选 2 个的不同的数对
//  	= C_{count(a*b)}^{2} = (count(a*b)*(count(a*b)-1))/2
// 最后/2 （因为选择的顺序不重要，所以我们需要除以 2 来消除重复计数（即，(数对1, 数对2) 和 (数对2, 数对1) 被视为同一个组合））
//  3. 每一种 [a*b] = [c*d] （每个数对组合可以生成 8 个不同的同积元组）
//  4. 最后的公式为：(count(a*b)*(count(a*b)-1))/2 * 8
//                = count(a*b)*(count(a*b)-1) * 4

func tupleSameProduct(nums []int) int {
	n := len(nums)
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cnt[nums[i]*nums[j]]++
		}
	}
	ans := 0
	for _, v := range cnt {
		ans += v * (v - 1) * 4
	}
	return ans
}
