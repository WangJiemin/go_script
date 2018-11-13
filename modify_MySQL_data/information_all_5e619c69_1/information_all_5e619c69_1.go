package main

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type InForMation_All_5e619c69 struct {
	ID      int
	Source  string
	Content string
	Media   string
}

func (InForMation_All_5e619c69) TablesName() string {
	return "information_all_5e619c69"
}

func main() {
	db_user := "jiemin"
	db_pass := "jiemin2017"
	db_host := "127.0.0.1"
	db_port := 3306
	db_name := "jiemin"
	/*
		db_user := "funbird"
		db_pass := "funbird2017"
		db_host := "10.10.206.205"
		db_port := 3306
		db_name := "funbird_analysis"
	*/

	dbConnectBaseStr := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnect := fmt.Sprintf(dbConnectBaseStr, db_user, db_pass, db_host, db_port, db_name)
	fmt.Printf(dbConnect, "\n")

	DB, err := gorm.Open("mysql", dbConnect)
	if err != nil {
		fmt.Println("Connection to MySQL database Failed！")
	}
	defer DB.Close()

	information_all := []*InForMation_All_5e619c69{}
	DB.Table("information_all_5e619c69").Where("media = ? AND content like ?", "上海浦东发展银行", "%2017%").Find(&information_all)

	for _, VALUES := range information_all {
		index := 0
		in_all_5e619c69 := strings.Split(VALUES.Content, "\n")[index]
		if len(strings.TrimSpace(in_all_5e619c69)) <= 0 {
			index = index + 1
			in_all_5e619c69 = strings.Split(VALUES.Content, "\n")[index]
		}

		splitResult := strings.Split(in_all_5e619c69, " ")
		if len(splitResult) > 1 {
			VALUES.Source = strings.Split(in_all_5e619c69, " ")[1]
		}
		//fmt.Printf("Coutent:%s\tSource:%s\n", in_all_5e619c69, VALUES.Source)
		//DB.Exec(fmt.Sprintf("update %s set content=REPLACE(%s,'%s','') where id =%d", "information_all_5e619c69", "content", in_all_5e619c69, VALUES.ID))
		VALUES.Content = strings.Replace(VALUES.Content, in_all_5e619c69, "", -1)
		//fmt.Printf("Coutent:%s\n", VALUES.Content)
		DB.Table("information_all_5e619c69").Where("id=?", VALUES.ID).Update(&VALUES)
		fmt.Printf("update_id:%d\tupdate_source:%s\tupdate_content:%s\n", VALUES.ID, VALUES.Source, VALUES.Content)
	}
}
