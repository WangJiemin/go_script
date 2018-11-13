package main

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	check := func(err error, msg string) {
		if err != nil {
			log.Fatal("%s error: %v", msg, err)
		}
	}

	client, err := ssh.Dial("tcp", "192.168.1.254", &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{ssh.Password("123456")},
	})
	check(err, "Dial connection error")

	session, err := client.NewSession()
	check(err, "New Session")
	defer session.Close()

	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(fd, oldState)

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	termWidth, termHeight, err := terminal.GetSize(fd)
	if err != nil {
		panic(err)
	}

	models := ssh.TerminalModes{
		ssh.ECHO:          0,     // enable echoing
		ssh.TTY_OP_OSPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_ISPEED: 14400, // output speed = 14.4kbaud
	}
	//err = session.RequestPty("xterm", 25, 100, models)
	err = session.RequestPty("xterm-256color", termHeight, termWidth, models)
	check(err, "Request pty")

	err = session.Shell()
	check(err, "Start Shell")

	err = session.Wait()
	check(err, "return")
}
