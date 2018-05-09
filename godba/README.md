作者:王杰民(278667010@qq.com)

版本: v0.1.1

程序介绍:

    参考阿里巴巴开源的dba工具orzdba

    功能全部用Golang的基础知识实现一个简单的模块化编程的示例
	
	

程序结构:
```jshelllanguage
Shell + Golang
```

编译：

Mac 
```go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```
linux
```go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
```