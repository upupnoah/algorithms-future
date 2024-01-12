package _251__Number_of_Flowers_in_Full_Bloom

import "sort"

// 暴力
//func fullBloomFlowers(flowers [][]int, people []int) []int {
//	var res []int
//	for _, p := range people {
//		cnt := 0
//		for _, f := range flowers {
//			if f[0] <= p && p <= f[1] {
//				cnt++
//			}
//		}
//		res = append(res, cnt)
//	}
//	return res
//}

// 做法 1：转换 + 二分
// 把 flowers 分成两个数组 starts 和 ends， 对他们进行排序
// 第 i 个人能看到的花的数目等于 (start 中 <= people[i]的数目) - (ends 中 < people[i]的数目)
// 如何快速计算 开花数 和 凋落数？
// 		二分！
//  - 把 starts 和 ends 从小到大排序

//   starts 中 <= people[i]的表示开花，这里可以直接找下标（下标从 0 开始，因此 > people[i] 的下一个数的下标为开花数
//   ends 中 < people[i]的表示凋落的（例如 在 2 时刻来看花，但是在 1 时刻就凋谢了）
//	- 在 starts 中二分查找 > people[i] 的下一个数的下标（不存在则为 n），即为开花数
//  - 在 ends   中二分查找 >= people[i] 的下一个数的下标（不存在则为 n），即为凋落数

func fullBloomFlowers(flowers [][]int, people []int) []int {
	n := len(flowers)
	starts := make([]int, n)
	ends := make([]int, n)
	for i, f := range flowers {
		starts[i] = f[0]
		ends[i] = f[1]
	}
	sort.Ints(starts)
	sort.Ints(ends)
	// 二分中，因为调用的是 sorts.SearchInts(), 默认是 return a[i] >= x
	// 因此这里要在 starts 中找 > x 的下一个数的下标， 可以转换成 找 >= x+1 的数的第一个下标
	for i, p := range people {
		people[i] = sort.SearchInts(starts, p+1) - sort.SearchInts(ends, p)
	}
	return people
}

// 方法二：差分
// 把flowers[i]看成是将区间[start，end]上的每个时间点都增加一朵花。
// 例如示例1，我们可以从一个全为0的数组开始，把下标在区间[1，6]，[3，7]， 12]，[4，13]内的数都加一
// 如果一个下标出现在多个区间内，就多次加一。
// 加完后，我们得到的数组就是a= [0，1，1，1，2，3, 3，2，1，2，1，2，2，2，2，1]
// 那么answer[i]就等于 person[i]时间点上有多少朵花，即a[person[i]]。

func fullBloomFlowers2(flowers [][]int, people []int) []int {
	// 初始化差分数组，在 d[i]的位置+x， 在 d[j+1]的位置-x
	// 最后通过sum 的方式可以还原 a 数组
	diff := map[int]int{}
	for _, f := range flowers {
		diff[f[0]]++
		diff[f[1]+1]--
	}
	// 这题元素值很大，需要吧数组改成有序集合，或者用哈希表 + 排序
	// 这里我们用哈希表 + 排序
	n := len(diff)
	times := make([]int, 0, n)
	for t := range diff {
		times = append(times, t)
	}
	sort.Ints(times)

	// 将 id 通过 people 排序
	id := make([]int, len(people))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return people[id[i]] < people[id[j]] })

	j, sum := 0, 0
	// times 就是 diff 中的 key，但是已经排序了，因此可以通过 time 来找到 diff 中的 value
	for _, i := range id {
		for ; j < n && times[j] <= people[i]; j++ {
			sum += diff[times[j]] // 累加不超过 people[i] 的差分值
		}
		people[i] = sum // 从而得到这个时刻花的数量
	}
	return people
}

func fullBloomFlowers3(flowers [][]int, people []int) []int {
	// 初始化差分数据
	diff := map[int]int{}
	for _, f := range flowers {
		diff[f[0]]++
		diff[f[1]+1]--
	}

	// 将差分数据中的 key 弄出来排序
	// 因为 map 是无序的，因此需要将 key 整出来排序
	n := len(diff)
	times := make([]int, 0, n)
	for k := range diff {
		times = append(times, k)
	}

	sort.Ints(times)

	// 为了最后一次遍历（累加） people，可以将 people 的下标排序
	id := make([]int, len(people))
	for i := range id {
		id[i] = i
	}
	// 映射的是排好序之后的 people 的下标
	sort.Slice(id, func(i, j int) bool {
		return people[i] < people[j]
	})
	// Note： 遍历id 数组， people[id[i]]  可以得到一个排序之后的 people 数组
	j, sum := 0, 0
	for _, i := range id {
		// 因为 id 数组是排好序之后的 people 的下标，因此从头开始遍历 id 的值，那么对应的 people[i]
		// 也是从小到大的，这样可以一次 sum，求完答案
		// 累加不超过people[i] 的所有的值（差分数组还原成原数组的过程）
		for ; j < n && times[j] <= people[i]; j++ {
			sum += diff[times[j]]
		}
		people[i] = sum
	}
	return people
}
