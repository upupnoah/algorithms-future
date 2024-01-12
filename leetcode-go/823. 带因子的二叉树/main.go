package _23__带因子的二叉树

import "sort"

// 思路
// 构造出的二叉树一定是完全二叉树（要么是只有两个儿子，要么一个儿子都没有）
// p = xy  p > x && p > y （因为 x 和 y 不等于 1，且是 p 的因子）
// 1. 排序
// 2. f[i] 表示以 arr[i] 为根的二叉树的个数
// 3. f[i] = sum(f[j] * f[k]) (j < i && k < i && arr[j] * arr[k] == arr[i])
func numFactoredBinaryTrees(arr []int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	n := len(arr)
	hash := make(map[int]int) // arr[i] -> i
	for i := 0; i < n; i++ {
		hash[arr[i]] = i
	}
	res, mod := 0, int(1e9+7)
	f := make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = 1 // 只有一个节点的二叉树是一个方案
		for j := 0; j < i; j++ {
			if arr[i]%arr[j] == 0 { // 可以整除，那么 arr[j] 是 arr[i] 的因子
				d := arr[i] / arr[j] // 求出 arr[k]
				if k, ok := hash[d]; ok {
					f[i] = (f[i] + f[j]*f[k]) % mod
				}
			}
		}
		res = (res + f[i]) % mod
	}
	return res
}
