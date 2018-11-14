package main

import (
	"fmt"
)

const LIM = 40

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return res
}

func main() {
	//result := 0
	//var array []int
	var array [LIM]int
	for i := 0; i < LIM; i++ {
		array[i] = fibonacci(i)
		//result = fibonacci(i)
		//array = append(array, result)
		//fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	fmt.Println(array)
}
