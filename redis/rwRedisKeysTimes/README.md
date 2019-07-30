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

## 帮助
```bash
db-pre-master ~ # /usr/local/bin/rwRedisKeysTimes -h
flag needs an argument: -h
rwRedisKeysTimes V0.1 By WangJiemin.
	E_mail: 278667010@qq.com

***************************************************************************************
	system_command: ./rwRedisKeysTimes
	system_goos: linux
	system_arch: amd64
	hostname: db-pre-master.zw.babytree-ops.org
	hostaddress: 10.25.1.78
	blog: https://jiemin.wang
***************************************************************************************

  -d int
    	Redis databases. default: 0
  -h string
    	Redis address and port. default: 127.0.0.1:6379 (default "127.0.0.1:6379")
  -k int
    	Redis Keys number. default: 10 (default 10)
  -m string
    	Redis models. options: client/cluster. default: cluster (default "cluster")
  -p string
    	Redis password. default: ""
  -t duration
    	Redis keys expiration time. default: 30 Second (default 30ns)
  -v	print version
db-pre-master ~ #
db-pre-master ~ # ./rwRedisKeysTimes -m="cluster" -h="10.40.7.21:6881,10.25.1.78:6881,10.40.7.21:6883,10.25.1.78:6883,10.40.7.21:6882,10.25.1.78:6880,10.40.7.22:6882,10.40.7.21:6880,10.25.1.78:6884,10.40.7.22:6881,10.40.7.22:6880,10.40.7.22:6883,10.40.7.22:6884,10.25.1.78:6882,10.40.7.21:6884" -k=20000 -t 2m30s
```

## 截图
![](https://github.com/WangJiemin/go_script/blob/master/redis/rwRedisKeysTimes/image/WechatIMG858.png)
![](https://github.com/WangJiemin/go_script/blob/master/redis/rwRedisKeysTimes/image/WechatIMG859.png)
![](https://github.com/WangJiemin/go_script/blob/master/redis/rwRedisKeysTimes/image/WechatIMG860.png)

