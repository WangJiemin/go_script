package main

import (
	"fmt"
	"log"
	//"os"
	"flag"

	"github.com/WangJiemin/gocomm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type ForRentInfo struct {
	UUID             string `json:"uuid"gorm:"column:uuid"`
	Id               int    `json:"id"gorm:"column:id"`
	TITLE            string `json:"title"gorm:"column:title"`
	RELEASE_DATETIME string `json:"release_datetime"gorm:"column:release_datetime"`
	VIEWS            string `json:"views"gorm:"column:views"`
	TOWNSHIP         string `json:"township"gorm:"column:township"`
	LEASE_METHOD     string `json:"lease_method"gorm:"column:lease_method"`
	RENT             string `json:"rent"gorm:"column:rent"`
	CONTACT_PERSON   string `json:"contact_person"gorm:"column:contact_person"`
	CONTACT_NUM      string `json:"contact_num"gorm:"column:contact_num"`
	RENT_DETAILS     string `json:"rent_details"gorm:"column:rent_details"`
	FLOOR            string `json:"floor"gorm:"column:floor"`
	ORIENTATION      string `json:"orientation"gorm:"column:orientation"`
	PAYMENT_METHOD   string `json:"payment_method"gorm:"column:payment_method"`
	PICTURE          string `json:"picture"gorm:"column:picture"`
	COMMUNITY        string `json:"community"gorm:"column:community"`
	HOUSE_SOURCE     string `json:"house_source"gorm:"column:house_source"`
	URL              string `json:"url"gorm:"column:url"`
	DECORATE         string `json:"decorate"gorm:"column:decorate"`
	MEDIA            string `json:"media"gorm:"column:media"`
	RENTAL_PERIOD    string `json:"rental_period"gorm:"column:rental_period"`
	COMMUNITY_INFO   string `json:"community_info"gorm:"column:community_info"`
	CHECK_IN         string `json:"check_in"gorm:"column:check_in"`
	LIVING_ROOM      string `json:"living_room"gorm:"column:living_room"`
	SITE_URL         string `json:"site_url"gorm:"column:site_url"`
	DATA_MD5         string `json:"data_md5"gorm:"column:data_md5"`
	GET_TIME         string `json:"get_time"gorm:"column:get_time"`
}

type RentInfo struct {
	//ID                string `json:"id"gorm:"column:id"`
	TITLE             string `json:"title"gorm:"column:title"`
	RENTAL_PERIOD     string `json:"rental_period"gorm:"column:rental_period"`
	PRICE             string `json:"price"gorm:"column:price"`
	LIVING_ROOM       string `json:"living_room"gorm:"column:living_room"`
	AREA              string `json:"area"gorm:"column:area"`
	DETAILED_LOCATION string `json:"detailed_location"gorm:"column:detailed_location"`
	RELEASE_DATETIME  string `json:"release_datetime"gorm:"column:release_datetime"`
	PAYMENT_METHOD    string `json:"payment_method"gorm:"column:payment_method"`
	ORIENTATION       string `json:"orientation"gorm:"column:orientation"`
	LEASE_METHOD      string `json:"lease_method"gorm:"column:lease_method"`
	COMMUNITY         string `json:"community"gorm:"column:community"`
	TOWNSHIP          string `json:"township"gorm:"column:township"`
	HOUSE_DETAILS     string `json:"house_details"gorm:"column:house_details"`
	CONTACT_NUM       string `json:"contact_num"gorm:"column:contact_num"`
	PICTURE           string `json:"picture"gorm:"column:picture"`
	CONTACT_PERSON    string `json:"contact_person"gorm:"column:contact_person"`
	HOUSE_SOURCE      string `json:"house_source"gorm:"column:house_source"`
	MEDIA             string `json:"media"gorm:"column:media"`
	HOUSE_CONF        string `json:"house_conf"gorm:"column:house_conf"`
	URL               string `json:"url"gorm:"column:url"`
	FLOOR             string `json:"floor"gorm:"column:floor"`
	VIEWS             string `json:"views"gorm:"column:views"`
	DECORATE          string `json:"decorate"gorm:"column:decorate"`
	COMMUNITY_INFO    string `json:"community_info"gorm:"column:community_info"`
	SITE_URL          string `json:"site_url"gorm:"column:site_url"`
	CHECK_IN          string `json:"check_in"gorm:"column:check_in"`
	DATA_MD5          string `json:"data_md5"gorm:"column:data_md5"`
	GET_TIME          string `json:"get_time"gorm:"column:get_time"`
}

