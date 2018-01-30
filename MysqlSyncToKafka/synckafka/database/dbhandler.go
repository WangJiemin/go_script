package database

import (
	"fmt"

	"golang_Learn/synckafka/datetime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Initialize(config map[string]string) {
	var dbHost, dbUser, dbPwd, dbName, dbPort string
	dbConnentBaseStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	dbHost = config["db_host"]
	dbUser = config["db_user"]
	dbPwd = config["db_pwd"]
	dbName = config["db_name"]
	dbPort = config["db_port"]

	dbConnent := fmt.Sprintf(dbConnentBaseStr, dbUser, dbPwd, dbHost, dbPort, dbName)
	fmt.Println(dbConnent)
	db, err = gorm.Open("mysql", dbConnent)
	if err != nil {
		fmt.Println(err)
	}
	db.LogMode(true)
}

func SELECTTABLES(table string) []*InForMationTwo {
	dt := datetime.GetNowTime()
	infotwo := []*InForMationTwo{}
	db.Table(fmt.Sprintf("%s_%s", table, dt)).Scan(&infotwo)
	//for _, v := range infotwo {
	//	v.CrawlSource = "蜂鸟"
	//}
	return infotwo
}

/*
func SELECTTABLES(table string) []*INFORMATRIONALL {
	dt := datetime.GetNowTime()
	info := []*INFORMATRIONALL{}
	db.Table(fmt.Sprintf("%s_%s", table, dt)).Scan(&info)
	db.Close()
	return info
}

*/
/*
func SELECTTABLES(table string) (jsondata []byte) {
	dt := datetime.GetNowTime()
	info := []*INFORMATRIONALL{}
	db.Table(fmt.Sprintf("%s_%s", table, dt)).Scan(&info)
	db.Close()
	jsondata, _ = json.Marshal(&info)
	return jsondata
}
*/
