package main

import (
	db "golang_Learn/synckafka/database"
	ka "golang_Learn/synckafka/kafka"

	"github.com/WangJiemin/gocomm"
)

func main() {
	config := gocomm.ReadConfig("./conf/app.cnf")
	db.Initialize(config)
	ka.SyncToKafka(config, "information_two_1e8e0d58")
}
