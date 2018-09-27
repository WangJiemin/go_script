package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type INFORMATRIONALL struct {
	ID               int    `gorm:"column:id"`
	Url              string `gorm:"column:url"`
	Title            string `gorm:"column:title"`
	Author           string `gorm:"column:author"`
	Source           string `gorm:"column:source"`
	Release_datetime string `gorm:"column:release_datetime"`
	Content          string `gorm:"column:content"`
	Media_type       string `gorm:"column:media_type"`
	Original_title   string `gorm:"column:original_title"`
	Editor           string `gorm:"column:editor"`
	Reporter         string `gorm:"column:reporter"`
	Contents         string `gorm:"column:contents"`
	Reading_times    string `gorm:"column:reading_times"`
	Abstract_data    string `gorm:"column:abstract_data"`
	Media            string `gorm:"column:media"`
	Media_channel    string `gorm:"column:media_channel"`
	Location         string `gorm:"column:location"`
	Location_path    string `gorm:"column:location_path"`
	Collection_tool  string `gorm:"column:collection_tool"`
	User             string `gorm:"column:user"`
	Site_url         string `gorm:"column:site_url"`
	Leaf_id          string `gorm:"column:leaf_id"`
	Task_id          string `gorm:"column:task_id"`
	Task_name        string `gorm:"column:task_name"`
	Get_time         string `gorm:"column:get_time"`
	Keyword          string `gorm:"column:keyword"`
	Pub_time         string `gorm:"column:pub_time"`
	Project_id       string `gorm:"column:project_id"`
	Error_msg        string `gorm:"column:error_msg"`
}

type BISHOUZHAN_FENGNIAO struct {
	ID                string `gorm:"column:id"`
	MEDIA_TYPE        string `gorm:"column:media_type"`
	URL               string `gorm:"column:url"`
	TITLE             string `gorm:"column:title"`
	ORIGINAL_TITLE    string `gorm:"column:original_title"`
	ABSTRACT_DATA     string `gorm:"column:abstract_data"`
	CONTENT           string `gorm:"column:content"`
	MEDIA             string `gorm:"column:media"`
	MEDIA_CHANNEL     string `gorm:"column:media_channel"`
	MEDIA_CHANNEL_SND string `gorm:"column:media_channel_snd"`
	MEDIA_CHANNEL_TRD string `gorm:"column:media_channel_trd"`
	LOCATION          string `gorm:"column:location"`
	LOCATIONPATH      string `gorm:"column:locationPath"`
	SOURCE            string `gorm:"column:source"`
	RELEASE_DATETIME  string `gorm:"column:release_datetime"`
	CONTAINS_PICTURES string `gorm:"column:contains_pictures"`
	CONTAINS_VIDEOS   string `gorm:"column:contains_videos"`
	EDITOR            string `gorm:"column:editor"`
	AUTHOR            string `gorm:"column:author"`
	REPORTER          string `gorm:"column:reporter"`
	CONTENTS          string `gorm:"column:contents"`
	READING_TIMES     string `gorm:"column:reading_times"`
	ANALYSIS_KEYWORD  string `gorm:"column:analysis_keyword"`
	SITE_URL          string `gorm:"column:site_url"`
	GET_TIME          string `gorm:"column:get_time"`
	HOTSPOT           string `gorm:"column:hotspot"`
	REPOSTS_COUNT     string `gorm:"column:reposts_count"`
	PRAISE_COUNT      string `gorm:"column:praise_count"`
	TERMINAL          string `gorm:"column:terminal"`
	USERS_TYPE        string `gorm:"column:users_type"`
	AGE               int    `gorm:"column:age"`
	GENDER            string `gorm:"column:gender"`
	REGION            string `gorm:"column:region"`
	ATTENTION_COUNT   string `gorm:"column:attention_count"`
	FANS_COUNT        string `gorm:"column:fans_count"`
	STATUSES_COUNT    string `gorm:"column:statuses_count"`
	TEXT_TYPE         string `gorm:"column:text_type"`
	ARTICLE_COUNT     string `gorm:"column:article_count"`
	AVGREADING_TIMES  string `gorm:"column:avgreading_times"`
	PUBLIC_FUNCTION   string `gorm:"column:public_function"`
	PUBLIC_SUBJECT    string `gorm:"column:public_subject"`
	TIEBA_FOLLOWERS   string `gorm:"column:tieba_followers"`
	COLLECTION_TOOL   string `gorm:"column:collection_tool"`
	USER              string `gorm:"column:user"`
	KEYWORD           string `gorm:"column:keyword"`
	ATTITUDE          string `gorm:"column:attitude"`
	COMPANY_NAME      string `gorm:"column:company_name"`
	CLASSIFICATION    string `gorm:"column:classification"`
	MEDIA_REGION      string `gorm:"column:media_region"`
	CENTRAL_MEDIA     string `gorm:"column:central_media"`
	PRODUCT           string `gorm:"column:product"`
	PRODUCT_ISNEW     string `gorm:"column:product_isnew"`
	TECHNOLOGY        string `gorm:"column:technology"`
	TECHNOLOGY_ISNEW  string `gorm:"column:technology_isnew"`
	HEAT              string `gorm:"column:heat"`
	REPRODUCED_COUNT  string `gorm:"column:reproduced_count"`
	SIMILARITY        string `gorm:"column:similarity"`
	RELATION          string `gorm:"column:relation"`
	NEWS_AUDIENCES    string `gorm:"column:news_audiences"`
	PUBLICITY_VALUES  string `gorm:"column:publicity_values"`
	MEDIA_RANK        string `gorm:"column:media_rank"`
	WORD_COUNT        string `gorm:"column:word_count"`
	LEAFID            string `gorm:"column:leafid"`
	SOURCECODEID      string `gorm:"column:sourcecodeid"`
	TASKID            string `gorm:"column:taskid"`
	TASKNAME          string `gorm:"column:taskname"`
	SYNSTATUS         string `gorm:"column:synstatus"`
	ISREPEAT          string `gorm:"column:isRepeat"`
	MARK              int    `gorm:"column:mark"`
	FLAG              int    `gorm:"column:flag"`
}

