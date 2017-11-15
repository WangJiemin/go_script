package models

import (
	cexcel "excel_informall/C_excel"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Initiailize(config map[string]string) {
	var dbUser, dbPwd, dbHost, dbPort, dbName string
	dbConnentBaseStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	dbUser = config["db_user"]
	dbPwd = config["db_pwd"]
	dbHost = config["db_host"]
	dbPort = config["db_port"]
	dbName = config["db_name"]

	dbConnent := fmt.Sprintf(dbConnentBaseStr, dbUser, dbPwd, dbHost, dbPort, dbName)
	fmt.Print(dbConnent, "\n")
	db, err = gorm.Open("mysql", dbConnent)
	if err != nil {
		log.Fatalln(err)
	}

	db.LogMode(true)
}

func FUNBIRDTABLES(tablesname, project_id string) {
	var tasknum int
	taskidnum := []*TASKCOUNT{}
	db.Table("task_infos").Where("project_id = ?", project_id).Select("count(*) as count").Scan(&taskidnum)
	for _, v := range taskidnum {
		//fmt.Println(v.COUNT)
		tasknum = v.COUNT
	}

	taskinfos := []*TASKINFOS{}
	//dict := map[string]string{}

	for i := 0; i <= tasknum; i = i + 4 {
		//db.Table("funbird.task_infos").Where("project_id = ?", project_id).Select("task_id, `name`").Limit(fmt.Sprintf("%d,%d", i, 4)).Scan(&taskinfos)
		db.Raw(fmt.Sprintf("select task_id,`name` from task_infos where project_id = '%s' limit %d, %d", project_id, i, 4)).Scan(&taskinfos)
		for _, v := range taskinfos {
			go FUNBIRTASKID(tablesname, v.TASKID)
		}
	}
}

func FUNBIRTASKID(tablesname, taskID string) {
	dict := map[string]string{}
	infoall := []*INFORALL{}
	db.Table(fmt.Sprintf("%s", tablesname)).Where("task_id = ?", taskID).Scan(&infoall)
	key := 0
	for _, values := range infoall {
		if values.Content != "" {
			key = key + 1
			var a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u string = "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U"
			StringStr_A := fmt.Sprintf("%s%d", a, key)
			StringStr_B := fmt.Sprintf("%s%d", b, key)
			StringStr_C := fmt.Sprintf("%s%d", c, key)
			StringStr_D := fmt.Sprintf("%s%d", d, key)
			StringStr_E := fmt.Sprintf("%s%d", e, key)
			StringStr_F := fmt.Sprintf("%s%d", f, key)
			StringStr_G := fmt.Sprintf("%s%d", g, key)
			StringStr_H := fmt.Sprintf("%s%d", h, key)
			StringStr_I := fmt.Sprintf("%s%d", i, key)
			StringStr_J := fmt.Sprintf("%s%d", j, key)
			StringStr_K := fmt.Sprintf("%s%d", k, key)
			StringStr_L := fmt.Sprintf("%s%d", l, key)
			StringStr_M := fmt.Sprintf("%s%d", m, key)
			StringStr_N := fmt.Sprintf("%s%d", n, key)
			StringStr_O := fmt.Sprintf("%s%d", o, key)
			StringStr_P := fmt.Sprintf("%s%d", p, key)
			StringStr_Q := fmt.Sprintf("%s%d", q, key)
			StringStr_R := fmt.Sprintf("%s%d", r, key)
			StringStr_S := fmt.Sprintf("%s%d", s, key)
			StringStr_T := fmt.Sprintf("%s%d", t, key)
			StringStr_U := fmt.Sprintf("%s%d", u, key)

			dict[StringStr_A] = values.Url
			dict[StringStr_B] = values.Title
			dict[StringStr_C] = values.Author
			dict[StringStr_D] = values.Source
			dict[StringStr_E] = values.Pub_time
			dict[StringStr_F] = values.Content
			dict[StringStr_G] = values.Media_type
			dict[StringStr_H] = values.Original_title
			dict[StringStr_I] = values.Editor
			dict[StringStr_J] = values.Reporter
			dict[StringStr_K] = values.Contents
			dict[StringStr_L] = values.Reading_times
			dict[StringStr_M] = values.Abstract_data
			dict[StringStr_N] = values.Media
			dict[StringStr_O] = values.Media_channel
			dict[StringStr_P] = values.Location
			dict[StringStr_Q] = values.Location_path
			dict[StringStr_R] = values.Collection_tool
			dict[StringStr_S] = values.User
			dict[StringStr_T] = values.Site_url
			dict[StringStr_U] = values.Get_time

			//fmt.Println(dict)
			os.Mkdir(fmt.Sprintf("/tmp/funbird/%s", tablesname), 0777)
			cexcel.CreateExcel(fmt.Sprintf("/tmp/funbird/%s/%s.xlsx", tablesname, taskID), dict)
		}
	}
}
