package main

import (
	"fmt"
	"time"
)

const (
	//TimeFormart = "2006-01-02 15:04:05"
	TimeFormart = "20060102"
)

func main() {
	NowTime := time.Now()
	fmt.Printf("getnowtime is %s\n", NowTime)

	diff, err := time.ParseDuration("-24h")
	if err != nil {
		fmt.Println(err)
	}

	yesterday := NowTime.Add(diff).Format(TimeFormart)
	fmt.Printf("yesterday is %s\n", yesterday)
	now := time.Now().Format("20060102")
	fmt.Printf("getnowtime in format Ymd is %s\n", now)
	now1 := time.Now().Format("20060102 15:04:05")
	fmt.Printf("getnowtime in format Ymd HMS is %s\n", now1)
	now2 := time.Now().Format("20060102 15:04:05.000 +0800")
	fmt.Printf("getnowtime in format Ymd HMS.ms mis nans TZ is %s\n", now2)

	now3 := time.Now()
	yesterday1 := now3.AddDate(0, 0, -1)
	bef_yes := yesterday1.AddDate(0, 0, -1)
	fmt.Printf("Yesterday:%s\n", yesterday1.Format("20060102"))
	fmt.Printf("Yesterdat before Yesterday:%s\n", bef_yes.Format("20060102"))

	t := time.Now()
	//year, month, day := t.Date()
	y, m, _ := t.Date()
	thisMonthFirstDay := time.Date(y, m, 1, 1, 1, 1, 1, t.Location())
	bef_yesday := thisMonthFirstDay.AddDate(0, -1, 0)
	fmt.Printf("Today:%s\n", t)
	fmt.Printf("Yesterday:%d\n", y)
	//fmt.Printf("Yesterdat before Yesterday:%s\n", bef_yesday)
	fmt.Printf("Yesterdat before Yesterday:%s\n", bef_yesday.Format("01"))

}
