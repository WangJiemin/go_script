package main

import (
	"learngo/goInception/models"
)

func main() {
	models.Initialize("F:\\Coding\\Golang\\src\\learngo\\goInception\\conf\\conf.yaml")
	models.DBAction(models.MySQLRemoteDBName)
}
