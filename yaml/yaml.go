package main

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"fmt"
	"log"
)

type Yaml struct {
	Mysql struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DbName   string `yaml:"dbname"`
	}
	PtQueryDigest struct {
		Enable   bool   `yaml:"enable"`
		LogPath  string `yaml:"logpath"`
		SlowName string `yaml:"slowname"`
		Values struct {
			OrderBy string `yaml:"orderby"`
			Limit   string `yaml:"limit"`
			Filter  string `yaml:"filter"`
		}
	}
}

/*func Initialize(config map[string]string) {

}*/

func main() {
	conf := Yaml{}
	//conf := new(Yaml)
	yamlFile, err := ioutil.ReadFile("/home/jiemin/code/go_dev/src/golang_Learn/yaml/conf.yaml")
	if err != nil {
		fmt.Println("yamlFile.Get %v", err)
	}

	//fmt.Println("yamlFile:", yamlFile)
	fmt.Println("yamlFile:", string(yamlFile))

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	fmt.Println("conf:", conf)

	fmt.Println("#########################################################################################")

	user := conf.Mysql.User
	pass := conf.Mysql.Password
	host := conf.Mysql.Host
	port := conf.Mysql.Port
	dbname := conf.Mysql.DbName

	enable := conf.PtQueryDigest.Enable
	logpath := conf.PtQueryDigest.LogPath
	slowname := conf.PtQueryDigest.SlowName
	orderby := conf.PtQueryDigest.Values.OrderBy
	limit := conf.PtQueryDigest.Values.Limit
	filter := conf.PtQueryDigest.Values.Filter

	fmt.Printf("yaml connent user is %s\n", user)
	fmt.Printf("yaml connent pass is %s\n", pass)
	fmt.Printf("yaml connent host is %s\n", host)
	fmt.Printf("yaml connent port is %s\n", port)
	fmt.Printf("yaml connent dbname is %s\n", dbname)
	fmt.Printf("yaml connent enable is %s\n", enable)
	fmt.Printf("yaml connent logpath is %s\n", logpath)
	fmt.Printf("yaml connent slowname is %s\n", slowname)
	fmt.Printf("yaml connent orderby is %s\n", orderby)
	fmt.Printf("yaml connent limit is %s\n", limit)
	fmt.Printf("yaml connent filter is %s\n", filter)
	/*
		resultMap := make(map[string]interface{})
		yamlFile, err := ioutil.ReadFile("/home/jiemin/code/go_dev/src/golang_Learn/yaml/conf.yaml")
		if err != nil {
			fmt.Println("yamlFile.Get %v", err)
		}

		//fmt.Println("yamlFile:", yamlFile)
		fmt.Println("yamlFile:", string(yamlFile))

		err = yaml.Unmarshal(yamlFile, &resultMap)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
		fmt.Println(resultMap)

		fmt.Println("#########################################################################################")
		//
		//for k, v := range resultMap {
		//	switch vv := v.(type) {
		//	case string:
		//		fmt.Println(k, "is string", vv)
		//	case int:
		//		fmt.Println(k, "is int", vv)
		//	case []interface{}:
		//		fmt.Println(k, "is an array:")
		//		for i, u := range vv {
		//			fmt.Println(i, u)
		//		}
		//	default:
		//		fmt.Println(k, "is of a type I don't know to handle")
		//	}
		//}

		for k, v := range resultMap {
			fmt.Println(k, v)
		}

		//user := resultMap["mysql"].(map[string]interface{})["user"]

		//fmt.Printf("yaml connent user is %s\n", user)
	*/
}