type SecondHandGood struct {
	//ID               string `json:"id"grom:"column:id"`
	PRICE            string `json:"price"grom:"column:price"`
	FINENESS         string `json:"fineness"grom:"column:fineness"`
	ADDRESS          string `json:"address"grom:"column:address"`
	SELLER           string `json:"seller"grom:"column:seller"`
	CONTACT_NUM      string `json:"contact_num"grom:"column:contact_num"`
	TITLE            string `json:"title"grom:"column:title"`
	PICTURE          string `json:"picture"grom:"column:picture"`
	QQ_NUM           string `json:"qq_num"grom:"column:qq_num"`
	URL              string `json:"url"grom:"column:url"`
	MEDIA            string `json:"media"grom:"column:media"`
	KEYWORD          string `json:"keyword"grom:"column:keyword"`
	SITE_URL         string `json:"site_url"grom:"column:site_url"`
	ITEM_DETAILS     string `json:"item_details"grom:"column:item_details"`
	USED_ITEM        string `json:"used_item"grom:"column:used_item"`
	RELEASE_DATETIME string `json:"release_datetime"grom:"column:release_datetime"`
	VIEWS            string `json:"views"grom:"column:views"`
	TRANSFER_SOURCE  string `json:"transfer_source"grom:"column:transfer_source"`
	ORIGINAL_PRICE   string `json:"original_price"grom:"column:original_price"`
	FACTORY_TIME     string `json:"factory_time"grom:"column:factory_time"`
	USAGE_TIME       string `json:"usage_time"grom:"column:usage_time"`
	INFO_NUM         string `json:"info_num"grom:"column:info_num"`
	COMMENTS         string `json:"comments"grom:"column:comments"`
	WANNA_NUM        string `json:"wanna_num"grom:"column:wanna_num"`
	DATA_MD5         string `json:"data_md5"gorm:"column:data_md5"`
	GET_TIME         string `json:"get_time"gorm:"column:get_time"`
}

type SecondHandCar struct {
	//ID                   string `json:"id"gorm:"column:id"`
	TITLE                string `json:"title"gorm:"column:title"`
	PRICE                string `json:"price"gorm:"column:price"`
	RELEASE_DATETIME     string `json:"release_datetime"gorm:"column:release_datetime"`
	MARKET_PRICE         string `json:"market_price"gorm:"column:market_price"`
	CONTACT_NUM          string `json:"contact_num"gorm:"column:contact_num"`
	SELLER               string `json:"seller"gorm:"column:seller"`
	ADDRESS              string `json:"address"gorm:"column:address"`
	PICTURE              string `json:"picture"gorm:"column:picture"`
	URL                  string `json:"url"gorm:"column:url"`
	MEDIA                string `json:"media"gorm:"column:media"`
	SITE_URL             string `json:"site_url"gorm:"column:site_url"`
	DISPLACEMENT         string `json:"displacement"gorm:"column:displacement"`
	CAR_COLOR            string `json:"car_color"gorm:"column:car_color"`
	CARD_TIME            string `json:"card_time"gorm:"column:card_time"`
	COMPULSORY_INSURANCE string `json:"compulsory_insurance"gorm:"column:compulsory_insurance"`
	GEARBOX              string `json:"gearbox"gorm:"column:gearbox"`
	BODY_STRUCTURE       string `json:"body_structure"gorm:"column:body_structure"`
	USE_NATURE           string `json:"use_nature"gorm:"column:use_nature"`
	ACCIDENT_SITUATION   string `json:"accident_situation"gorm:"column:accident_situation"`
	EMISSION_STANDARDS   string `json:"emission_standards"gorm:"column:emission_standards"`
	DRIVEN_DISTANCE      string `json:"driven_distance"gorm:"column:driven_distance"`
	ANNUAL_CHECK         string `json:"annual_check"gorm:"column:annual_check"`
	CAR_AGE              string `json:"car_age"gorm:"column:car_age"`
	SELLER_MESSAGE       string `json:"seller_message"gorm:"column:seller_message"`
	CAR_DIS              string `json:"car_dis"gorm:"column:car_dis"`
	TRANSFER_SOURCE      string `json:"transfer_source"gorm:"column:transfer_source"`
	VIEWS                string `json:"views"gorm:"column:views"`
	TRANSFER_NUM         string `json:"transfer_num"gorm:"column:transfer_num"`
	DISTRICT             string `json:"district"gorm:"column:district"`
	INFO_NUM             string `json:"info_num"gorm:"column:info_num"`
	MAINTAIN             string `json:"maintain"gorm:"column:maintain"`
	COMMERCIAL_INSURANCE string `json:"commercial_insurance"gorm:"column:commercial_insurance"`
	DATA_MD5             string `json:"data_md5"gorm:"column:data_md5"`
	GET_TIME             string `json:"get_time"gorm:"column:get_time"`
}

