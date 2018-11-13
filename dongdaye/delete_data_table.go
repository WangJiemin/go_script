package main

import (
	"fmt"
	"log"
	"sync"
	//"time"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/WangJiemin/gocomm"

)

type TMPLIULIANG struct {
	ID             int   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id"`
	WorkflowId     int64 `gorm:"index:idx_WorkflowID"`
	WorkflowStepId int64 `gorm:"index:idx_WorkflowStepID"`
}

var (
	db  *gorm.DB
	err error
	// 源数据库的参数
	dbHost, dbUser, dbPwd, dbName, dbPort, MinID, MaxID string
	wg sync.WaitGroup
)

// 初始化数据库参数
func Init(config map[string]string) {
	// 连接串
	dbConnentBaseStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	// 源数据库获取参数的values
	dbHost = config["db_host"]
	dbUser = config["db_user"]
	dbPwd = config["db_pwd"]
	dbName = config["db_name"]
	dbPort = config["db_port"]
	MinID = config["minid"]
	MaxID = config["maxid"]

	dbConnent := fmt.Sprintf(dbConnentBaseStr, dbUser, dbPwd, dbHost, dbPort, dbName)
	fmt.Print(dbConnent, "\n")
	db, err = gorm.Open("mysql", dbConnent)
	if err != nil {
		log.Fatalln(err)
	}
	//defer dbS.Close()
	db.LogMode(true)
}

func main() {
	conf := gocomm.ReadConfig("./app.cnf")
	Init(conf)

	if (!db.HasTable(&TMPLIULIANG{})) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&TMPLIULIANG{})
	}

	wg.Add(1)
	go SelectTmpLiuliangsTable(MinID, MaxID)
	wg.Wait()

	wg.Add(2)
	go deleteExecution()
	go deleteWorkflowStep()
	wg.Wait()

	fmt.Println("Operation completed!")
	db.Raw("set GLOBAL foreign_key_checks=1").Row()
}

func SelectTmpLiuliangsTable(minid, maxid string) {

	//value := make(chan int64, 10000)

	tmpliuliang := []*TMPLIULIANG{}
	tmpliuliang1 := []*TMPLIULIANG{}
	tmpliuliang2 := []*TMPLIULIANG{}

	//db.Raw("TRUNCATE TABLE tmpliuliangs").Row()
	//db.Raw("TRUNCATE TABLE tmpliuliangs")
	db.Exec("TRUNCATE TABLE tmpliuliangs")

	db.Table("execution").Where("id > ? AND id < ?", MinID, MaxID).Select("workflow_id").Find(&tmpliuliang)
	for _, v := range tmpliuliang {
		//value <- v.WorkflowId
		//val := <-value
		//db.Table("tmpliuliangs").Create(&val)
		db.Table("tmpliuliangs").Save(&v)
	}

	db.Table("tmpliuliangs").Select("workflow_id").Find(&tmpliuliang1)
	for _, v := range tmpliuliang1 {
		db.Table("workflow_workflow_step").Where("workflow_commands_id = ?", v.WorkflowId).Select("workflow_step_id").Find(&tmpliuliang2)
		for _, vv := range tmpliuliang2 {
			//value <- vv.WorkflowStepId
			//val := <-value
			//db.Table("tmpliuliangs").Save(&val)
			db.Table("tmpliuliangs").Save(&vv)
		}
	}

	/*
	row := db.Table("tmpliuliangs").Select("workflow_id, workflow_step_id").Row()
	row.Scan(&workflowid, &workflowstepid)
	fmt.Println(workflowid, workflowstepid)
	if len(workflowid) == 0 {
		fmt.Println("Exce SELECT execution to tmpliuliangs")
		time.Sleep(time.Second * 10)
		db.Table("execution").Where("id > ? AND id < ?", minid, maxid).Select("workflow_id").Find(&tmpliuliang)
		for _, v := range tmpliuliang {
			//value <- v.WorkflowId
			//val := <-value
			//db.Table("tmpliuliangs").Create(&val)
			db.Table("tmpliuliangs").Create(&v)
		}
	} else if len(workflowstepid) == 0 {
		fmt.Println("Exce SELECT tmpliuliangs to workflow_workflow_step")
		time.Sleep(time.Second * 10)
		db.Table("tmpliuliangs").Select("workflow_id").Find(&tmpliuliang1)
		for _, v := range tmpliuliang1 {
			db.Table("workflow_workflow_step").Where("workflow_commands_id = ?", v.WorkflowId).Select("workflow_step_id").Find(&tmpliuliang2)
			for _, vv := range tmpliuliang2 {
				//value <- vv.WorkflowStepId
				//val := <-value
				//db.Table("tmpliuliangs").Save(&val)
				db.Table("tmpliuliangs").Save(&vv)
			}
		}
	} else if len(workflowid) != 0 && len(workflowstepid) != 0 {
		//db.Raw("TRUNCATE TABLE tmpliuliangs").Row()
		//db.Raw("TRUNCATE TABLE tmpliuliangs")
		db.Exec("TRUNCATE TABLE tmpliuliangs")
	}
*/

	wg.Done()
}

func deleteExecution() {
	tmpliuliang := []*TMPLIULIANG{}
	db.Table("tmpliuliangs").Select("workflow_id").Find(&tmpliuliang)
	db.Raw("set GLOBAL foreign_key_checks=0").Row()
	for _, v := range tmpliuliang {
		db.Raw("DELETE FROM execution WHERE workflow_id = ?", v.WorkflowId).Row()
		db.Raw("DELETE FROM workflow WHERE id = ?", v.WorkflowId).Row()
		db.Raw("DELETE FROM workflow_workflow_step WHERE workflow_commands_id = ?", v.WorkflowId).Row()
	}

	wg.Done()
}

/*
func deleteWorkflow() {
	tmpliuliang := []*TMPLIULIANG{}
	db.Table("tmpliuliangs").Select("workflow_id").Find(&tmpliuliang)
	for _, v := range tmpliuliang {
		db.Raw("DELETE FROM workflow WHERE id = ?",v.WorkflowId).Row()
	}

    wg.Done()
}
*/

func deleteWorkflowStep() {
	tmpliuliang := []*TMPLIULIANG{}
	db.Raw("set GLOBAL foreign_key_checks=0").Row()
	db.Table("tmpliuliangs").Select("workflow_step_id").Find(&tmpliuliang)
	for _, v := range tmpliuliang {
		db.Raw("DELETE FROM workflow_step WHERE id = ?", v.WorkflowStepId).Row()
	}

	wg.Done()
}

/*
func deleteWorkflowWorkflowStep() {
	tmpliuliang := []*TMPLIULIANG{}
	db.Table("tmpliuliangs").Select("workflow_id").Find(&tmpliuliang)
	for _, v := range tmpliuliang {
		db.Raw("DELETE FROM workflow_workflow_step WHERE workflow_commands_id = ?",v.WorkflowId).Row()
	}

	wg.Done()
}
*/
