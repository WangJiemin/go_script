package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type FUNBIRDTABLE struct {
	ID    int    `gorm:"column:id"`
	NAME  string `gorm:"column:tablename"`
	NUMID int    `gorm:"column:numid"`
}

type BAIDUNEWS struct {
	ID                string `gorm:"column:id"`
	Media_type        string `gorm:"column:media_type"`
	Url               string `gorm:"column:url"`
	Title             string `gorm:"column:title"`
	Original_title    string `gorm:"column:original_title"`
	Abstract_data     string `gorm:"column:abstract_data"`
	Content           string `gorm:"column:content"`
	Media             string `gorm:"column:media"`
	Media_channel     string `gorm:"column:media_channel"`
	Media_channel_snd string `gorm:"column:media_channel_snd"`
	Media_channel_trd string `gorm:"column:media_channel_trd"`
	Location          string `gorm:"column:location"`
	LocationPath      string `gorm:"column:locationPath"`
	Source            string `gorm:"column:source"`
	Release_datetime  string `gorm:"column:release_datetime"`
	Contains_pictures string `gorm:"column:contains_pictures"`
	Contains_videos   string `gorm:"column:contains_videos"`
	Editor            string `gorm:"column:editor"`
	Author            string `gorm:"column:author"`
	Reporter          string `gorm:"column:reporter"`
	Contents          string `gorm:"column:contents"`
	Reading_times     string `gorm:"column:reading_times"`
	Tag               string `gorm:"column:tag"`
	Analysis_keyword  string `gorm:"column:analysis_keyword"`
	Site_url          string `gorm:"column:site_url"`
	Get_time          string `gorm:"column:get_time"`
	Hotspot           string `gorm:"column:hotspot"`
	Reposts_count     string `gorm:"column:reposts_count"`
	Praise_count      string `gorm:"column:praise_count"`
	Terminal          string `gorm:"column:terminal"`
	Users_type        string `gorm:"column:users_type"`
	Gender            string `gorm:"column:gender"`
	Region            string `gorm:"column:region"`
	Attention_count   string `gorm:"column:attention_count"`
	Fans_count        string `gorm:"column:fans_count"`
	Statuses_count    string `gorm:"column:statuses_count"`
	Text_type         string `gorm:"column:text_type"`
	Article_count     string `gorm:"column:article_count"`
	Avgreading_times  string `gorm:"column:avgreading_times"`
	Public_function   string `gorm:"column:public_function"`
	Public_subject    string `gorm:"column:public_subject"`
	Tieba_followers   string `gorm:"column:tieba_followers"`
	Collection_tool   string `gorm:"column:collection_tool"`
	User              string `gorm:"column:user"`
	Keyword           string `gorm:"column:keyword"`
	Attitude          string `gorm:"column:attitude"`
	Company_name      string `gorm:"column:company_name"`
	Classification    string `gorm:"column:classification"`
	Media_region      string `gorm:"column:media_region"`
	Central_media     string `gorm:"column:central_media"`
	Product           string `gorm:"column:product"`
	Product_isnew     string `gorm:"column:product_isnew"`
	Technology        string `gorm:"column:technology"`
	Technology_isnew  string `gorm:"column:technology_isnew"`
	Heat              string `gorm:"column:heat"`
	Reproduced_count  string `gorm:"column:reproduced_count"`
	Similarity        string `gorm:"column:similarity"`
	Relation          string `gorm:"column:relation"`
	News_audiences    string `gorm:"column:news_audiences"`
	Publicity_values  string `gorm:"column:publicity_values"`
	Media_rank        string `gorm:"column:media_rank"`
	Word_count        string `gorm:"column:word_count"`
	Leafid            string `gorm:"column:leafid"`
	Sourcecodeid      string `gorm:"column:sourcecodeid"`
	Taskid            string `gorm:"column:taskid"`
	Taskname          string `gorm:"column:taskname"`
	Synstatus         string `gorm:"column:synstatus"`
	IsRepeat          string `gorm:"column:isRepeat"`
	Mark              int    `gorm:"column:mark"`
	ReadStatus        int    `gorm:"column:readStatus"`
}

type INFOR_88a18cd9 struct {
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

func (INFOR_88a18cd9) TableName() string {
	return "information_all_88a18cd9"
}

func (BAIDUNEWS) TableName() string {
	return "baidunews"
}

var db124, db140 *gorm.DB
var err error

func main() {
	db_user_S := "funbird"
	db_pass_S := "funbird2017"
	db_host_S := "192.168.95.186"
	db_port_S := 3306
	//db_name_S := "funbird"
	db_name_S := "funbird_storage"

	dbConnectBaseStr_S := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnect_S := fmt.Sprintf(dbConnectBaseStr_S, db_user_S, db_pass_S, db_host_S, db_port_S, db_name_S)
	fmt.Print(dbConnect_S)
	db124, err = gorm.Open("mysql", dbConnect_S)
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

	dbConnectBaseStr_D := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnect_D := fmt.Sprintf(dbConnectBaseStr_D, db_user_D, db_pass_D, db_host_D, db_port_D, db_name_D)
	fmt.Print(dbConnect_D)
	db140, err = gorm.Open("mysql", dbConnect_D)
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
		db124.Table("funbird_analysis.funbird_table").Where("tablename = ?", "information_all_88a18cd9").Select("numid").Find(&funbridtable)
		for _, v := range funbridtable {
			numid = v.NUMID
		}
		fmt.Printf("update table name is:\t%s\tAUTO INCREMENT number is:\t%d\n", "information_all_88a18cd9", numid)
	*/

	infor := []*INFOR_88a18cd9{}
	//db124.Table("information_all_88a18cd9").Where("id > ? and pub_time!=''", numid).Find(&infor)
	db124.Table(fmt.Sprintf("%s_%s", "information_all_88a18cd9", yesterday)).Where("pub_time != ''").Find(&infor)
	for _, v := range infor {
		data := BAIDUNEWS{}
		data.ID = v.Leaf_id
		data.Url = v.Url
		data.Title = v.Title
		data.Author = v.Author
		data.Source = v.Source
		data.Release_datetime = v.Pub_time
		data.Content = v.Content
		data.Media_type = v.Media_type
		data.Original_title = v.Original_title
		data.Editor = v.Editor
		data.Reporter = v.Reporter
		data.Contents = v.Contents
		data.Reading_times = v.Reading_times
		data.Abstract_data = v.Abstract_data
		data.Media = v.Media
		data.Media_channel = v.Media_channel
		data.Location = v.Location
		data.LocationPath = v.Location_path
		data.Collection_tool = v.Collection_tool
		data.User = v.User
		data.Site_url = v.Site_url
		data.Leafid = v.Leaf_id
		data.Taskid = v.Task_id
		data.Taskname = v.Task_name
		data.Get_time = v.Get_time
		data.Keyword = v.Keyword
		data.Tag = "政策"
		data.Media_type = "新闻"

		//fmt.Printf("informationAll_88 tables ID NUM is %d\n", v.ID)
		db140.Table("baidunews").Save(&data)
		//db124.Table("funbird_analysis.funbird_table").Where("tablename = ?", "information_all_88a18cd9").Update("numid", v.ID)
	}
}
