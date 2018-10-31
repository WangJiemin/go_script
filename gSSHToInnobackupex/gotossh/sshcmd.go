package gotossh

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

//创建命令行对象
//@param ip IP地址
//@param username 用户名
//@param password 密码
//@param port 端口号,默认22
func New(ip, username, password, port string) (*Cli) {
	cli := new(Cli)
	cli.IP, cli.Username, cli.Password = ip, username, password
	l := len(port)
	if l <= 0 {
		cli.Port = "22"
	} else {
		cli.Port = port
	}
	if err := cli.sshConnect(); err != nil {
		log.Fatal(err)
		return nil
	}
	return cli
}

// 创建 ssh 连接
func (cli *Cli) sshConnect() (error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		err          error
	)

	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(cli.Password))

	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) (error) {
		return nil
	}

	clientConfig = &ssh.ClientConfig{
		User:            cli.Username,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// connect to ssh
	//addr = fmt.Sprintf("%s:%d", cli.IP, cli.Port)
	addr = fmt.Sprintf("%s:%s", cli.IP, cli.Port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return err
	}
	cli.client = client

	session, err := cli.client.NewSession()
	if err != nil {
		return err
	}
	cli.session = session

	return nil
}

//执行shell，需要等待结果
//@param shell shell脚本命令
func (cli *Cli) Run(shell string) (string, error) {
	buf, err := cli.session.CombinedOutput(shell)
	cli.LastResult = string(buf)
	return cli.LastResult, err
}

//执行 shell，不需要等待结果
//@param shell shell脚本命令
func (cli *Cli) Start(shell string) (error) {
	err := cli.session.Start(shell)
	if err != nil {
		return err
	}
	return nil
}

// 关闭连接
func (cli *Cli) Close() {
	cli.session.Close()
	cli.client.Close()
}

// RunTerminal
func (cli *Cli) RunTerminal(shell string, stdout, stderr io.Writer) (error) {
	if cli.client == nil {
		if err := cli.sshConnect(); err != nil {
			return err
		}
	}

	session, err := cli.client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(fd, oldState)

	session.Stdout = stdout
	session.Stderr = stderr
	session.Stdin = os.Stdin

	termWidth, termHeight, err := terminal.GetSize(fd)
	if err != nil {
		panic(err)
	}

	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal
	if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
		return err
	}
	session.Run(shell)
	return nil
}
