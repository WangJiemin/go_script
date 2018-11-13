package main

import (
	"fmt"
	"math"
)

func StringToIntArray(input string) []int {
	output := []int{}
	for _, v := range input {
		output = append(output, int(v))
	}
	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}
	return output
}

func getInput(input string) (<-chan int) {
	out := make(chan int)
	go func() {
		for _, b := range StringToIntArray(input) {
			out <- b
		}
		close(out)
	}()

	return out
}
func sq(in <-chan int) (<-chan int) {
	out := make(chan int)

	var base, i float64 = 2, 0
	go func() {
		for n := range in {
			out <- (n - 48) * int(math.Pow(base, i))
			i++
		}
		close(out)
	}()
	return out
}

func main() {
	input := "1010"
	c := getInput(input)
	out := sq(c)
	sum := 0
	for o := range out {
		// fmt.Println(o)
		sum += o
	}
	fmt.Println(sum)
	// fmt.Println(sum / 1024 / 1024)
}
