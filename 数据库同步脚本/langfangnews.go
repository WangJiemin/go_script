package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type InForMationAll struct {
	Id               int
	Url              string
	Title            string
	Author           string
	Source           string
	Release_datetime string
	Content          string
	Media_type       string
	Original_title   string
	Editor           string
	Reporter         string
	Contents         string
	Reading_times    string
	Abstract_data    string
	Media            string
	Media_channel    string
	Location         string
	Location_path    string
	Collection_tool  string
	User             string
	Site_url         string
	Leaf_id          string
	Task_id          string
	Task_name        string
	Get_time         string
	Keyword          string
	Pub_time         string
	Project_id       string
	Error_msg        string
}

type FUNBIRDTABLE struct {
	ID    int    `gorm:"column:id"`
	NAME  string `gorm:"column:tablename"`
	NUMID int    `gorm:"column:numid"`
}

type NEWS struct {
	ID                string `gorm:"cloumn:id"`
	MEDIA_TYPE        string `gorm:"cloumn:media_type"`
	URL               string `gorm:"cloumn:url"`
	TITLE             string `gorm:"cloumn:title"`
	ORIGINAL_TITLE    string `gorm:"cloumn:original_title"`
	ABSTRACT_DATA     string `gorm:"cloumn:abstract_data"`
	CONTENT           string `gorm:"cloumn:content"`
	MEDIA             string `gorm:"cloumn:media"`
	MEDIA_CHANNEL     string `gorm:"cloumn:media_channel"`
	MEDIA_CHANNEL_SND string `gorm:"cloumn:media_channel_snd"`
	MEDIA_CHANNEL_TRD string `gorm:"cloumn:media_channel_trd"`
	LOCATION          string `gorm:"cloumn:location"`
	LOCATIONPATH      string `gorm:"cloumn:locationPath"`
	SOURCE            string `gorm:"cloumn:source"`
	RELEASE_DATETIME  string `gorm:"cloumn:release_datetime"`
	CONTAINS_PICTURES string `gorm:"cloumn:contains_pictures"`
	ContainsVideos    string `gorm:"cloumn:contains_videos"`
	EDITOR            string `gorm:"cloumn:editor"`
	AUTHOR            string `gorm:"cloumn:author"`
	REPORTER          string `gorm:"cloumn:reporter"`
	CONTENTS          string `gorm:"cloumn:contents"`
	READING_TIMES     string `gorm:"cloumn:reading_times"`
	TAG               string `gorm:"cloumn:tag"`
	ANALYSIS_KEYWORD  string `gorm:"cloumn:analysis_keyword"`
	SITE_URL          string `gorm:"cloumn:site_url"`
	GET_TIME          string `gorm:"cloumn:get_time"`
	HOTSPOT           string `gorm:"cloumn:hotspot"`
	REPOSTS_COUNT     string `gorm:"cloumn:reposts_count"`
	PRAISE_COUNT      string `gorm:"cloumn:praise_count"`
	TERMINAL          string `gorm:"cloumn:terminal"`
	USERS_TYPE        string `gorm:"cloumn:users_type"`
	GENDER            string `gorm:"cloumn:gender"`
	REGION            string `gorm:"cloumn:region"`
	ATTENTION_COUNT   string `gorm:"cloumn:attention_count"`
	FANS_COUNT        string `gorm:"cloumn:fans_count"`
	STATUSES_COUNT    string `gorm:"cloumn:statuses_count"`
	TEXT_TYPE         string `gorm:"cloumn:text_type"`
	ARTICLE_COUNT     string `gorm:"cloumn:article_count"`
	AVGREADING_TIMES  string `gorm:"cloumn:avgreading_times"`
	PUBLIC_FUNCTION   string `gorm:"cloumn:public_function"`
	PUBLIC_SUBJECT    string `gorm:"cloumn:public_subject"`
	TIEBA_FOLLOWERS   string `gorm:"cloumn:tieba_followers"`
	COLLECTION_TOOL   string `gorm:"cloumn:collection_tool"`
	USER              string `gorm:"cloumn:user"`
	KEYWORD           string `gorm:"cloumn:keyword"`
	ATTITUDE          string `gorm:"cloumn:attitude"`
	COMPANY_NAME      string `gorm:"cloumn:company_name"`
	CLASSIFICATION    string `gorm:"cloumn:classification"`
	MEDIA_REGION      string `gorm:"cloumn:media_region"`
	CENTRAL_MEDIA     string `gorm:"cloumn:central_media"`
	PRODUCT           string `gorm:"cloumn:product"`
	PRODUCT_ISNEW     string `gorm:"cloumn:product_isnew"`
	TECHNOLOGY        string `gorm:"cloumn:technology"`
	TECHNOLOGY_ISNEW  string `gorm:"cloumn:technology_isnew"`
	HEAT              string `gorm:"cloumn:heat"`
	REPRODUCED_COUNT  string `gorm:"cloumn:reproduced_count"`
	SIMILARITY        string `gorm:"cloumn:similarity"`
	RELATION          string `gorm:"cloumn:relation"`
	NEWS_AUDIENCES    string `gorm:"cloumn:news_audiences"`
	PUBLICITY_VALUES  string `gorm:"cloumn:publicity_values"`
	MEDIA_RANK        string `gorm:"cloumn:media_rank"`
	WORD_COUNT        string `gorm:"cloumn:word_count"`
	Leafid            string `gorm:"cloumn:leafid"`
	Sourcecodeid      string `gorm:"cloumn:sourcecodeid"`
	Taskid            string `gorm:"cloumn:taskid"`
	TASKNAME          string `gorm:"cloumn:taskname"`
	SYNSTATUS         string `gorm:"cloumn:synstatus"`
	ISREPEAT          string `gorm:"cloumn:isRepeat"`
	MARK              int    `gorm:"cloumn:mark"`
	READSTATUS        int    `gorm:"cloumn:readStatus"`
}

