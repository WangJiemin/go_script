package main

import (
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/WangJiemin/gocomm"
)

// 需要同步的表的结构,每次同步表结构就需要改成需要同步的表结构
type SyncTablesName struct {
	ID                string `gorm:"cloumn:id"`
	MEDIA_TYPE        string `gorm:"cloumn:media_type"`
	URL               string `gorm:"cloumn:url"`
	TITLE             string `gorm:"cloumn:title"`
	ORIGINAL_TITLE    string `gorm:"cloumn:original_title"`
	ABSTRACT_DATA     string `gorm:"cloumn:abstract_data"`
	CONTENT           string `gorm:"cloumn:content"`
	MEDIA             string `gorm:"cloumn:media"`
	MEDIA_CHANNEL     string `gorm:"cloumn:media_channel"`
	MEDIA_CHANNEL_SND string `gorm:"cloumn:media_channel_snd"`
	MEDIA_CHANNEL_TRD string `gorm:"cloumn:media_channel_trd"`
	LOCATION          string `gorm:"cloumn:location"`
	LOCATIONPATH      string `gorm:"cloumn:locationPath"`
	SOURCE            string `gorm:"cloumn:source"`
}

// 定义同步表的记录的最后的id数值。用于增量数据的同步
type SYNCTABLESDATA struct {
	ID    uint   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id"`
	NAME  string `gorm:"type:varchar(200);unique_index:UN_tablename;not null;column:tablename"`
	NUMID int    `gorm:"not null;column:numid"`
}

// 定义全局参数
var (
	dbS *gorm.DB
	dbD *gorm.DB
	err error
	// 源数据库的参数
	dbHostS, dbUserS, dbPwdS, dbNameS, dbTablesS, dbPortS string
	// 目的数据库的参数
	dbHostD, dbUserD, dbPwdD, dbNameD, dbTablesD, dbPortD string
)

// 初始化数据库参数
func Init(config map[string]string) {
	// 连接串
	dbConnentBaseStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	// 源数据库获取参数的values
	dbHostS = config["db_host_s"]
	dbUserS = config["db_user_s"]
	dbPwdS = config["db_pwd_s"]
	dbNameS = config["db_name_s"]
	dbTablesS = config["db_table_s"]
	dbPortS = config["db_port_s"]

	dbConnentS := fmt.Sprintf(dbConnentBaseStr, dbUserS, dbPwdS, dbHostS, dbPortS, dbNameS)
	fmt.Print(dbConnentS, "\n")
	dbS, err = gorm.Open("mysql", dbConnentS)
	if err != nil {
		log.Fatalln(err)
	}
	//defer dbS.Close()
	dbS.LogMode(true)

	// 目标数据库获取参数的values
	dbHostD = config["db_host_d"]
	dbUserD = config["db_user_d"]
	dbPwdD = config["db_pwd_d"]
	dbNameD = config["db_name_d"]
	dbTablesD = config["db_table_d"]
	dbPortD = config["db_port_d"]

	dbConnentD := fmt.Sprintf(dbConnentBaseStr, dbUserD, dbPwdD, dbHostD, dbPortD, dbNameD)
	fmt.Print(dbConnentD, "\n")
	dbD, err = gorm.Open("mysql", dbConnentD)
	if err != nil {
		log.Fatalln(err)
	}
	//defer dbD.Close()
	dbD.LogMode(true)
}

// 程序入口
func main() {
	// 读取配置文件app.cof
	conf := gocomm.ReadConfig("./app.cnf")
	// 初始化数据库连接参数
	Init(conf)

	// 判断同步记录ID表存在与否，不存在，则创建
	if (!dbS.HasTable(&SYNCTABLESDATA{})) {
		dbS.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&SYNCTABLESDATA{})
	}
	//initData:=&SYNCTABLESDATA{NAME:dbTablesS,NUMID:0}
	//dbS.Create(initData)

	var numid int
	synctablesstruct := []*SYNCTABLESDATA{}

	// 判断同步记录表有没有要同步表名在。如果在，就忽略。不在则插入要同步的表名到同步记录表中
	//if err = dbS.Table("synctablesdata").Where("tablename = ?", dbTablesS).First(&synctablesstruct).Error; err != nil {
	//}

	//dbS.Table("synctablesdata").Where("tablename = ?", dbTablesS).First(&synctablesstruct).RecordNotFound()
	//if dbS.Model(&synctablesstruct).Related(&synctablesstruct).RecordNotFound() {
	//}

	//if err = dbS.Table("synctablesdata").Where("tablename = ?", dbTablesS).First(&synctablesstruct).Error; gorm.IsRecordNotFoundError(err) {
	//	initData := &SYNCTABLESDATA{NAME: dbTablesS, NUMID: 0}
	//	dbS.Create(initData)
	//}

	//err = dbS.Table("synctablesdata").Where("tablename = ?", dbTablesS).First(&synctablesstruct).Error
	//if err != gorm.ErrRecordNotFound {
	//	initData := &SYNCTABLESDATA{NAME: dbTablesS, NUMID: 0}
	//	dbS.Create(initData)
	//}
	var tableField string
	row := dbS.Table("synctablesdata").Where("tablename = ?", dbTablesS).Select("tablename").Row()
	row.Scan(&tableField)
	if len(tableField) == 0 {
		initData := &SYNCTABLESDATA{NAME: dbTablesS, NUMID: 0}
		dbS.Create(initData)
	}

	// 读取同步记录表的记录的最后的id数值
	// SELECT numid FROM synctablesdata WHERE tablename = 'XXXXX'
	dbS.Table("synctablesdata").Where("tablename = ?", dbTablesS).Select("numid").Find(&synctablesstruct)
	for _, v := range synctablesstruct {
		numid = v.NUMID
	}
	//fmt.Printf("update table name is:\t%s\tAUTO INCREMENT number is:\t%d\n", "", numid)

	// 读取需要同步表的数据
	synctablesname := []*SyncTablesName{}
	// 写上同步的表名字
	dbS.Table(dbTablesS).Where("id > ?", numid).Find(&synctablesname)
	for _, v := range synctablesname {
		//data := SyncTablesName{}
		// 同步目的表的名字
		dbD.Table(dbTablesD).Create(&v)
		//dbD.Table(dbTablesD).Save(&v)
		dbS.Table("synctablesdata").Where("tablename = ?", dbTablesS).Update("numid", v.ID)
	}
}
