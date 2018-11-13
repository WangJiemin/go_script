package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go(i) // goroutine运行Go函数
	}
	<-ch
}

func Go(index int) { // 定一个Go函数，在里面传递进去一个参数，是int类型的
	num := 1                         // 定义一个参数，把1赋值给num
	for i := 0; i < 100000000; i++ { // for循环,循环叠加一亿次
		num += i // num加等i
	}
	fmt.Println(i, num) // 打印i的值，然后在输出num的值
}
