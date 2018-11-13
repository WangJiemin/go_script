package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type InForMation_All_5e619c69 struct {
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
}

func (InForMation_All_5e619c69) TableName() string {
	return "information_all_5e619c69"
}

type InForMation_All_5e619c69_1 struct {
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
	//error_msg        string
}

func (InForMation_All_5e619c69_1) TableName() string {
	return "information_all_5e619c69_1"
}

var db_5e619c69 *gorm.DB
var db_5e619c69_1 *gorm.DB

func main() {
	/*
		db_user := "funbird"
		db_pass := "funbird2017"
		db_host := "10.10.206.205"
		db_port := 3306
		db_name := "funbird_analysis"
		dbConnectBaseStr := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
		dbConnect := fmt.Sprintf(dbConnectBaseStr, db_user, db_pass, db_host, db_port, db_name)
		fmt.Print(dbConnect)
		db_5e619c69, err := gorm.Open("mysql", dbConnect)
		if err != nil {
			fmt.Println(err)
			//fmt.Printf("Connection to MySQL database is Name:%s Failed！\n", dbInformationAll)
		}
		defer db_5e619c69.Close()
		db_5e619c69.LogMode(true)

		db_user_t := "funbird"
		db_pass_t := "funbird2017"
		db_host_t := "10.10.206.205"
		db_port_t := 3306
		db_name_t := "funbird_analysis"
		dbConnectBaseStr_t := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
		dbConnect_t := fmt.Sprintf(dbConnectBaseStr_t, db_user_t, db_pass_t, db_host_t, db_port_t, db_name_t)
		fmt.Print(dbConnect_t)
		db_5e619c69_1, err := gorm.Open("mysql", dbConnect_t)
		if err != nil {
			fmt.Println(err)
			//fmt.Printf("Connection to MySQL database is Name:%s Failed！\n", dbInformationAll)
		}
		defer db_5e619c69_1.Close()
		db_5e619c69_1.LogMode(true)
	*/

	db_user := "jiemin"
	db_pass := "jiemin2017"
	db_host := "127.0.0.1"
	db_port := 3306
	db_name := "jiemin"
	dbConnectBaseStr := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnect := fmt.Sprintf(dbConnectBaseStr, db_user, db_pass, db_host, db_port, db_name)
	fmt.Print(dbConnect)
	db_5e619c69, err := gorm.Open("mysql", dbConnect)
	if err != nil {
		fmt.Println(err)
		//fmt.Printf("Connection to MySQL database is Name:%s Failed！\n", dbInformationAll)
	}
	defer db_5e619c69.Close()
	db_5e619c69.LogMode(true)

	db_user_t := "jiemin"
	db_pass_t := "jiemin2017"
	db_host_t := "127.0.0.1"
	db_port_t := 3306
	db_name_t := "jiemin"
	dbConnectBaseStr_t := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnect_t := fmt.Sprintf(dbConnectBaseStr_t, db_user_t, db_pass_t, db_host_t, db_port_t, db_name_t)
	fmt.Print(dbConnect_t)
	db_5e619c69_1, err := gorm.Open("mysql", dbConnect_t)
	if err != nil {
		fmt.Println(err)
		//fmt.Printf("Connection to MySQL database is Name:%s Failed！\n", dbInformationAll)
	}
	defer db_5e619c69_1.Close()
	db_5e619c69_1.LogMode(true)

	table_inforall := []*InForMation_All_5e619c69{}

	for i := 0; i < 100; {
		a := i + 4
		if a < 200 {
			db_5e619c69.Table("information_all_5e619c69").Where("id > ? and id <= ?", i, a).Find(&table_inforall)
			//SQLString := "insert into information_all_5e619c69_1 (id , url , title , author , source, release_datetime, content, media_type, original_title, editor, reporter, contents, reading_times , abstract_data, media, media_channel, location, location_path, collection_tool, user, site_url , leaf_id , task_id, task_name, get_time, keyword, pub_time, project_id) VALUES"
			SQLString := "insert into information_all_5e619c69_1 (title) VALUES"
			//SQL_EXEC := "(%d,'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s'),"

			for _, v := range table_inforall {
				/*
					table_inforall_t := InForMation_All_5e619c69_1{}
					table_inforall_t.Id = v.Id
					table_inforall_t.Url = v.Url
					table_inforall_t.Title = v.Title
					table_inforall_t.Author = v.Author
					table_inforall_t.Source = v.Source
					table_inforall_t.Release_datetime = v.Release_datetime
					table_inforall_t.Content = v.Content
					table_inforall_t.Media_type = v.Media_type
					table_inforall_t.Original_title = v.Original_title
					table_inforall_t.Editor = v.Editor
					table_inforall_t.Reporter = v.Reporter
					table_inforall_t.Contents = v.Contents
					table_inforall_t.Reading_times = v.Reading_times
					table_inforall_t.Abstract_data = v.Abstract_data
					table_inforall_t.Media = v.Media
					table_inforall_t.Media_channel = v.Media_channel
					table_inforall_t.Location = v.Location
					table_inforall_t.Location_path = v.Location_path
					table_inforall_t.Collection_tool = v.Collection_tool
					table_inforall_t.User = v.User
					table_inforall_t.Site_url = v.Site_url
					table_inforall_t.Leaf_id = v.Leaf_id
					table_inforall_t.Task_id = v.Task_id
					table_inforall_t.Task_name = v.Task_name
					table_inforall_t.Get_time = v.Get_time
					table_inforall_t.Keyword = v.Keyword
					table_inforall_t.Pub_time = v.Pub_time
					table_inforall_t.Project_id = v.Project_id

					SQL_EXEC := "insert into information_all_5e619c69_1(id , url , title , author , source, release_datetime, content, media_type, original_title, editor, reporter, contents, reading_times , abstract_data, media, media_channel, location, location_path, collection_tool, user, site_url , leaf_id , task_id, task_name, get_time, keyword, pub_time, project_id) VALUES(%d,'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')"
					//db_5e619c69_1.Table("information_all_5e619c69_1").Save(&table_inforall_t)
					db_5e619c69_1.Exec(fmt.Sprintf(SQL_EXEC, v.Id, v.Url, v.Title, v.Author, v.Source, v.Release_datetime, v.Content, v.Media_type, v.Original_title, v.Editor, v.Reporter, v.Contents, v.Reading_times, v.Abstract_data, v.Media, v.Media_channel, v.Location, v.Location_path, v.Collection_tool, v.User, v.Site_url, v.Leaf_id, v.Task_id, v.Task_name, v.Get_time, v.Keyword, v.Pub_time, v.Project_id))
				*/

				//SQL_EXEC := "(%d,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s),"
				//SQLString = fmt.Sprintf("%s%s", SQLString, fmt.Sprintf(SQL_EXEC, v.Id, v.Url, v.Title, v.Author, v.Source, v.Release_datetime, v.Content, v.Media_type, v.Original_title, v.Editor, v.Reporter, v.Contents, v.Reading_times, v.Abstract_data, v.Media, v.Media_channel, v.Location, v.Location_path, v.Collection_tool, v.User, v.Site_url, v.Leaf_id, v.Task_id, v.Task_name, v.Get_time, v.Keyword, v.Pub_time, v.Project_id))

				//valueString := fmt.Sprintf(SQL_EXEC, v.Id, v.Url, v.Title, v.Author, v.Source, v.Release_datetime, v.Content, v.Media_type, v.Original_title, v.Editor, v.Reporter, v.Contents, v.Reading_times, v.Abstract_data, v.Media, v.Media_channel, v.Location, v.Location_path, v.Collection_tool, v.User, v.Site_url, v.Leaf_id, v.Task_id, v.Task_name, v.Get_time, v.Keyword, v.Pub_time, v.Project_id)
				valueString := "(" + "'" + v.Title + "'" + ")"
				SQLString = SQLString + valueString + ","
			}
			fmt.Print(SQLString, "\n")
			//db_5e619c69_1.Exec(SQLString)
		}
		i = i + 4
	}
}
