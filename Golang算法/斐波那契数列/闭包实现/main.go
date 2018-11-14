package main

import "fmt"

const LIM = 40

func fibonacci() (func() int) {
	back1, back2 := 0, 1
	return func() int {
		// 重新赋值
		back1, back2 = back2, (back1 + back2)
		return back1
	}
}

func main() {
	f := fibonacci() //返回一个闭包函数
	var array [LIM]int
	for i := 0; i < LIM; i++ {
		array[i] = f()
	}
	fmt.Println(array)
}
