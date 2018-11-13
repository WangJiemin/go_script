package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type InForMation_All_b42661db struct {
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
	//Project_id       string
}

func (InForMation_All_b42661db) TableName() string {
	return "information_all_b42661db"
}

type BaiDuNews struct {
	Id               string
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
	Locationpath     string
	Collection_tool  string
	User             string
	Site_url         string
	Leafid           string
	Taskid           string
	Taskname         string
	Get_time         string
	Keyword          string
	//Pub_time         string
	//Project_id       string
}

func (BaiDuNews) TableName() string {
	return "baidunews"
	//return "baidunews_jiemin"
}

var dbInformationAll *gorm.DB

var dbBaiDuNews *gorm.DB

func main() {
	/*
		db_user := "jiemin"
		db_pass := "jiemin2017"
		db_host := "127.0.0.1"
		db_port := 3306
		db_name := "jiemin"
	*/

	db_user := "funbird"
	db_pass := "funbird2017"
	db_host := "10.10.206.205"
	db_port := 3306
	db_name := "funbird_analysis"

	dbConnectBaseStr := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnect := fmt.Sprintf(dbConnectBaseStr, db_user, db_pass, db_host, db_port, db_name)
	fmt.Print(dbConnect)
	dbInformationAll, err := gorm.Open("mysql", dbConnect)
	if err != nil {
		fmt.Println(err)
		//fmt.Printf("Connection to MySQL database is Name:%s Failed！\n", dbInformationAll)
	}
	defer dbInformationAll.Close()
	dbInformationAll.LogMode(true)

	/*
		db_user_b := "jiemin"
		db_pass_b := "jiemin2017"
		db_host_b := "127.0.0.1"
		db_port_b := 3306
		db_name_b := "jiemin"
	*/

	db_user_b := "iscloud"
	db_pass_b := "iscloud"
	db_host_b := "192.168.95.140"
	db_port_b := 3306
	db_name_b := "network_data"

	dbConnectBaseStr_b := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnect_b := fmt.Sprintf(dbConnectBaseStr_b, db_user_b, db_pass_b, db_host_b, db_port_b, db_name_b)
	fmt.Print(dbConnect_b)
	dbBaiDuNews, err := gorm.Open("mysql", dbConnect_b)
	if err != nil {
		fmt.Println(err)
		//fmt.Printf("Connection to MySQL database is Name:%s Failed！\n", dbBaiDuNews)
	}
	defer dbBaiDuNews.Close()

	dbBaiDuNews.LogMode(true)

	inforall := []*InForMation_All_b42661db{}
	//baiDNews := []*BaiDuNews{}

	dbInformationAll.Table("information_all_b42661db").Where("pub_time is not null or pub_time !=''").Find(&inforall)
	//dbBanknews.Table("baidunews").Count(&baiDNews)

	for _, v := range inforall {
		//fmt.Print(v)
		baiDNews := BaiDuNews{}
		baiDNews.Id = v.Leaf_id
		baiDNews.Url = v.Url
		baiDNews.Title = v.Title
		baiDNews.Author = v.Author
		baiDNews.Source = v.Source
		baiDNews.Release_datetime = v.Pub_time
		baiDNews.Content = v.Content
		baiDNews.Media_type = v.Media_type
		baiDNews.Original_title = v.Original_title
		baiDNews.Editor = v.Editor
		baiDNews.Reporter = v.Reporter
		baiDNews.Contents = v.Contents
		baiDNews.Reading_times = v.Reading_times
		baiDNews.Abstract_data = v.Abstract_data
		baiDNews.Media = v.Media
		baiDNews.Media_channel = v.Media_channel
		baiDNews.Location = v.Location
		baiDNews.Locationpath = v.Location_path
		baiDNews.Collection_tool = v.Collection_tool
		baiDNews.User = v.User
		baiDNews.Site_url = v.Site_url
		baiDNews.Leafid = v.Leaf_id
		baiDNews.Taskid = v.Task_id
		baiDNews.Taskname = v.Task_name
		baiDNews.Get_time = v.Get_time
		baiDNews.Keyword = v.Keyword

		dbBaiDuNews.Table("baidunews").Save(&baiDNews)
		//dbBaiDuNews.Table("baidunews_jiemin").Save(&baiDNews)
	}

}
