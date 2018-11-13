package main

import (
	"fmt"
	"strings"

	"github.com/WangJiemin/gocomm/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db205 *gorm.DB
var db140 *gorm.DB

type TaskResults struct {
	TASKID   string `gorm:"column:task_id"`
	DATA_MD5 string `gorm:"column:data_md5"`
	URL      string `gorm:"column:url"`
	CONTENT  string `gorm:"column:content"`
	GET_TIME string `gorm:"column:get_time"`
}

/*
type ForRentInfo struct {
	ID             string `json:"id"gorm:"column:id"`
	TITLE          string `json:"title"gorm:"column:title"`
	RELEASE_TIME   string `json:"release_time"gorm:"column:release_time"`
	VIEWS          string `json:"views"gorm:"column:views"`
	TOWNSHIP       string `json:"township"gorm:"column:township"`
	LEASE_METHOD   string `json:"lease_method"gorm:"column:lease_method"`
	RENT           string `json:"rent"gorm:"column:rent"`
	CONTACT_PERSON string `json:"contact_person"gorm:"column:contact_person"`
	CONTACT_NUM    string `json:"contact_num"gorm:"column:contact_num"`
	RENT_DETAILS   string `json:"rent_details"gorm:"column:rent_details"`
	FLOOR          string `json:"floor"gorm:"column:floor"`
	ORIENTATION    string `json:"orientation"gorm:"column:orientation"`
	PAYMENT_METHOD string `json:"payment_method"gorm:"column:payment_method"`
	PICTURE        string `json:"picture"gorm:"column:picture"`
	COMMUNITY      string `json:"community"gorm:"column:community"`
	HOUSE_SOURCE   string `json:"house_source"gorm:"column:house_soource"`
	URL            string `json:"url"gorm:"column:url"`
	DECORATE       string `json:"decorate"gorm:"column:decorate"`
	MEDIA          string `json:"media"gorm:"column:media"`
	RENTAL_PERIOD  string `json:"rental_period"gorm:"column:rental_period"`
	COMMUNITY_INFO string `json:"community_info"gorm:"column:community_info"`
	CHECK_IN       string `json:"check_in"gorm:"column:check_in"`
	LIVING_ROOM    string `json:"living_room"gorm:"column:living_room"`
	SITE_URL       string `json:"site_url"gorm:"column:site_url"`
	DATA_MD5       string `json:"data_md5"gorm:"column:data_md5"`
	GET_TIME       string `json:"get_time"gorm:"column:get_time"`
}

type RentInfo struct {
	ID                string `json:"id"gorm:"column:id"`
	TITLE             string `json:"title"gorm:"column:title"`
	RENTAL_PERIOD     string `json:"rental_period"gorm:"column:rental_period"`
	PRICE             string `json:"price"gorm:"column:price"`
	LIVING_ROOM       string `json:"living_room"gorm:"column:living_room"`
	AREA              string `json:"area"gorm:"column:area"`
	DETAILED_LOCATION string `json:"detailed_location"gorm:"column:detailed_location"`
	RELEASE_TIME      string `json:"release_time"gorm:"column:release_time"`
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
	ID               string `json:"id"grom:"column:id"`
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
	ID                   string `json:"id"gorm:"column:id"`
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
*/

type FINACIALPRODUCTS struct {
	NAME                       string `json:"name"gorm:"column:name"`
	PROJECTCODE                string `json:"projectCode"gorm:"column:projectCode"`
	PRODUCTCATAGORY            string `json:"product_catagory"gorm:"column:product_catagory"`
	CURRENCY                   string `json:"currency"gorm:"column:currency"`
	SALESTARTDATE              string `json:"sale_start_date"gorm:"column:sale_start_date"`
	SALEENDDATE                string `json:"sale_end_date"gorm:"column:sale_end_date"`
	VALUEDATE                  string `json:"value_date"gorm:"column:value_date"`
	ESTABLISHMENTDATE          string `json:"establishment_date"gorm:"column:establishment_date"`
	DUEDATE                    string `json:"due_date"gorm:"column:due_date"`
	TREM                       string `json:"term"gorm:"column:term"`
	OPENDATE                   string `json:"open_date"gorm:"column:open_date"`
	EXPECTEDRATE               string `json:"expected_rate"gorm:"column:expected_rate"`
	PRODUCTSTATUS              string `json:"product_status"gorm:"column:product_status"`
	RISKLEVEL                  string `json:"risk_level"gorm:"column:risk_level"`
	INCOMETYPE                 string `json:"income_type"gorm:"column:income_type"`
	PURCHASEAMOUNT             string `json:"purchase_amount"gorm:"column:purchase_amount"`
	DISTRICT                   string `json:"district"gorm:"column:district"`
	INCREMENTALAMOUNT          string `json:"incremental_amount"gorm:"column:incremental_amount"`
	PRODEUCTTYPE               string `json:"product_type"gorm:"column:product_type"`
	COSTOMER                   string `json:"costomer"gorm:"column:costomer"`
	DISTRIBUTIONCHANNEL        string `json:"distribution_channel"gorm:"column:distribution_channel"`
	REMAININGQUOTA             string `json:"remaining_quota"gorm:"column:remaining_quota"`
	PRODUCTMANUAL              string `json:"product_manual"gorm:"column:product_manual"`
	TRANSFERABLE               string `json:"transferable"gorm:"column:transferable"`
	MANAGEMENT                 string `json:"management"gorm:"column:management"`
	BUYUNIT                    string `json:"buy_unit"gorm:"column:buy_unit"`
	NETPRODUCT                 string `json:"net_product"gorm:"column:net_product"`
	REDEMPTIONLIMIT            string `json:"redemption_limit"gorm:"column:redemption_limit"`
	REDEMPTIONLOWERLIMIT       string `json:"redemption_lower_limit"gorm:"column:redemption_lower_limit"`
	MINIMUMSUBSCRIBE           string `json:"minimum_subscribe"gorm:"column:minimum_subscribe"`
	SUBSCRIPTIONCAP            string `json:"subscription_cap"gorm:"column:subscription_cap"`
	SUBSCRIPTIONLOWLIMIT       string `json:"subscription_lowlimit"gorm:"column:subscription_lowlimit"`
	PURCHASECAP                string `json:"purchase_cap"gorm:"column:purchase_cap"`
	PURCHASELOWLIMIT           string `json:"purchase_lowlimit"gorm:"column:purchase_lowlimit"`
	REDEMPTIONBASE             string `json:"redemption_base"gorm:"column:redemption_base"`
	SUBSCRIPTIONBASE           string `json:"subscription_base"gorm:"column:subscription_base"`
	PERCHASE_BASE              string `json:"perchase_base"gorm:"column:perchase_base"`
	REDEMPTIONALLOWED          string `json:"redemption_allowed"gorm:"column:redemption_allowed"`
	RESERVATIONREDEMPTION      string `json:"reservation_redemption"gorm:"column:reservation_redemption"`
	REALTIMEREDEMPTION         string `json:"real_time_redemption"gorm:"column:real_time_redemption"`
	EXCLUSIVEFINANCIALPRODUCTS string `json:"exclusive_financial_products"gorm:"column:exclusive_financial_products"`
	EXCLUSIVEVALIDDATE         string `json:"exclusive_validDate"gorm:"column:exclusive_validDate"`
	DAYPURCHASELIMIT           string `json:"day_purchase_limit"gorm:"column:day_purchase_limit"`
	REMITTANCEMARK             string `json:"remittance_mark"gorm:"column:remittance_mark"`
	MINIMUMPOSITION            string `json:"minimum_position"gorm:"column:minimum_position"`
	NEXTOPENDATE               string `json:"next_open_date"gorm:"column:next_open_date"`
	APPOINTMENTDATE            string `json:"appointment_date"gorm:"column:appointment_date"`
	CURRENTNETWORTH            string `json:"current_net_worth"gorm:"column:current_net_worth"`
	NETWORTHDATE               string `json:"net_worth_date"gorm:"column:net_worth_date"`
	BONUSWAY                   string `json:"bonus_way"gorm:"column:bonus_way"`
	GET_TIME                   string `json:"get_time"gorm:"column:get_time"`
	URL                        string `json:"url"gorm:"column:url"`
	BANKNAME                   string `json:"bankName"gorm:"column:bankName"`
}

func main() {
	// DB is SOURCE
	/*
		db_user_S := "jiemin"
		db_pass_S := "jiemin2017"
		db_host_S := "127.0.0.1"
		db_port_S := 3306
		db_name_S := "jiemin"
	*/

	db_user_S := "funbird"
	db_pass_S := "funbird2017"
	db_host_S := "10.10.206.205"
	db_port_S := 3306
	db_name_S := "funbird"
	//db_name_S := "funbird_analysis"

	dbConnectBaseStr_S := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnect_S := fmt.Sprintf(dbConnectBaseStr_S, db_user_S, db_pass_S, db_host_S, db_port_S, db_name_S)
	fmt.Print(dbConnect_S)
	db205, err := gorm.Open("mysql", dbConnect_S)
	if err != nil {
		fmt.Println(err)
	}
	defer db205.Close()
	db205.LogMode(true)

	// DB is DESC
	/*
		db_user_D := "jiemin"
		db_pass_D := "jiemin2017"
		db_host_D := "127.0.0.1"
		db_port_D := 3306
		db_name_D := "jiemin"
	*/

	db_user_D := "iscloud"
	db_pass_D := "iscloud"
	db_host_D := "192.168.95.140"
	db_port_D := 3306
	db_name_D := "network_data"

	dbConnectBaseStr_D := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnect_D := fmt.Sprintf(dbConnectBaseStr_D, db_user_D, db_pass_D, db_host_D, db_port_D, db_name_D)
	fmt.Print(dbConnect_D)
	db140, err := gorm.Open("mysql", dbConnect_D)
	if err != nil {
		fmt.Println(err)
	}
	defer db140.Close()
	db140.LogMode(true)

	// Business code
	/*
		results := []*TaskResults{}
		//db205.Raw("select * from task_results where task_id in (select task_id from task_infos where is_error=0 and project_id ='20564fc6-12fd-4086-a7a4-c26febe349cc')").Find(&results)
		db205.Table("task_results a").Select("a.*").Joins("JOIN task_infos b on b.task_id=a.task_id").Where("a.is_error=0 and b.project_id = '20564fc6-12fd-4086-a7a4-c26febe349cc'").Find(&results)
		for _, result := range results {
			data := ForRentInfo{}
			str := strings.Replace(result.CONTENT, "\\r", "", -1)
			json.Unmarshal([]byte(str), &data)
			data.DATA_MD5 = result.DATA_MD5
			fmt.Printf("Md5:%d\n", data.ID)
			db140.Table("find_rent_xingtai").Create(&data)
		}

		results2 := []*TaskResults{}
		//db205.Raw("select * from task_results where task_id in (select task_id from task_infos where is_error=0 and project_id ='5939d727-2c8d-41e3-9426-2f1497fd1018')").Find(&results2)
		db205.Table("task_results a").Select("a.*").Joins("JOIN task_infos b on b.task_id=a.task_id").Where("a.is_error=0 and b.project_id = '5939d727-2c8d-41e3-9426-2f1497fd1018'").Find(&results2)
		for _, result := range results2 {
			data := RentInfo{}
			str := strings.Replace(result.CONTENT, "\\r", "", -1)
			json.Unmarshal([]byte(str), &data)
			data.DATA_MD5 = result.DATA_MD5
			fmt.Printf("Md5:%d\n", data.ID)
			db140.Table("rent_xingtai").Create(&data)
		}

		results3 := []*TaskResults{}
		//db205.Raw("select * from task_results where task_id in (select task_id from task_infos where is_error=0 and project_id ='8cefe08b-40c7-4df1-8305-147c45c52170')").Find(&results3)
		db205.Table("task_results a").Select("a.*").Joins("JOIN task_infos b on b.task_id=a.task_id").Where("a.is_error=0 and b.project_id = '8cefe08b-40c7-4df1-8305-147c45c52170'").Find(&results3)
		for _, result := range results3 {
			data := SecondHandCar{}
			str := strings.Replace(result.CONTENT, "\\r", "", -1)
			json.Unmarshal([]byte(str), &data)
			data.DATA_MD5 = result.DATA_MD5
			fmt.Printf("Md5:%d\n", data.ID)
			db140.Table("second_hand_cars").Create(&data)
		}

		results4 := []*TaskResults{}
		//db205.Raw("select * from task_results where task_id in (select task_id from task_infos where is_error=0 and project_id ='9887fe7d-b2d6-4562-ac7e-5d44b35246c0')").Find(&results4)
		db205.Table("task_results a").Select("a.*").Joins("JOIN task_infos b on b.task_id=a.task_id").Where("a.is_error=0 and b.project_id = '9887fe7d-b2d6-4562-ac7e-5d44b35246c0'").Find(&results4)
		for _, result := range results4 {
			data := SecondHandGood{}
			str := strings.Replace(result.CONTENT, "\\r", "", -1)
			json.Unmarshal([]byte(str), &data)
			data.DATA_MD5 = result.DATA_MD5
			fmt.Printf("Md5:%d\n", data.ID)
			db140.Table("second_hand_goods").Create(&data)
		}
	*/

	results5 := []*TaskResults{}
	db205.Raw("select * from task_results where task_id in (select task_id from task_infos where project_id='70220d5f-0c3b-4728-a543-24ea3d4448f4' and is_delete=0)").Find(&results5)
	for _, result := range results5 {
		data := FINACIALPRODUCTS{}
		str := strings.Replace(result.CONTENT, "\\r", "", -1)
		json.Unmarshal([]byte(str), &data)
		data.GET_TIME = result.GET_TIME
		db140.Table("finacial_products").Create(&data)
	}
}
