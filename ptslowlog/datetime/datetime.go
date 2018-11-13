package datetime

import (
	"time"
	"log"
)

const (
	TimeFormart="2006-01-02 15:04:05"
	TimeFormartNumYMD="20060102"
)

func DateTime() (now, ysday string) {
	tnow:=time.Now()
	diff,err:=time.ParseDuration("-24h")
	if err !=nil{
		log.Println(err)
	}
	tday:=tnow.Format(TimeFormartNumYMD)
	yestime:=tnow.Add(diff).Format(TimeFormartNumYMD)
	return tday,yestime
}
