package _094__Car_Pooling

import "sort"

// 算法：差分
// 思路：因为本题的数据只有 1000，且下标都是 >= 0 的，因此可以创建一个数组进行差分
// 为了练习 哈希 + 排序 -> 缩小数据范围
func carPooling(trips [][]int, capacity int) bool {
	diff := map[int]int{}
	for _, t := range trips {
		diff[t[1]] += t[0]
		diff[t[2]+1] -= t[0]
	}
	n := len(diff)
	passengers := make([]int, 0, n)
	for k := range diff {
		passengers = append(passengers, k)
	}
	sort.Ints(passengers)
	sum := 0
	for i := range passengers {
		sum += diff[passengers[i]]
		if sum > capacity {
			return false
		}
	}
	return true
}

// 简化版本
func carPooling2(trips [][]int, capacity int) bool {
	diff := make([]int, 1001)
	for _, t := range trips {
		diff[t[1]] += t[0]
		diff[t[2]] -= t[0]
	}
	sum := 0
	for _, d := range diff {
		sum += d
		if sum > capacity {
			return false
		}
	}
	return true
}
