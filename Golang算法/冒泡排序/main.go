package main

import (
	"fmt"
)

func bubbleSort(items []int) {
	var (
		n       = len(items)
		swapped = true
	)
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		n = n - 1
	}
}

func main() {
	items := []int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}
	fmt.Println(items)
	bubbleSort(items)
	fmt.Println(items)
}
