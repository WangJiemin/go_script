package osexce

import (
	"os/exec"
	"log"
)

func ExecCommand(commands string) (string) {
	out, err := exec.Command("bash", "-c", commands).Output()
	if err != nil {
		log.Printf("ERR: %s", err)
	}
	return string(out)
}
