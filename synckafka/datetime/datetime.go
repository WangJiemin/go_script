package datetime

import (
	"log"
	"time"
)

const (
	//TimeFormart  = "2006-01-02 15:04:05"
	TimeFormart = "20060102"
)

func GetNowTime() (yesterday string) {
	now := time.Now()
	diff, err := time.ParseDuration("-24h")
	if err != nil {
		log.Println(err)
	}
	yesterday = now.Add(diff).Format(TimeFormart)
	return yesterday
}