func (NEWS) TableName() string {
	return "langfangnews"
}

var (
	db124 *gorm.DB
	db140 *gorm.DB
	err   error
)

func main() {
	db_user_S := "funbird"
	db_pass_S := "funbird2017"
	db_host_S := "192.168.95.186"
	db_port_S := 3306
	//db_name_S := "funbird"
	db_name_S := "funbird_storage"

	dbConnectBaseStr := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnect := fmt.Sprintf(dbConnectBaseStr, db_user_S, db_pass_S, db_host_S, db_port_S, db_name_S)
	fmt.Print(dbConnect)
	db124, err = gorm.Open("mysql", dbConnect)
	if err != nil {
		fmt.Println(err)
	}
	defer db124.Close()
	db124.LogMode(true)

	db_user_D := "iscloud"
	db_pass_D := "iscloud"
	db_host_D := "192.168.182.168"
	db_port_D := 3306
	db_name_D := "network_bigdata"
	dbConnectBaseStr_b := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnect_b := fmt.Sprintf(dbConnectBaseStr_b, db_user_D, db_pass_D, db_host_D, db_port_D, db_name_D)
	fmt.Print(dbConnect_b)
	db140, err = gorm.Open("mysql", dbConnect_b)
	if err != nil {
		fmt.Println(err)
	}
	defer db140.Close()

	db140.LogMode(true)

	now := time.Now()
	diff, err := time.ParseDuration("-24h")
	if err != nil {
		fmt.Println(err)
	}
	yesterday := now.Add(diff).Format("20060102")

	/*
		var numid int
		funbridtable := []*FUNBIRDTABLE{}
		db124.Table("funbird_analysis.funbird_table").Where("tablename = ?", "information_all_b42661db").Select("numid").Find(&funbridtable)
		for _, v := range funbridtable {
			numid = v.NUMID
		}
		fmt.Printf("update table name is:\t%s\tAUTO INCREMENT number is:\t%d\n", "information_all_b42661db", numid)
	*/

	inforall4b := []*InForMationAll{}

	//db124.Table("information_all_b42661db").Where("id > ? and title IS NOT NULL and title != '' and content IS NOT NULL and content != ''", numid).Find(&inforall4b)
	db124.Table(fmt.Sprintf("%s_%s", "information_all_b42661db", yesterday)).Where("title IS NOT NULL and title != '' and content IS NOT NULL and content != ''").Find(&inforall4b)

	for _, v := range inforall4b {
		//fmt.Print(v)
		langfannews := NEWS{}
		langfannews.ID = v.Leaf_id
		langfannews.URL = v.Url
		langfannews.TITLE = v.Title
		langfannews.AUTHOR = v.Author
		langfannews.SOURCE = v.Source
		langfannews.RELEASE_DATETIME = v.Pub_time
		langfannews.CONTENT = v.Content
		langfannews.MEDIA_TYPE = v.Media_type
		langfannews.ORIGINAL_TITLE = v.Original_title
		langfannews.EDITOR = v.Editor
		langfannews.REPORTER = v.Reporter
		langfannews.CONTENTS = v.Contents
		langfannews.READING_TIMES = v.Reading_times
		langfannews.ABSTRACT_DATA = v.Abstract_data
		langfannews.MEDIA = v.Media
		langfannews.MEDIA_CHANNEL = v.Media_channel
		langfannews.LOCATION = v.Location
		langfannews.LOCATIONPATH = v.Location_path
		langfannews.COLLECTION_TOOL = v.Collection_tool
		langfannews.USER = v.User
		langfannews.SITE_URL = v.Site_url
		langfannews.Leafid = v.Leaf_id
		langfannews.Taskid = v.Task_id
		langfannews.TASKNAME = v.Task_name
		langfannews.GET_TIME = v.Get_time
		langfannews.KEYWORD = v.Keyword
		langfannews.MARK = 0
		langfannews.READSTATUS = 0

		if v.Pub_time == "" {
			langfannews.RELEASE_DATETIME = v.Get_time
		}
		db140.Table("langfangnews").Save(&langfannews)
		//db124.Table("funbird_analysis.funbird_table").Where("tablename = ?", "information_all_b42661db").Update("numid", v.Id)
	}
}
