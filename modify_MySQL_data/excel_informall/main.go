package main

import (
	"golang_Learn/modify_MySQL_data/excel_informall/compress"
	db "golang_Learn/modify_MySQL_data/excel_informall/models"
	"log"
	"os"

	"github.com/WangJiemin/gocomm"
)

/*
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		os.MkdirAll()
		err := os.Mkdir(path, 0777)
		if err != nil {
			log.Fatalln("Directory creation failed\n")
		}
	}
	return false, err
}
*/

func main() {
	config := gocomm.ReadConfig("./conf/app.cnf")
	db.Initiailize(config)

	//PathExists("/tmp/funbird")
	os.Mkdir("/tmp/funbird", 0777)

	db.FUNBIRDTABLES("information_all_86a18cd9", "86a18cd9-6a06-4794-b7eb-6ce8d8c57bb0")

	f1, err := os.Open("/tmp/funbird")
	if err != nil {
		log.Fatalln(err)
	}
	defer f1.Close()
	var files = []*os.File{f1}
	dest := "/tmp/inforall.zip"
	compress.ComZIP(files, dest)
	//os.RemoveAll("/tmp/funbird")
}
