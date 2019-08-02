package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DiscuzTopicRelation struct {
	Id           int    `gorm:"column:id"`
	UserId       int    `grom:"column:user_id"`
	DiscussionId int    `gorm:"column:discussion_id"`
	Code         string `gorm:"column:code"`
	TopicType    int    `gorm:"column:topic_type"`
	Extra        string `gorm:"column:extra"`
	CreateTs     int    `gorm:"column:create_ts"`
	UpdateTs     int    `gorm:"column:update_ts"`
}

//const (
//	TimeFormart = "2006-01-02 15:04:05"
//)

var (
	db_s *gorm.DB
	//db_t *gorm.DB
	err error
)

func main() {

	dbConnectBaseStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	db_host_s := "10.30.2.201"
	db_port_s := "930"
	//db_host_s := "10.25.1.50"
	//db_port_s := "3306"
	db_user_s := "mico_r"
	db_pwd_s := "7P7uI7gM10JWrxAG"
	db_name_s := "pregnancy_two"
	dbConnect_s := fmt.Sprintf(dbConnectBaseStr, db_user_s, db_pwd_s, db_host_s, db_port_s, db_name_s)
	fmt.Print(dbConnect_s, "\t")
	//db_s, err = gorm.Open("mysql", dbConnect_s)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer db_s.Close()
	//db_s.DB().SetMaxIdleConns(10)
	//db_s.DB().SetMaxOpenConns(100)
	//db_s.LogMode(true)

	//db_host_t := "117.107.241.67"
	//db_port_t := "3306"
	//db_user_t := "spider"
	//db_pwd_t := "rosqvv*Pp6Uz9ca"
	//db_name_t := "datacenter_spider"
	//dbConnect_t := fmt.Sprintf(dbConnectBaseStr, db_user_t, db_pwd_t, db_host_t, db_port_t, db_name_t)
	//fmt.Print(dbConnect_t, "\n")
	//db_t, err = gorm.Open("mysql", dbConnect_t)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer db_t.Close()
	//db_t.LogMode(true)

	//t1, t2 := GetNowTime()

	// SELECT id,code FROM DiscuzTopicRelation where discussion_id = 72123734
	taskinfo := []*DiscuzTopicRelation{}

	for idnum := 72123734; idnum <= 92123734; {
		db_s, err = gorm.Open("mysql", dbConnect_s)
		if err != nil {
			fmt.Println(err)
		}
		db_s.DB().SetMaxIdleConns(100)
		db_s.DB().SetMaxOpenConns(100)
		db_s.LogMode(true)
		db_s.Table("DiscuzTopicRelation").Where("discussion_id = ? ", idnum).Scan(&taskinfo)
		db_s.Close()
		idnum = idnum + 100
		fmt.Println(taskinfo)

	}

	//db_s.Table("task_infos").Where("project_id='a69fe420-9671-4859-a493-54a873da34dd'").Select("task_id").Scan(&taskinfo)
	//for _, taskid := range taskinfo {
	//	db_s.Raw(fmt.Sprintf("select * from task_results where task_id = '%s' and get_time >= '%s' and get_time <= '%s'", taskid.TASKID, t1, t2)).Scan(&taskresults)
	//	for _, v := range taskresults {
	//		//datamd5results := []*TASKRESULTS{}
	//		//db_t.Table("task_results").Select("data_md5").Scan(&datamd5results)
	//		//for _, vv := range datamd5results {
	//		//if vv.DATAMD5 != v.DATAMD5 {
	//		db_t.Table("task_results").Save(&v)
	//		//}
	//		//}
	//	}
	//}
}

//func GetNowTime() (starttime string, stoptime string) {
//	now := time.Now()
//	yesterday, err := time.ParseDuration("-2h")
//	if err != nil {
//		log.Println(err)
//	}
//	start := now.Add(yesterday).Format(TimeFormart)
//	current := now.Format(TimeFormart)
//	return start, current
//}
