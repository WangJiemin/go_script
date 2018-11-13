package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadConfig(filepath string) map[string]string {
	res := map[string]string{}
	file, err := os.Open(filepath)
	if err != nil {
		return res
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				return res
			}
			if len(line) == 0 {
				break
			}
		}

		if len(line) == 0 || line == "#" || line == "\r\n" {
			//break
			continue
		}

		if line[0] == '/' {
			//fmt.Println("line[0] =", line[0])
			continue
		}
		//fmt.Println(line)
		//fmt.Println("len(line) =", len(line))
		i := strings.IndexAny(line, "=")
		//fmt.Println("i = ", i)
		value := strings.TrimSpace(line[i+1 : len(line)])
		//fmt.Println("value =", value)
		res[strings.TrimSpace(line[0:i])] = value
		//fmt.Printf("res[strings.TrimSpace(line[0:%d])] = %s\n", i, value)
	}
	return res
}

func main() {
	var User, Host, Port, Touser string
	conf := ReadConfig("F:\\Coding\\Golang\\src\\learngo\\goReadConfig\\app.cnf")
	User = conf["user"]
	Port = conf["port"]
	Host = conf["host"]
	Touser = conf["touser"]
	fmt.Printf("User is %s\nPort is %s\nHost is %s\nTouser is %s\n", User, Port, Host, Touser)
}
