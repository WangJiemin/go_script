package models

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db124 *gorm.DB
	err   error
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
	fmt.Print(dbConnent, "\n")
	db124, err = gorm.Open("mysql", dbConnent)
	if err != nil {
		log.Fatalln(err)
	}

	db124.LogMode(true)
}

func Inforall_dataReport(tables_name, project_id string) map[string]string {
	//infor := []*INFO{}
	db124.Raw("TRUNCATE TABLE info_temp").Row()
	//db124.Table("information_all_86a18cd9").Select("task_id").Group("task_id").Find(&infor).Count("task_id")
	db124.Raw(fmt.Sprintf("insert into info_temp(task_id,count) select task_id,count(0) from `%s` group by task_id", tables_name)).Row()

	//db124.Raw("select task_id,count(0) from information_all_86a18cd9 group by task_id").Find(&infor)
	/*
		for _, values := range infor {
			fmt.Println(values)
			//data := INFOTEMP{}
			//data.TASKID = values.TASK_ID
			//data.COUNT = values.COUNT
			//db124.Table("info_temp").Save(&data)
		}
	*/
	/*
		rows, err := db124.Raw("select task_id,count(0) from information_all_86a18cd9 group by task_id").Rows()
		if err != nil {
			log.Fatalln(err)
		}
		defer rows.Close()
		for rows.Next() {
			data := INFOTEMP{}
			rows.Scan(&data)
		}
	*/
	taskinfos := []*TASKINFOS{}
	//dic := TaskLogs()
	db124.Raw(fmt.Sprintf("select a.task_id, a.`name`, a.entry_link, a.owner_name, a.create_time, sum(c.page_count) as page_count, sum(c.collected_count) as collected_count, sum(c.insert_count) as insert_count, b.count from task_infos a, info_temp b left join task_logs c on b.task_id=c.task_id where a.task_id=b.task_id and a.`status`!=2 and a.is_start=1 and a.project_id= '%s' group by b.task_id", project_id)).Scan(&taskinfos)
	Dict := map[string]string{}
	for key, values := range taskinfos {
		//fmt.Println(values)
		key = key + 1
		var a, b, c, d, e, f, g, h, i string = "A", "B", "C", "D", "E", "F", "G", "H", "I"
		StringStr_A := fmt.Sprintf("%s%d", a, key)
		StringStr_B := fmt.Sprintf("%s%d", b, key)
		StringStr_C := fmt.Sprintf("%s%d", c, key)
		StringStr_D := fmt.Sprintf("%s%d", d, key)
		StringStr_E := fmt.Sprintf("%s%d", e, key)
		StringStr_F := fmt.Sprintf("%s%d", f, key)
		StringStr_G := fmt.Sprintf("%s%d", g, key)
		StringStr_H := fmt.Sprintf("%s%d", h, key)
		StringStr_I := fmt.Sprintf("%s%d", i, key)

		Dict[StringStr_A] = values.TASK_ID
		Dict[StringStr_B] = values.NAME
		Dict[StringStr_C] = values.ENTRY_LINK
		Dict[StringStr_D] = values.OWNER_NAME
		Dict[StringStr_E] = values.CREATE_TIME
		Dict[StringStr_F] = values.PAGE_COUNT
		Dict[StringStr_G] = values.COLLECTED_COUNT
		Dict[StringStr_H] = values.INSERT_COUNT
		Dict[StringStr_I] = values.COUNT

		//fmt.Printf("F:\t%s\tG:\t%s\tH:\t%s\tI:\t%s\t", Dict[StringStr_F], Dict[StringStr_G], Dict[StringStr_H], Dict[StringStr_I])
		/*
			if dic[values.TASK_ID] != nil {
				Dict[StringStr_G] = dic[values.TASK_ID].COLLECTED_COUNT
				Dict[StringStr_H] = dic[values.TASK_ID].INSERT_COUNT
				Dict[StringStr_I] = dic[values.TASK_ID].PAGE_COUNT
			}
			fmt.Printf("F:\t%s\tG:\t%s\tH:\t%s\tI:\t%s\t", Dict[StringStr_F], Dict[StringStr_G], Dict[StringStr_H], Dict[StringStr_I])
		*/
	}

	return Dict
}

func InfoallChayi(tables_name, project_id string) map[string]string {
	taskinfos := []*TASKINFOS{}
	db124.Raw(fmt.Sprintf("select task_id, name, entry_link, owner_name,create_time from task_infos a where task_id not in (select distinct(task_id) from `%s`) and a.status!=2 and a.is_start=1 and a.project_id='%s'", tables_name, project_id)).Scan(&taskinfos)
	Dict := map[string]string{}
	for key, values := range taskinfos {
		//fmt.Println(values)
		key = key + 1
		var a, b, c, d, e, f, g, h, i string = "A", "B", "C", "D", "E", "F", "G", "H", "I"
		StringStr_A := fmt.Sprintf("%s%d", a, key)
		StringStr_B := fmt.Sprintf("%s%d", b, key)
		StringStr_C := fmt.Sprintf("%s%d", c, key)
		StringStr_D := fmt.Sprintf("%s%d", d, key)
		StringStr_E := fmt.Sprintf("%s%d", e, key)
		StringStr_F := fmt.Sprintf("%s%d", f, key)
		StringStr_G := fmt.Sprintf("%s%d", g, key)
		StringStr_H := fmt.Sprintf("%s%d", h, key)
		StringStr_I := fmt.Sprintf("%s%d", i, key)

		Dict[StringStr_A] = values.TASK_ID
		Dict[StringStr_B] = values.NAME
		Dict[StringStr_C] = values.ENTRY_LINK
		Dict[StringStr_D] = values.OWNER_NAME
		Dict[StringStr_E] = values.CREATE_TIME
		Dict[StringStr_F] = "0"
		Dict[StringStr_G] = "0"
		Dict[StringStr_H] = "0"
		Dict[StringStr_I] = "0"
		//fmt.Println(dict[StringStr_B])
	}

	return Dict
}

/*
func TaskLogs() map[string]*TASKLOGS {
	tasklogs := []*TASKLOGS{}
	db124.Table("task_logs").Select("task_id,SUM(collected_count),SUM(insert_count),SUM(page_count)").Group("task_id").Scan(&tasklogs)
	Dict := map[string]*TASKLOGS{}
	for _, v := range tasklogs {
		Dict[v.TASK_ID] = v
	}
	return Dict
}

func TaskLogs() map[string]*TASKINFOS {
	taskinfos := []*TASKINFOS{}
	db124.Table("task_logs").Select("task_id,SUM(collected_count),SUM(insert_count),SUM(page_count)").Group("task_id").Scan(&taskinfos)
	Dict := map[string]*TASKINFOS{}
	for _, v := range taskinfos {
		Dict[v.TASK_ID] = v
	}
	return Dict
}
*/
