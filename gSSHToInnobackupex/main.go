package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"

	"gSSHToInnobackupex/gotossh"
	"github.com/WangJiemin/gocomm"
)

var (
	sshUser, sshPass, sshHost, sshPort                                         string
	xtrabackupUser, xtrabackupPass, xtrabackupHost, xtrabackupPort             string
	xtrabackupConfig, xtrabackupTmpDir, xtrabackupBackupDir, xtrabackupCommand string
	wg                                                                         sync.WaitGroup
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func initialize(config map[string]string) {
	sshUser = config["sshuser"]
	sshPass = config["sshpass"]
	sshHost = config["sshhost"]
	sshPort = config["sshport"]

	xtrabackupUser = config["xtrabackupuser"]
	xtrabackupPass = config["xtrabackuppass"]
	xtrabackupHost = config["xtrabackuphost"]
	xtrabackupPort = config["xtrabackupport"]
	xtrabackupConfig = config["xtrabackupconfig"]
	xtrabackupTmpDir = config["xtrabackuptmpdir"]
	xtrabackupBackupDir = config["xtrabackupbackupdir"]
	xtrabackupCommand = config["xtrabackupcommand"]
}

func RunSSH() {
	sshExceSession := gotossh.New(sshHost, sshUser, sshPass, sshPort)
	defer sshExceSession.Close()
	innobackupexCommand := fmt.Sprintf("nohup %s --defaults-file=%s --user=%s --password=%s --host=%s --port=%s"+
		" --stream=xbstream --encrypt-threads=6  --compress --compress-threads=8 --tmpdir=%s"+
		" %s &", xtrabackupCommand, xtrabackupConfig, xtrabackupUser, xtrabackupPass, xtrabackupHost, xtrabackupPort, xtrabackupTmpDir, xtrabackupBackupDir)
	fmt.Println(innobackupexCommand)
	exceErr := sshExceSession.Start(innobackupexCommand)
	checkErr(exceErr)
	wg.Done()
}

func RunSSHTerminal() {
	sshExceSession := gotossh.New(sshHost, sshUser, sshPass, sshPort)
	sshExceSession.RunTerminal("top", os.Stdout, os.Stdin)
	wg.Done()
}

func main() {
	conf := gocomm.ReadConfig("./conf/app.cnf")
	initialize(conf)
	runtime.GOMAXPROCS()

	//wg.Add(1)
	//go RunSSH()
	//wg.Wait()

	wg.Add(1)
	go RunSSHTerminal()
	wg.Wait()
}
