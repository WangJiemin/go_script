package models

import (
	"fmt"
	"golang_Learn/dataReport/gettime"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db124 *gorm.DB
	db186 *gorm.DB
	err   error
)

func Initialize(config map[string]string) {
	var dbHost, dbUser, dbPwd, dbName, dbPort string
	dbConnentBaseStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	dbHost = config["db_host_s"]
	dbUser = config["db_user_s"]
	dbPwd = config["db_pwd_s"]
	dbName = config["db_name_s"]
	dbPort = config["db_port_s"]

	dbConnent_s := fmt.Sprintf(dbConnentBaseStr, dbUser, dbPwd, dbHost, dbPort, dbName)
	fmt.Print(dbConnent_s, "\n")
	db124, err = gorm.Open("mysql", dbConnent_s)
	if err != nil {
		log.Fatalln(err)
	}

	db124.LogMode(true)

	dbHost = config["db_host_d"]
	dbUser = config["db_user_d"]
	dbPwd = config["db_pwd_d"]
	dbName = config["db_name_d"]
	dbPort = config["db_port_d"]

	dbConnent_d := fmt.Sprintf(dbConnentBaseStr, dbUser, dbPwd, dbHost, dbPort, dbName)
	fmt.Print(dbConnent_d, "\n")
	db186, err = gorm.Open("mysql", dbConnent_d)
	if err != nil {
		log.Fatalln(err)
	}

	db186.LogMode(true)
}

func INFO_TEMP(project_id string) {
	db186.Raw("TRUNCATE TABLE funbird_storage.info_temp").Row()
	taskinfos := []*TASKINFOS{}
	db124.Table("funbird.task_infos").Where("project_id = ?", project_id).Scan(&taskinfos)
	for _, values := range taskinfos {
		data := INFOTEMP{}
		data.TASKID = values.TASK_ID
		data.NAME = values.NAME
		data.ENTRY_LINK = values.ENTRY_LINK
		data.OWNER_NAME = values.OWNER_NAME
		data.CREATE_TIME = values.CREATE_TIME
		db186.Table("funbird_storage.info_temp").Save(&data)
	}
}

func UPDATEINFO(tables_name string) {
	yd := gettime.YESTERDAY()
	t_one, t_two := gettime.TWODAYS()

	info := []*INFO{}
	db186.Raw(fmt.Sprintf("select task_id, COUNT(*) AS count from `%s_%s` GROUP BY task_id", tables_name, yd)).Scan(&info)
	for _, v := range info {
		db186.Table("funbird_storage.info_temp").Where("task_id = ?", v.TASK_ID).Update("count", v.COUNT)
	}
	info_yd := []*INFO{}
	db186.Raw(fmt.Sprintf("select task_id, COUNT(*) AS count from `%s_%s` GROUP BY task_id", tables_name, t_one)).Scan(&info_yd)
	for _, v := range info_yd {
		db186.Table("funbird_storage.info_temp").Where("task_id = ?", v.TASK_ID).Update("count_yesday", v.COUNT)
	}
	info_beyd := []*INFO{}
	db186.Raw(fmt.Sprintf("select task_id, COUNT(*) AS count from `%s_%s` GROUP BY task_id", tables_name, t_two)).Scan(&info_beyd)
	for _, v := range info_beyd {
		db186.Table("funbird_storage.info_temp").Where("task_id = ?", v.TASK_ID).Update("count_beyesday", v.COUNT)
	}
}

func Inforall_dataReport(tables_name, project_id string) map[string]string {

	INFO_TEMP(project_id)
	UPDATEINFO(tables_name)

	infotemp := []*INFOTEMP{}
	Dict := map[string]string{}
	db186.Table("funbird_storage.info_temp").Scan(&infotemp)
	for key, values := range infotemp {
		//fmt.Println(values)
		key = key + 1
		var a, b, c, d, e, f, g, h, _ string = "A", "B", "C", "D", "E", "F", "G", "H", "I"
		StringStr_A := fmt.Sprintf("%s%d", a, key)
		StringStr_B := fmt.Sprintf("%s%d", b, key)
		StringStr_C := fmt.Sprintf("%s%d", c, key)
		StringStr_D := fmt.Sprintf("%s%d", d, key)
		StringStr_E := fmt.Sprintf("%s%d", e, key)
		StringStr_F := fmt.Sprintf("%s%d", f, key)
		StringStr_G := fmt.Sprintf("%s%d", g, key)
		StringStr_H := fmt.Sprintf("%s%d", h, key)
		//StringStr_I := fmt.Sprintf("%s%d", i, key)

		Dict[StringStr_A] = values.TASKID
		Dict[StringStr_B] = values.NAME
		Dict[StringStr_C] = values.ENTRY_LINK
		Dict[StringStr_D] = values.OWNER_NAME
		Dict[StringStr_E] = values.CREATE_TIME
		Dict[StringStr_F] = values.COUNT
		Dict[StringStr_G] = values.COUNT_YESDAY
		Dict[StringStr_H] = values.COUNT_BEYESDAY
		//Dict[StringStr_I] = values.COUNT
	}
	return Dict
}
