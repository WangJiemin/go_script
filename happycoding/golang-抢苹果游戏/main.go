package main

import (
	"fmt"
)

//游戏规则，a和b 一起数0～9 数字，谁先数到9，谁赢

func play(p chan<- int, index int) {
	for i := 0; i < index; i++ {
		p <- i
	}
	close(p)
}

//a 和 b 抢苹果
//游戏规则，a和b 一起数0～9 数字，谁先数到9，谁赢

func main() {
	a := make(chan int)
	b := make(chan int)

	//a和b加入游戏
	go play(a, 10)
	go play(b, 10)

	s := make(chan string)

	//裁判

	go func(a, b <-chan int, s chan<- string) {
		for {
			select {
			case v, ok := <-a:
				if !ok {
					s <- "a  赢了"
				}
				fmt.Printf("a %d \n", v)
			case v, ok := <-b:
				if !ok {
					s <- "b  赢了"
				}
				fmt.Printf("b %d \n", v)
			}
		}
	}(a, b, s)

	//观众等待结果
	fmt.Printf("%s \n", <-s)

}


//func main() {
//	a := make(chan int)
//	b := make(chan int)
//	s := make(chan string)
//
//	go play(a, 10)
//	go play(b, 10)
//
//	//裁判
//	go func(a <-chan int, b <-chan int, s chan<- string) {
//		for {
//			select {
//			case v, ok := <-a:
//				if !ok {
//					s <- "a"
//				}
//				fmt.Printf("a:%d \n", v)
//			case v, ok := <-b:
//				if !ok {
//					s <- "b"
//				}
//				fmt.Printf("b:%d \n", v)
//			}
//
//		}
//	}(a, b, s)
//
//	fmt.Printf("game over %s \n", <-s)
//}
//
////游戏规则
//func play(p chan<- int, index int) {
//	for i := 0; i < index; i++ {
//		p <- i
//	}
//	close(p)
//}
