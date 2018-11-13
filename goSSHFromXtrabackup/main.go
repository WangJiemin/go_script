package main

import (
	"bytes"
	"log"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/WangJiemin/gocomm"

	"learngo/goSSHFromXtrabackup/gossh"
)

var (
	stdOut, stdErr                                                             bytes.Buffer
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
	sshExceSession, err := gossh.SSHConnect(sshUser, sshPass, sshHost, sshPort)
	checkErr(err)
	defer sshExceSession.Close()

	sshExceSession.Stdout = &stdOut
	sshExceSession.Stderr = &stdErr

	innobackupexCommand := fmt.Sprintf("%s --defaults-file=%s --user=%s --password=%s --host=%s --port=%s"+
		" --stream=xbstream --encrypt-threads=6  --compress --compress-threads=8 --tmpdir=%s"+
		" %s", xtrabackupCommand, xtrabackupConfig, xtrabackupUser, xtrabackupPass, xtrabackupHost, xtrabackupPort, xtrabackupTmpDir, xtrabackupBackupDir)
	fmt.Println(innobackupexCommand)
	sshExceSession.Run(innobackupexCommand)

	str, err := strconv.Atoi(strings.Replace(stdOut.String(), "\n", "", -1))
	checkErr(err)
	//str := strings.Replace(stdOut.String(), "\n", "", -1)
	fmt.Printf("%s, %s\n", str, stdErr.String())

	wg.Done()
}

func main() {
	conf := gocomm.ReadConfig("./conf/app.cnf")
	initialize(conf)

	runtime.GOMAXPROCS(2)

	wg.Add(1)
	go RunSSH()
	wg.Wait()
}
