package main

import (
	"fmt"
)

const LIM = 40

func fibonacci(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else {
		a, b := 1, 1
		result := 0
		for i := 3; i <= n; i++ {
			result = a + b
			a, b = b, result
		}
		return result
	}
}

func main() {
	fmt.Println(fibonacci(LIM))
}
