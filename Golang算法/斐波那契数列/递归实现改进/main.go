package main

import "fmt"

const LIM = 40

var fibs [LIM]uint64

func fibonacci(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}

func main() {
	//var result uint64 = 0
	var array [LIM]uint64
	for i := 0; i < LIM; i++ {
		array[i] = fibonacci(i)
		//result = fibonacci(i)
		//array = append(array, result)
		//fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	fmt.Println(array)
}
