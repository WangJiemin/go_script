## 简介

rwRedisKeysTime 是测试redis cluster/redis 主从/redis单节点 读写速度脚本

## 使用
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o rwRedisKeysTimes -ldflags "-s -w" rwRedisKeysTimes.go
```
或者
```bash
go run rwRedisKeysTimes.go
```

## redis client driver
```bash
go get -u github.com/go-redis/redis
```

```bash
db-pre-master ~ # /usr/local/bin/rwRedisKeysTimes -h
flag needs an argument: -h
rwRedisKeysTimes V0.1 By WangJiemin.
	E_mail: 278667010@qq.com

***************************************************************************************
	system_command: /usr/local/bin/rwRedisKeysTimes
	system_goos: linux
	system_arch: amd64
	hostname: db-pre-master.zw.babytree-ops.org
	hostaddress: 10.25.1.78
	blog: https://jiemin.wang
***************************************************************************************

  -d int
    	Redis databases. default: 0
  -h string
    	Redis address and port. default: 127.0.0.1:6390 (default "127.0.0.1:6390")
  -m string
    	Redis models. options: client/cluster. default: cluster (default "cluster")
  -p string
    	Redis password. default: ""
  -t int
    	Redis Keys number. default: 10 (default 10)
  -v	print version
db-pre-master ~ #
```

## 截图
![]()
![]()
![]()