var (
	db186 *gorm.DB
	db168 *gorm.DB
	err   error
)

const (
	//TimeFormart = "2006-01-02 15:04:05"
	TimeFormart = "20060102"
)

func main() {
	db_user_S := "funbird"
	db_pass_S := "funbird2017"
	db_host_S := "192.168.95.186"
	db_port_S := 3306
	//db_name_S := "funbird"
	db_name_S := "funbird_storage"

	dbConnectBaseStr := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnect_S := fmt.Sprintf(dbConnectBaseStr, db_user_S, db_pass_S, db_host_S, db_port_S, db_name_S)
	fmt.Print(dbConnect_S, "\n")
	db186, err = gorm.Open("mysql", dbConnect_S)
	if err != nil {
		fmt.Println(err)
	}
	defer db186.Close()
	db186.LogMode(true)

	db_user_D := "iscloud"
	db_pass_D := "iscloud"
	db_host_D := "192.168.182.168"
	db_port_D := 3306
	db_name_D := "test1"

	dbConnect_D := fmt.Sprintf(dbConnectBaseStr, db_user_D, db_pass_D, db_host_D, db_port_D, db_name_D)
	fmt.Print(dbConnect_D, "\n")
	db168, err = gorm.Open("mysql", dbConnect_D)
	if err != nil {
		fmt.Println(err)
	}
	defer db168.Close()
	db168.LogMode(true)

	now := time.Now()
	diff, err := time.ParseDuration("-24h")
	if err != nil {
		fmt.Println(err)
	}
	yd := now.Add(diff).Format(TimeFormart)

	inforall := []*INFORMATRIONALL{}
	db186.Table(fmt.Sprintf("%s_%s", "information_all_d7178819", yd)).Where("pub_time != ''").Find(&inforall)
	for _, v := range inforall {
		data := BISHOUZHAN_FENGNIAO{}

		data.ID = v.Leaf_id
		data.URL = v.Url
		data.TITLE = v.Title
		data.AUTHOR = v.Author
		data.SOURCE = v.Source
		data.RELEASE_DATETIME = v.Pub_time
		data.CONTENT = v.Content
		data.MEDIA_TYPE = v.Media_type
		data.ORIGINAL_TITLE = v.Original_title
		data.EDITOR = v.Editor
		data.REPORTER = v.Reporter
		data.CONTENTS = v.Contents
		data.READING_TIMES = v.Reading_times
		data.ABSTRACT_DATA = v.Abstract_data
		data.MEDIA = v.Media
		data.MEDIA_CHANNEL = v.Media_channel
		data.LOCATION = v.Location
		data.LOCATIONPATH = v.Location_path
		data.COLLECTION_TOOL = v.Collection_tool
		data.USER = v.User
		data.SITE_URL = v.Site_url
		data.LEAFID = v.Leaf_id
		data.TASKID = v.Task_id
		data.TASKNAME = v.Task_name
		data.GET_TIME = v.Get_time
		data.KEYWORD = v.Keyword
		//data.TAG = "政策"
		data.MEDIA_TYPE = "新闻"
		db168.Table("bishouzhan_fengniao").Create(&data)
	}
}
