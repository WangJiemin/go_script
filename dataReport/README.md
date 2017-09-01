作者:WangJiemin(Jamin.Wang)-王杰民

版本:示例版本 v0.1

程序介绍:

    实现了获取数据生成EXCEL数据报表功能

    功能全部用Golang的基础知识实现一个简单的模块化编程的示例

    注意:只实现了"生成EXCEL"和"发送Email"

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
	├── models   # models目录("操作数据库")
	│	├──	models.go 	# 定义数据库结构文件
	│	└── datareport.go # 数据库操作文件
	│
	├── sendmail # Email目录
	│   └── sendmail.go # 发送Email文件
	└── main.go  # golang main文件

	
程序版本/

	├── v0.1版本
	│	└── 实现从MySQL数据库中获取数据生成EXCEL表发送到邮箱里(现实)
	│
	├── v0.2版本
	│       └── 实现把生成的EXCEL合并成一个EXCEL。一个EXCEL中包含不同的所有表(未实现,已有代码测试阶段)
	│
	└── v0.3版本
		└── 实现把操作数据库(优化)大SQL拆成小SQL,进行字段直接的拼接组成。
