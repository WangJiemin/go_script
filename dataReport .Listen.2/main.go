package main

import (
	//"fmt"
	"funbird-dataReport/compress"
	"funbird-dataReport/execl"
	db "funbird-dataReport/models"
	mail "funbird-dataReport/sendmail"
	"log"
	"os"

	"github.com/WangJiemin/gocomm"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0777)
		if err != nil {
			log.Fatalln("Directory creation failed\n")
		}
	}
	return false, err
}

func main() {
	config := gocomm.ReadConfig("./conf/app.cnf")
	db.Initialize(config)

	PathExists("/tmp/dataReport")

	dict_86 := db.Inforall_dataReport("information_all_86a18cd9", "86a18cd9-6a06-4794-b7eb-6ce8d8c57bb0")
	execl.Execl_install("/tmp/dataReport/86表.xlsx", dict_86)
	//dict_chayi86 := db.InfoallChayi("information_all_86a18cd9", "86a18cd9-6a06-4794-b7eb-6ce8d8c57bb0")
	//execl.Execl_install("/tmp/dataReport/86差异表.xlsx", dict_chayi86)

	//compress.ComTarGz("/tmp/dataReport", "/tmp/dataReport/dataReport.tar.gz")

	/*
		f1, err := os.Open("/tmp/dataReport")
		if err != nil {
			log.Fatalln(err)
		}
		defer f1.Close()
		var files = []*os.File{f1}
	*/
	f1, err := os.Open("/tmp/dataReport/86表.xlsx")
	if err != nil {
		log.Fatalln(err)
	}
	defer f1.Close()

	var files = []*os.File{f1}
	dest := "/tmp/dataReport.zip"
	compress.ComZIP(files, dest)

	mail.SendGoMail(config, "数据报表", "/tmp/dataReport.zip")

	os.RemoveAll("/tmp/dataReport")
	os.Remove("/tmp/dataReport.zip")

}
