package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type TASKRESULTS struct {
	//Id       int    `gorm:"column:id"`
	TASKID   string    `grom:"column:task_id"`
	DATAMD5  string    `gorm:"column:data_md5"`
	URL      string    `gorm:"column:url"`
	CONTENT  string    `gorm:"column:content"`
	ISERROR  int       `gorm:"column:is_error"`
	ISDELETE int       `gorm:"column:is_delete"`
	GETTIME  time.Time `gorm:"column:get_time"`
	DOMAIN   string    `gorm:"column:domain"`
}

type TASKINFOS struct {
	TASKID string `gorm:"column:task_id"`
}

const (
	TimeFormart = "2006-01-02 15:04:05"
)

var (
	db_s *gorm.DB
	db_t *gorm.DB
	err  error
)

func main() {

	dbConnectBaseStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	db_host_s := "192.168.181.124"
	db_port_s := "3306"
	db_user_s := "funbird"
	db_pwd_s := "funbird2017"
	db_name_s := "funbird"
	dbConnect_s := fmt.Sprintf(dbConnectBaseStr, db_user_s, db_pwd_s, db_host_s, db_port_s, db_name_s)
	fmt.Print(dbConnect_s, "\t")
	db_s, err = gorm.Open("mysql", dbConnect_s)
	if err != nil {
		fmt.Println(err)
	}
	defer db_s.Close()
	db_s.LogMode(true)

	db_host_t := "117.107.241.67"
	db_port_t := "3306"
	db_user_t := "spider"
	db_pwd_t := "rosqvv*Pp6Uz9ca"
	db_name_t := "datacenter_spider"
	dbConnect_t := fmt.Sprintf(dbConnectBaseStr, db_user_t, db_pwd_t, db_host_t, db_port_t, db_name_t)
	fmt.Print(dbConnect_t, "\n")
	db_t, err = gorm.Open("mysql", dbConnect_t)
	if err != nil {
		fmt.Println(err)
	}
	defer db_t.Close()
	db_t.LogMode(true)

	taskinfo := []*TASKINFOS{}
	db_s.Table("task_infos").Where("project_id='a69fe420-9671-4859-a493-54a873da34dd' and `name` like '万得-%'").Select("task_id").Scan(&taskinfo)
	taskinfo = append(taskinfo, TASKINFOS{taskid: "215c67e0-3b7d-40eb-b08e-25a42663a202"})
	for _, taskid := range taskinfo {
		taskresults := []*TASKRESULTS{}
		db_s.Raw(fmt.Sprintf("select * from task_results where task_id = '%s'", taskid.TASKID)).Scan(&taskresults)
		for _, v := range taskresults {
			//datamd5results := []*TASKRESULTS{}
			//db_t.Table("task_results").Select("data_md5").Scan(&datamd5results)
			//for _, vv := range datamd5results {
			//if vv.DATAMD5 != v.DATAMD5 {
			db_t.Table("task_results").Save(&v)
			//}
			//}
		}
	}
}