var (
	db205 *gorm.DB
	db140 *gorm.DB
	err   error
)

func Initialize(ConfigPath string) {
	var dbHost, dbUser, dbPwd, dbPort, dbName string
	config := gocomm.ReadConfig(ConfigPath)
	dbHost = config["db_host_s"]
	dbUser = config["db_user_s"]
	dbPwd = config["db_pass_s"]
	dbPort = config["db_port_s"]
	dbName = config["db_name_s"]

	dbConnectBaseStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnect_s := fmt.Sprintf(dbConnectBaseStr, dbUser, dbPwd, dbHost, dbPort, dbName)
	fmt.Print(dbConnect_s)
	db205, err = gorm.Open("mysql", dbConnect_s)
	if err != nil {
		fmt.Println(err)
	}
	//defer db205.Close()
	db205.LogMode(true)

	dbHost = config["db_host_d"]
	dbUser = config["db_user_d"]
	dbPwd = config["db_pass_d"]
	dbPort = config["db_port_d"]
	dbName = config["db_name_d"]

	dbConnect_d := fmt.Sprintf(dbConnectBaseStr, dbUser, dbPwd, dbHost, dbPort, dbName)
	fmt.Print(dbConnect_d, "\n")
	db140, err = gorm.Open("mysql", dbConnect_d)
	if err != nil {
		fmt.Println(err)
	}
	//defer db140.Close()
	db140.LogMode(true)
}

func main() {

	Initialize("./app.cnf")

	// os.Args方式
	//user_arg := os.Args
	//fmt.Printf("\narray is:%s\n", user_arg)

	// flag.Args方式
	flag.Parse()
	var user_arg []string = flag.Args()
	//fmt.Printf("\narray is:%s\n", user_arg)

	array_Str := [...]string{0: "find_rent_xingtai", 1: "rent_xingtai", 2: "second_hand_goods", 3: "second_hand_cars"}
	//if len(user_arg) > 1 {
	if user_arg != nil && len(user_arg) > 0 {
		switch {
		case user_arg[0] == array_Str[0]:
			Forreninfo()
			//fallthrough
		case user_arg[0] == array_Str[1]:
			Rentinfo()
			//fallthrough
		case user_arg[0] == array_Str[2]:
			Secondhandgood()
			//fallthrough
		case user_arg[0] == array_Str[3]:
			Secondhandcar()
			//fallthrough
		default:
			log.Fatal("Ple input:\n\ttable_name:[find_rent_xingtai or rent_xingtai or second_hand_goods or second_hand_cars]")
		}
	} else {
		log.Fatalln("No command line arguments are passed in,exit with 1")
	}

	db205.Close()
	db140.Close()
}

func Forreninfo() {
	table_db_205 := []*ForRentInfo{}
	db205.Table("find_rent_xingtai").Find(&table_db_205)
	for _, v := range table_db_205 {
		db140.Table("find_rent_xingtai").Save(&v)
		//db140.Table("find_rent_xingtai_1").Save(&v)
	}
}

func Rentinfo() {
	table_db_205_2 := []*RentInfo{}
	db205.Table("rent_xingtai").Find(&table_db_205_2)
	for _, v := range table_db_205_2 {
		db140.Table("rent_xingtai").Save(&v)
		//db140.Table("rent_xingtai_1").Save(&v)
	}
}

func Secondhandgood() {
	table_db_205_3 := []*SecondHandGood{}
	db205.Table("second_hand_goods").Find(&table_db_205_3)
	for _, v := range table_db_205_3 {
		db140.Table("second_hand_goods").Save(&vip)
		//db140.Table("second_hand_goods_1").Save(&v)
	}
}

func Secondhandcar() {
	table_db_205_4 := []*SecondHandCar{}
	db205.Table("second_hand_cars").Find(&table_db_205_4)
	for _, v := range table_db_205_4 {
		db140.Table("second_hand_cars").Save(&v)
		//db140.Table("second_hand_cars_1").Save(&v)
	}
}
