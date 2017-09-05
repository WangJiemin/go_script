package gettime

import (
	"log"
	"time"
)

const (
	TimeFormart = "2006-01-02 15:04:05"
)

func GetNowTime() (starttime string, stoptime string) {
	now := time.Now()
	yesterday, err := time.ParseDuration("-24h")
	if err != nil {
		log.Println(err)
	}
	start := now.Add(yesterday).Format(TimeFormart)
	current := now.Format(TimeFormart)
	return start, current
}
