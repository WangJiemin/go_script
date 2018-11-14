package main

import (
	"fmt"
)

func merge(left, right []int) (result []int) {
	lf, lg := len(left), len(right)
	result = make([]int, lf+lg)

	i := 0
	for lf > 0 && lg > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	// Either left or right may have elements left; consume them. 2
	// (Only one of the following loops will actually be entered.)
	for j := 0; j < lf; j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < lg; j++ {
		result[i] = right[j]
		i++
	}
	return
}

func mergeSort(items []int) []int {
	var n = len(items)

	if n == 1 {
		return items
	}

	middle := int(n / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, n-middle)
	)
	for i := 0; i < n; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func main() {
	items := []int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}
	fmt.Println(items)
	sortedItems := mergeSort(items)
	fmt.Println(sortedItems)
}
