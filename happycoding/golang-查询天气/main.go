package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Resp struct {
	Data   Weather `json:"data"`
	Status int     `json:"status"`
	Desc   string  `json:"desc"`
}

type Weather struct {
	Yesterday WeatherInfo   `json:"yesterday"`
	Forecast  []WeatherInfo `json:"forecast"`
	Wendu     string        `json:"wendu"`
	Qanmao    string        `json:"ganmao"`
	City      string        `json:"city"`
	Aqi       string        `json:"aqi"`
}

type WeatherInfo struct {
	Date      string `json:"date"`
	High      string `json:"high"`
	Fengli    string `json:"fengli"`
	Low       string `json:"low"`
	Fengxiang string `json:"fengxiang"`
	Type      string `json:"type"`
	fx        string `json:"fx"`
	fl        string `json:"fl"`
}

// http://wthrcdn.etouch.cn/weather_mini?city=
func main() {
	if len(os.Args) <= 1 {
		log.Fatal("请输入城市")
	}
	city := os.Args[1]

	resp, _ := http.Get("http://wthrcdn.etouch.cn/weather_mini?city=" + city)
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)
	var data Resp
	json.Unmarshal(bs, &data)

	fmt.Printf("城市：%s \n", data.Data.City)
	fmt.Printf("天气：%s \n", data.Data.Forecast[0].Type)
	fmt.Printf("温度：%s %s \n", data.Data.Forecast[0].Low, data.Data.Forecast[0].High)
}
