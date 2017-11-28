package gettime

import (
	"log"
	"time"
)

const (
	TimeFormart  = "2006-01-02 15:04:05"
	TimeFarmart1 = "20060102"
)

func GetNowTime() (starttime string, stoptime string) {
	now := time.Now()
	yesterday, err := time.ParseDuration("-48h")
	if err != nil {
		log.Println(err)
	}
	start := now.Add(yesterday).Format(TimeFormart)
	current := now.Format(TimeFormart)
	return start, current
}

func YESTERDAY() (yesday string) {
	NowTime := time.Now()
	//fmt.Println(NowTime)

	diff, err := time.ParseDuration("-24h")
	if err != nil {
		log.Println(err)
	}

	yesday = NowTime.Add(diff).Format(TimeFarmart1)
	//fmt.Println(yesterday)
	return yesday
}

func TWODAYS() (one, two string) {
	NowTime := time.Now()
	//fmt.Println(NowTime)

	diff_one, err := time.ParseDuration("-48h")
	if err != nil {
		log.Println(err)
	}

	diff_two, err := time.ParseDuration("-72h")
	if err != nil {
		log.Println(err)
	}

	one = NowTime.Add(diff_one).Format(TimeFarmart1)
	two = NowTime.Add(diff_two).Format(TimeFarmart1)
	//fmt.Println(yesterday)
	return one, two
}
