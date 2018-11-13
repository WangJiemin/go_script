package main

import (
	SL "golang_Learn/ptslowlog/slowlog"
	"fmt"
)

func main(){
	//SL.SlowLogForMart("/home/jiemin/code/go_dev/src/golang_Learn/ptslowlog/conf/conf.yaml")
	slowlogStr:=SL.SlowLogForMart("/home/jiemin/code/go_dev/src/golang_Learn/ptslowlog/conf/conf.yaml")
	fmt.Println(slowlogStr)
}

