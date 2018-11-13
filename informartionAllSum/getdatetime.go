package informartionAllSum

import (
	"fmt"
	"time"
)

const (
	//TimeFormart = "2006-01-02 15:04:05"
	TimeFormart = "20060102"
)

func GetYesterDay() (yd string) {
	t1 := time.Now()
	diff, err := time.ParseDuration("-24h")
	if err != nil {
		fmt.Println(err)
	}

	yd = t1.Add(diff).Format(TimeFormart)
	return yd

}
