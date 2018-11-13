package main

import (
	"math/rand"
	"time"
	"fmt"
	"strconv"
)

func main() {
	var (
		inpstr string // 获取用户输入
		intput int    // 用户输入的值
		result int    // 答案
		err    error
	)

	//1.获取 100 以内随机整数作为 答案
	rand.Seed(time.Now().Unix())
	result = rand.Intn(100)

	for {
		//2.等待用户输入，非法就继续让用户输入
		if _, err = fmt.Scan(&inpstr); err != nil {
			fmt.Println("输入非法")
			continue
		}
		if intput, err = strconv.Atoi(inpstr); err != nil {
			fmt.Println("输入非法")
			continue
		}

		//3.用户输入的值 和 答案 比较，提示是大还是小?。
		switch {
		case intput > result:
			fmt.Println("数字大了")
		case intput < result:
			fmt.Println("数字小了")
		case intput == result:
			fmt.Println("你赢了")
			goto LOOP
		}
	}
	LOOP:
}
