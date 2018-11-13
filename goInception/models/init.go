package models

import (
	"fmt"
	"io/ioutil"
	"log"

	//"database/sql"
	"github.com/jinzhu/gorm"
	//"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"

	"gopkg.in/yaml.v2"
)

var (
	MysqlUser, MysqlPass, MysqlHost, MysqlPort, MysqlDBName, MysqlDBTableName                                     string
	MySQLRemoteUser, MySQLRemotePass, MySQLRemoteHost, MySQLRemotePort, MySQLRemoteDBName, MySQLRemoteDBTableName string
	InceptionValues                                                                                               string = "enable-query-print"
	//SqlFormat                                                                                                     string
	Db *gorm.DB
	//Db  *sql.DB
	//Db  *xorm.Engine
	err error
)

func Initialize(conf string) {
	cnf := UserConfig{}
	yamlFile, err := ioutil.ReadFile(conf)
	CheckErr(err)

	err = yaml.Unmarshal(yamlFile, &cnf)
	CheckErr(err)

	MysqlUser = cnf.Inc.User
	MysqlPass = cnf.Inc.Password
	MysqlHost = cnf.Inc.Host
	MysqlPort = cnf.Inc.Port
	MysqlDBName = cnf.Inc.DB

	MySQLRemoteUser = cnf.Remote.User
	MySQLRemotePass = cnf.Remote.Password
	MySQLRemoteHost = cnf.Remote.Host
	MySQLRemotePort = cnf.Remote.Port
	MySQLRemoteDBName = cnf.Remote.DB
	MySQLRemoteDBTableName = cnf.Remote.DBTable
	//SqlFormat = cnf.Remote.SqlFormat

	//dbConnentBaseStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnentBaseStr := "%s:%s@tcp(%s:%s)/%s"
	dbConnent := fmt.Sprintf(dbConnentBaseStr, MysqlUser, MysqlPass, MysqlHost, MysqlPort, MysqlDBName)
	//fmt.Print(dbConnent, "\n")
	Db, err = gorm.Open("mysql", dbConnent)
	//Db, err = sql.Open("mysql", dbConnent)
	//Db, err = xorm.NewEngine("mysql", dbConnent)
	if err != nil {
		log.Fatalln(err)
	}
	//defer Db.Close()
	//Db.LogMode(true)
}

func CheckErr(errinfo error) {
	if errinfo != nil {
		fmt.Println(errinfo.Error())
	}
}
