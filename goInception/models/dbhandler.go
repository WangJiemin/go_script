package models

import (
	"fmt"
)

func DBhandler(db, sqlformat string) {
	sql := fmt.Sprintf("/*--user=%s;--password=%s;--host=%s;--execute=1;--port=%s;*/"+
		"inception_magic_start;"+
		"use %s;"+
		"%s;"+
		"inception_magic_commit;", MySQLRemoteUser, MySQLRemotePass, MySQLRemoteHost, MySQLRemotePort, db, sqlformat)

	fmt.Println(sql)
	incinformationinfo := []*IncInformationInfo{}
	Db.Exec(sql).Scan(incinformationinfo)
	for _,v :=range incinformationinfo{
		fmt.Println(v)
	}

	//incinformationinfo := []*IncInformationInfo{}
	//Db.Raw(sql).Scan(&incinformationinfo)
	//for _, v := range incinformationinfo {
	//	fmt.Println(v)
	//}
}
