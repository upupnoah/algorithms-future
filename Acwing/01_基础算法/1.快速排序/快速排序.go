package main

import "fmt"

func quickSort(q []int, l int, r int) {
	if l >= r { // Check the boundary condition, return if there's only one element or no elements in the range
		return
	}
	x := q[(l+r)>>1] // Pivot
	i, j := l-1, r+1
	for i < j {
		i++ // Move i to the right until finding an element >= x
		for q[i] < x {
			i++
		}
		j-- // Move j to the left until finding an element <= x
		for q[j] > x {
			j--
		}
		if i < j { // Swap q[i] and q[j] if i is still to the left of j
			q[i], q[j] = q[j], q[i]
		}
	}
	quickSort(q, l, j)   // Recursively sort the left segment
	quickSort(q, j+1, r) // Recursively sort the right segment
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	quickSort(q, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}
	return
}
