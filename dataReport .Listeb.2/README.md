作者:WangJiemin(Jamin.Wang)-王杰民

版本:示例版本 v0.2

程序介绍:

    实现了获取数据生成EXCEL数据报表功能

    功能全部用Golang的基础知识实现一个简单的模块化编程的示例
	
	

程序结构:

dataReport/

	├── README
	│
	├── compress # 压缩解压缩("支持tar.gz/zip")
	│   ├── tar_gz.go 	# tar.gz的压缩
	│   ├── untar_gz.go # tar.gz的解压缩
	│   ├── unzip.go    # zip的压缩
	│   └── zip.go	    # zip的解压缩
	│
	├── conf     # 配置文件目录
	│   └── app.cnf     # 项目配置文件
	│
	├── execl    # EXCEL目录
	│	└── execl.go    # 生成EXCEL
	│
	├── gettime    # 时间目录
	│	└── getnowtime.go    # 获取当前时间与前一天的时间
	│
	├── models   # models目录("操作数据库")
	│	├──	models.go 	# 定义数据库结构文件
	│	└── datareport.go # 数据库操作文件
	│
	├── sendmail # Email目录
	│   └── sendmail.go # 发送Email文件
	└── main.go  # golang main文件

	
程序版本/

	├── v0.1版本
	│	└── 实现从MySQL数据库中获取数据生成EXCEL表发送到邮箱里(上线)
	│
	├── v0.1.1版本
	│	└── 实现从MySQL数据库中获取当前时间和一天前的数据生成EXCEL表发送到邮箱里(上线)
	│
	├── v0.2版本
	│       └── 实现把数据库按天分表进行统计查询
	└── 
