package _333__Filter_Restaurants_by_Vegan_Friendly__Price_and_Distance

import "sort"

// 题意：按照 3 个过滤器（vegan素食，maxPrice，maxDistance）来筛选符合条件的餐厅
//
//	如果筛选出来的几个餐厅 rating 一样，那么按照 id 从高到底排序
//	veganFriendly = 0/1 (0表示：any，1 表示： vegan餐厅）
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	// 1. 实现一个排序函数
	// 对于 vegan 为 1：vegan 为 0 的都小
	// 对于 vegan 为 0：那就考虑 rating，如果 rating 一样，则按照 id 从高到低排序
	// ** 排序的时候似乎不需要考虑 vegan 是否为 1，因为最后遍历的时候还是需要判断 vegan 的值

	// 2. 遍历结果，判断是否符合 maxPrice 和 maxDistance
	//	将符合条件的加入答案中

	// sort
	sort.Slice(restaurants, func(i int, j int) bool {
		a, b := restaurants[i], restaurants[j]
		if a[1] == b[1] {
			return a[0] > b[0]
		}
		return a[1] > b[1]
	})

	var res []int
	for _, r := range restaurants {
		if veganFriendly == 1 && r[2] == 0 {
			continue
		}
		if r[3] <= maxPrice && r[4] <= maxDistance {
			res = append(res, r[0])
		}
	}
	return res
}
