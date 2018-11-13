package models

/*
import "time"

type IncText struct {
	ID         int       `gorm:"primary_key;AUTO_INCREMENT;not null;column:id"`
	CommentId  int64
	BlogId     int64
	CreateAt   string    `gorm:"type:varchar(32);default null`
	Text       string
	Source     string    `gorm:"type:varchar(100);default null`
	LikeCounts int
	ReplyId    int64
	ReplyText  string
	UserId     int64
	ScreenName string    `gorm:"type:varchar(100);default null`
	InsertTime time.Time `gorm:"type:timestamp;default null`
}
*/

type IncInformationInfo struct {
	ID            int    // 用来表示检查的sql序号
	STAGE         string // 这个列显示当前语句进行的状态
	ERRLEVEL      int    // 返回值
	STAGESTATUS   string // 表示检查及执行的过程是成功还是失败
	ERRORMESSAGE  string // 表示出错错误信息,用换行符分隔
	SQL           string // 用来表示当前检查的是哪条sql语句
	AFFECTED_ROWS int    // 表示当前语句执行时预计影响的行数，在执行时显示的是真实影响行数
	SEQUENCE      string // 这个列与上面说的备份功能有关系，其实就是对应$$Inception_backup_information$$表中的 opid_time 这个列
	BACKUP_DBNAME string // 表示的是当前语句产生的备份信息，存储在备份服务器的哪个数据库中，这是一个字符串类型的值，只针对需要备份的语句，数据库名由IP地址、端口、源数据库名组成，由下划线连接，而如果是不需要备份的语句，则返回字符串None。
	EXECUTE_TIME  string // 表示当前语句执行时间，单位为秒，精确到小数点后两位。列类型为字符串, 使用时可能需要转换成DOUBLE类型的值，如果只是审核而不执行，则这个列返回的值为0。
	SQLSHA1       string // 存储当前这个语句的一个HASH值，这是用来标识这个语句是不是会使用OSC功能，如果返回信息中有值，则表示这个语句在执行的时候会使用OSC，因为在执行前，会有一次单独的审核操作，此时上层已经可以拿到这个值，审核通过之后，语句是不会改变的，当然这个值也不会改变，那么在执行时就可以使用这个值来查看OSC执行的进度等信息，这个值一般长的样子如下：D0210DFF35F0BC0A7C95CD98F5BCD4D9B0CA8154*
}

type UserConfig struct {
	Inc struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DB       string `yaml:"db"`
	}
	Remote struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DB       string `yaml:"db"`
		DBTable  string `yaml:"dbtable"`
		//SqlFormat string `yaml:"sqlformat"`
	}
}
