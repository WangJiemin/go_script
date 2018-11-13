package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func ExceFulshLogs(user, password, host, dbname, port string) {
	dbConnect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	db, err := sql.Open("mysql", dbConnect)
	if err != nil {
		fmt.Println("ERR: ", err)
	}
	//err = db.Ping()
	//if err != nil {
	//	log.Fatal(err)
	//}
	_, err = db.Exec("FLUSH LOGS")
	if err != nil {
		fmt.Println("ERR: ", err)
	}
	db.Close()
}
