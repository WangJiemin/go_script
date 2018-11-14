package main

import "fmt"

func fibonacci(num int) int {

	if num < 2 {
		return 1
	}

	return fibonacci(num-1) + fibonacci(num-2)

}

func main() {

	for i := 0; i < 10; i++ {
		nums := fibonacci(i)
		fmt.Println(nums)
	}
}
