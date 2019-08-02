## 简介
学习go 生成二维码

## 执行
```
go run qrcode.go
```

## 编译 
Linux:
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o qrcode_linux_amd64 -ldflags "-s -w" qrcode.go
```
Windows:
```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o qrcode_windows_amd64.exe -ldflags "-s -w" qrcode.go
```
MacOS:
```
go build -o qrcode_osx_amd64 -ldflags "-s -w" qrcode.go
```

## 使用
```
Last login: Fri Aug  2 16:51:23 2019 from 172.16.9.10
dev-oplog-mongodb1 ~ # ./qrcode_linux_amd64 -h
Usage of ./qrcode_linux_amd64:
  -o string
        输出文件
  -t string
        二维图内容
dev-oplog-mongodb1 ~ # ./qrcode_linux_amd64 --help
Usage of ./qrcode_linux_amd64:
  -o string
        输出文件
  -t string
        二维图内容
dev-oplog-mongodb1 ~ # 
```
![](https://github.com/WangJiemin/go_script/blob/master/qrcode_encoder/reader_qrcode/image/qrcode.gif)

## 帮助
[Email](278667010@qq.com)

