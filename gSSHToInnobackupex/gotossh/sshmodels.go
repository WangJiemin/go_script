package gotossh

import "golang.org/x/crypto/ssh"

type Cli struct {
	IP       string //IP地址
	Username string //用户名
	Password string //密码
	//Port       int          //端口号
	Port       string       //端口号
	client     *ssh.Client  //ssh客户端
	session    *ssh.Session //ssh session
	LastResult string       //最近一次Run的结果
}
