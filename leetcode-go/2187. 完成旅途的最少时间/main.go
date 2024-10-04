package main

import (
	"slices"
	"sort"
)

func minimumTime(time []int, totalTrips int) int64 {
	minT := slices.Min(time)
	return int64(sort.Search(totalTrips*minT, func(x int) bool {
		sum := 0
		for _, t := range time {
			sum += x / t
			if sum >= totalTrips {
				return true
			}
		}
		return false
	}))
}

func minimumTimeWithBinarySearch(time []int, totalTrips int) int64 {
	minT := slices.Min(time)
	left := minT - 1           // 循环不变量：sum >= totalTrips 恒为 false
	right := totalTrips * minT // 循环不变量：sum >= totalTrips 恒为 true
	for left+1 < right {       // 开区间 (left, right) 不为空
		mid := (left + right) / 2
		sum := 0
		for _, t := range time {
			sum += mid / t
		}
		if sum >= totalTrips {
			right = mid // 缩小二分区间为 (left, mid)
		} else {
			left = mid // 缩小二分区间为 (mid, right)
		}
	}
	return int64(right) // 最小的 true
}
