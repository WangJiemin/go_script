package slowlog

import (
	"fmt"
	"io/ioutil"
	"regexp"
	//"strings"
	yaml "gopkg.in/yaml.v2"
	dt "golang_Learn/ptslowlog/datetime"
	ml "golang_Learn/ptslowlog/models"
	oe "golang_Learn/ptslowlog/osexce"
	"golang_Learn/ptslowlog/db"
)

var (
	enable    bool
	ptcommand string
	logpath   string
	slowname  string
	orderby   string
	limit     string
	filter    string
	user      string
	pass      string
	host      string
	port      string
	dbname    string
)

/*
func Initialize(config map[string]string) {

}
*/

func checkErr(errinfo error) {
	if errinfo != nil {
		fmt.Println(errinfo.Error())
	}
}

func SlowLogForMart(config string) (string) {
	conf := ml.YamlStruct{}
	yamlFile, err := ioutil.ReadFile(config)
	checkErr(err)
	//fmt.Println("yamlFile: ", string(yamlFile))
	err = yaml.Unmarshal(yamlFile, &conf)
	checkErr(err)
	//fmt.Println("conf: ", conf)

	//dy,ysdy:=dt.DateTime()
	_, ysdy := dt.DateTime()

	enable = conf.PtQueryDigest.Enable
	if (enable == true) {
		ptcommand = conf.PtQueryDigest.Ptcommand
		logpath = conf.PtQueryDigest.LogPath
		slowname = conf.PtQueryDigest.SlowName
		orderby = conf.PtQueryDigest.Values.OrderBy
		limit = conf.PtQueryDigest.Values.Limit
		filter = conf.PtQueryDigest.Values.Filter
		user = conf.Mysql.User
		pass = conf.Mysql.Password
		host = conf.Mysql.Host
		port = conf.Mysql.Port
		dbname = conf.Mysql.DbName

	}

	/*
	fmt.Printf("yaml connent enable is %s\n", enable)
	fmt.Printf("yaml connent logpath is %s\n", logpath)
	fmt.Printf("yaml connent slowname is %s\n", slowname)
	fmt.Printf("yaml connent orderby is %s\n", orderby)
	fmt.Printf("yaml connent limit is %s\n", limit)
	fmt.Printf("yaml connent filter is %s\n", filter)
	fmt.Printf("yaml connent user is %s\n", user)
	fmt.Printf("yaml connent pass is %s\n", pass)
	fmt.Printf("yaml connent host is %s\n", host)
	fmt.Printf("yaml connent port is %s\n", port)
	fmt.Printf("yaml connent dbname is %s\n", dbname)
	*/

	slowlogname := fmt.Sprintf("%s%s%s_%s", logpath, "/", slowname, ysdy)
	//fmt.Println(slowlogname)

	slowlog_cmd := "mv " + logpath + "/" + slowname + " " + slowlogname
	oe.ExecCommand(slowlog_cmd)

	db.ExceFulshLogs(user, pass, host, dbname, port)

	//grepcommands := " | egrep 'Time range|Attribute|============|Exec time|Lock time|Rows sent|Rows examine|Rows affecte|Bytes sent|Query size|Databases|Hosts|Users|SELECT' | grep -v '0x'"
	//pt_commands_cmd := ptcommand + " --filter" + " '" + filter + "'" + " --order-by " + orderby + " --limit=" + limit + " " + slowlogname + grepcommands
	pt_commands_cmd := ptcommand + " --filter" + " '" + filter + "'" + " --order-by " + orderby + " --limit=" + limit + " " + slowlogname
	//fmt.Println(pt_commands_cmd)
	pt_commands_str := oe.ExecCommand(pt_commands_cmd)
	//fmt.Println(pt_commands_str)
	re, _ := regexp.Compile("# ")
	pt_commands_str_grep := re.ReplaceAllString(pt_commands_str, "")
	//fmt.Println(len(pt_commands_str_grep))

	//var pt_slowlog_split string
	//for i := 0; i <= len(pt_commands_str_grep); i++ {
	//	index := 0
	//	pt_slowlog_split = strings.Split(pt_commands_str_grep, "\n")[index]
	//	if (len(strings.TrimSpace(pt_slowlog_split)) <= 0) {
	//		index = index + 1
	//		pt_slowlog_split = strings.Split(pt_commands_str_grep, "\n")[index]
	//	}
	//}
	//return pt_slowlog_split
	return pt_commands_str_grep
}
