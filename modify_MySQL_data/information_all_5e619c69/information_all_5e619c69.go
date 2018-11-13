package main

import (
	"fmt"
	"strings"
	//"github.com/WangJiemin/gocomm/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type InForMation_All_5e619c69 struct {
	ID    int
	Title string
}

func (InForMation_All_5e619c69) TableName() string {
	return "information_all_5e619c69"
}

func main() {
	/*db_user := "jiemin"
	db_pass := "jiemin2017"
	db_host := "127.0.0.1"
	db_port := 3306
	db_name := "jiemin"*/

	db_user := "funbird"
	db_pass := "funbird2017"
	db_host := "10.10.206.205"
	db_port := 3306
	db_name := "funbird_analysis"
	dbConnectBaseStr := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnect := fmt.Sprintf(dbConnectBaseStr, db_user, db_pass, db_host, db_port, db_name)
	fmt.Print(dbConnect)

	DB, _ = gorm.Open("mysql", dbConnect)
	/*DB, err = gorm.Open("mysql", dbConnect)
	if err != nil {
		fmt.Println("Connection to MySQL database Failed！")
	}
	defer DB.Close()*/

	information_all := []*InForMation_All_5e619c69{}

	//DB.Table("information_all_5e619c69").Where("title like  CONCAT('%',CHAR(10),'%')").Find(&information_all)
	DB.Where("title like  CONCAT('%',CHAR(10),'%')").Find(&information_all)
	for _, in_all := range information_all {
		//informationAllmap := map[string]string{}
		//json.Unmarshal([]byte(in_all.Title), &informationAllmap)
		in_all.Title = strings.Split(in_all.Title, "\n")[0]
		DB.Table("information_all_5e619c69").Where("id=?", in_all.ID).Update(&in_all)
		fmt.Printf("Update_id:%d\tUpdate_title:%s\n", in_all.ID, in_all.Title)
	}
}
