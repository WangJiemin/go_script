package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 实现一个猜数字额游戏
func main() {
	var (
		inpstr string //获取用户输入
		intput int    //用户输入的值
		result int    //答案
		err    error
	)
	// 1.获取 100以内 随机书作为，答案
	rand.Seed(time.Now().Unix())
	result = rand.Intn(100)
	// 赢了就退出程序
Out:
	for {
		// 2.等待用户输入 非法就继续让用户输入
		if _, err = fmt.Scan(&inpstr); err != nil {
			fmt.Println("输入非法")
			continue
		}
		if intput, err = strconv.Atoi(inpstr); err != nil {
			fmt.Println("输入非法")
			continue
		}
		// 3.用户输入的值 和 答案 比较，提示是大还是小？。
		switch {
		case intput > result:
			fmt.Println("数字大了")
		case intput < result:
			fmt.Println("数字小了")
		case intput == result:
			fmt.Println("你赢了")
			break Out
		}
	}

}
