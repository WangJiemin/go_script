package main

import (
	"fmt"
	//"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Myphoto struct {
	Id                    int
	Title                 string
	Description           string
	Type                  int
	Status                int
	Tags                  string
	Rank                  int
	DisplayOrder          int
	Server                string
	FileType              string
	FileSize              int
	Storepath             string
	BigName               string
	BigWidth              int
	BigHeight             int
	MiddleName            string
	MiddleWidth           int
	MiddleHeight          int
	MiddlesquareName      string
	MiddlesquareWidth     int
	MiddlesquareHeight    int
	SmallsquareName       string
	SmallsquareWidth      int
	SmallsquareHeight     int
	SmallheadsquareName   string
	SmallheadsquareWidth  int
	SmallheadsquareHeight int
	PrivateCtrlId         int
	PrivateSetting        string
	AlbumId               int
	CreateTs              int
	UpdateTs              int
	UserId                int
	ExifTime              string
	Watermarkflag         int
}

type Myphotoalbum struct {
	Id              int
	Title           string
	Description     string
	ParentId        int
	FaceId          int
	IsDefault       int
	DisplayOrder    int
	PrivateCtrlType string
	PrivateCtrlId   int
	CreateTs        int
	UpdateTs        int
	UserId          int
	Privacy         int
}

type Photorecyclebin struct {
	Id                 int
	RecordSource       string
	RecordId           int
	OpTs               time.Time
	OpUserId           int
	DelMode            string
	DelReason          string
	Processed          int
	RecyclePath        string
	Title              string
	Description        string
	Type               int
	Status             int
	Rank               int
	DisplayOrder       int
	Server             string
	FileType           string
	FileSize           int
	Storepath          string
	BigName            string
	BigWidth           int
	BigHeight          int
	MiddleName         string
	MiddleWidth        int
	MiddleHeight       int
	MiddlesquareName   string
	MiddlesquareWidth  int
	MiddlesquareHeight int
	SmallsquareName    string
	SmallsquareWidth   int
	SmallsquareHeight  int
	PrivateCtrlId      int
	PrivateSetting     string
	AlbumId            int
	CreateTs           int
	UpdateTs           int
	UserId             int
}

type Photouploadlog struct {
	UserId       int
	SessionId    string
	UploadMethod string
	PhotoName    string
	TotalNumber  int
	ErrorNumber  int
	CreateTs     int
	UploadTime   int
	UploadSize   int
}

type Myphotoindex struct {
	Id           int
	UserId       int
	AlbumId      int
	DisplayOrder int
	PhotoPrivate int
	CreateTs     int
	UpdateTs     int
	ExifTime     int
}

type SYNCTABLESDATA struct {
	ID    uint   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id"`
	NAME  string `gorm:"type:varchar(200);unique_index:UN_tablename;not null;column:name"`
	NUMID int    `gorm:"not null;column:numid"`
}

type Count struct {
	Numid int
}

var (
	dbs *gorm.DB
	dbd *gorm.DB
	err error
	//WG  sync.WaitGroup
)

var flag = true

func main() {

	db_user_S := "jiemin"
	db_pass_S := "jiemin"
	db_host_S := "10.25.1.197"
	db_port_S := 3306
	db_name_S := "babyphoto"

	dbConnectBaseStr := "%s:%s@tcp(%s:%d)/%s?charset=latin1&parseTime=True&loc=Local"
	dbConnect := fmt.Sprintf(dbConnectBaseStr, db_user_S, db_pass_S, db_host_S, db_port_S, db_name_S)
	fmt.Print(dbConnect, "\n")
	dbs, err = gorm.Open("mysql", dbConnect)
	if err != nil {
		fmt.Println(err)
	}
	defer dbs.Close()
	dbs.LogMode(true)

	db_user_D := "jiemin"
	db_pass_D := "jiemin"
	db_host_D := "10.25.1.193"
	db_port_D := 3306
	db_name_D := "babytree"
	dbConnectBaseStr_b := "%s:%s@tcp(%s:%d)/%s?charset=latin1&parseTime=True&loc=Local"
	dbConnect_b := fmt.Sprintf(dbConnectBaseStr_b, db_user_D, db_pass_D, db_host_D, db_port_D, db_name_D)
	fmt.Print(dbConnect_b, "\n")
	dbd, err = gorm.Open("mysql", dbConnect_b)
	if err != nil {
		fmt.Println(err)
	}
	defer dbd.Close()

	dbd.LogMode(true)

	var maxnum int
	num := []*Count{}

	//tablesa := []string{"MyPhotoAlbum", "MyPhoto", "MyPhotoIndex0", "MyPhotoIndex10", "MyPhotoIndex11", "MyPhotoIndex12", "MyPhotoIndex13", "MyPhotoIndex14", "MyPhotoIndex15", "MyPhotoIndex16", "MyPhotoIndex17", "MyPhotoIndex18", "MyPhotoIndex19", "MyPhotoIndex1", "MyPhotoIndex20", "MyPhotoIndex21", "MyPhotoIndex22", "MyPhotoIndex23", "MyPhotoIndex24", "MyPhotoIndex25", "MyPhotoIndex26", "MyPhotoIndex27", "MyPhotoIndex28", "MyPhotoIndex29", "MyPhotoIndex2", "MyPhotoIndex30", "MyPhotoIndex31", "MyPhotoIndex32", "MyPhotoIndex33", "MyPhotoIndex34", "MyPhotoIndex35", "MyPhotoIndex36", "MyPhotoIndex37", "MyPhotoIndex38", "MyPhotoIndex39", "MyPhotoIndex3", "MyPhotoIndex40", "MyPhotoIndex41", "MyPhotoIndex42", "MyPhotoIndex43", "MyPhotoIndex44", "MyPhotoIndex45", "MyPhotoIndex46", "MyPhotoIndex47", "MyPhotoIndex48", "MyPhotoIndex49", "MyPhotoIndex4", "MyPhotoIndex50", "MyPhotoIndex51", "MyPhotoIndex52", "MyPhotoIndex53", "MyPhotoIndex54", "MyPhotoIndex55", "MyPhotoIndex56", "MyPhotoIndex57", "MyPhotoIndex58", "MyPhotoIndex59", "MyPhotoIndex5", "MyPhotoIndex60", "MyPhotoIndex61", "MyPhotoIndex62", "MyPhotoIndex63", "MyPhotoIndex64", "MyPhotoIndex65", "MyPhotoIndex66", "MyPhotoIndex67", "MyPhotoIndex68", "MyPhotoIndex69", "MyPhotoIndex6", "MyPhotoIndex70", "MyPhotoIndex71", "MyPhotoIndex72", "MyPhotoIndex73", "MyPhotoIndex74", "MyPhotoIndex75", "MyPhotoIndex76", "MyPhotoIndex77", "MyPhotoIndex78", "MyPhotoIndex79", "MyPhotoIndex7", "MyPhotoIndex80", "MyPhotoIndex81", "MyPhotoIndex82", "MyPhotoIndex83", "MyPhotoIndex84", "MyPhotoIndex85", "MyPhotoIndex86", "MyPhotoIndex87", "MyPhotoIndex88", "MyPhotoIndex89", "MyPhotoIndex8", "MyPhotoIndex90", "MyPhotoIndex91", "MyPhotoIndex92", "MyPhotoIndex93", "MyPhotoIndex94", "MyPhotoIndex95", "MyPhotoIndex96", "MyPhotoIndex97", "MyPhotoIndex98", "MyPhotoIndex99", "MyPhotoIndex9", "PhotoRecycleBin"}
	//tablesa := []string{"MyPhoto", "MyPhotoIndex0", "MyPhotoIndex10", "MyPhotoIndex11", "MyPhotoIndex12", "MyPhotoIndex13", "MyPhotoIndex14", "MyPhotoIndex15", "MyPhotoIndex16", "MyPhotoIndex17", "MyPhotoIndex18", "MyPhotoIndex19", "MyPhotoIndex1", "MyPhotoIndex20", "MyPhotoIndex21", "MyPhotoIndex22", "MyPhotoIndex23", "MyPhotoIndex24", "MyPhotoIndex25", "MyPhotoIndex26", "MyPhotoIndex27", "MyPhotoIndex28", "MyPhotoIndex29", "MyPhotoIndex2", "MyPhotoIndex30", "MyPhotoIndex31", "MyPhotoIndex32", "MyPhotoIndex33", "MyPhotoIndex34", "MyPhotoIndex35", "MyPhotoIndex36", "MyPhotoIndex37", "MyPhotoIndex38", "MyPhotoIndex39", "MyPhotoIndex3", "MyPhotoIndex40", "MyPhotoIndex41", "MyPhotoIndex42", "MyPhotoIndex43", "MyPhotoIndex44", "MyPhotoIndex45", "MyPhotoIndex46", "MyPhotoIndex47", "MyPhotoIndex48", "MyPhotoIndex49", "MyPhotoIndex4", "MyPhotoIndex50", "MyPhotoIndex51", "MyPhotoIndex52", "MyPhotoIndex53", "MyPhotoIndex54", "MyPhotoIndex55", "MyPhotoIndex56", "MyPhotoIndex57", "MyPhotoIndex58", "MyPhotoIndex59", "MyPhotoIndex5", "MyPhotoIndex60", "MyPhotoIndex61", "MyPhotoIndex62", "MyPhotoIndex63", "MyPhotoIndex64", "MyPhotoIndex65", "MyPhotoIndex66", "MyPhotoIndex67", "MyPhotoIndex68", "MyPhotoIndex69", "MyPhotoIndex6", "MyPhotoIndex70", "MyPhotoIndex71", "MyPhotoIndex72", "MyPhotoIndex73", "MyPhotoIndex74", "MyPhotoIndex75", "MyPhotoIndex76", "MyPhotoIndex77", "MyPhotoIndex78", "MyPhotoIndex79", "MyPhotoIndex7", "MyPhotoIndex80", "MyPhotoIndex81", "MyPhotoIndex82", "MyPhotoIndex83", "MyPhotoIndex84", "MyPhotoIndex85", "MyPhotoIndex86", "MyPhotoIndex87", "MyPhotoIndex88", "MyPhotoIndex89", "MyPhotoIndex8", "MyPhotoIndex90", "MyPhotoIndex91", "MyPhotoIndex92", "MyPhotoIndex93", "MyPhotoIndex94", "MyPhotoIndex95", "MyPhotoIndex96", "MyPhotoIndex97", "MyPhotoIndex98", "MyPhotoIndex99", "MyPhotoIndex9"}

	//ch := make(chan struct{}, 1)
	// 新建计时器，1毫秒以后触发，go触发计时器的方法比较特别，就是在计时器的channel中发送值
	tick := time.NewTicker(time.Microsecond)
	for {
		select {
		// 此处在等待channel中的信号，因此执行此段代码时会阻塞1秒
		case <-tick.C:
			if flag == false {
				continue
			}

			flag = false
			//for _, v := range tablesa {
			//	//rows := dbd.Raw(fmt.Sprintf("select id from %s order by id desc limit 1", v)).Row()
			//	//rows.Scan(&maxnum)
			//	//data := SYNCTABLESDATA{}
			//	//data.NAME = v
			//	//data.NUMID = maxnum
			//	//dbs.Table("synctable").Save(&data)
			//	dbs.Table("synctable").Where("name = ?", v).Select("numid").Find(&num)
			//	for _, vv := range num {
			//		maxnum = vv.Numid
			//	}
			//
			//	switch {
			//	//case v == "MyPhotoAlbum":
			//	//	//my := []*Myphotoalbum{}
			//	//	//dbs.Table(v).Where("id >= ?", maxnum).Find(&my)
			//	//	//for _, dbdv := range my {
			//	//	//	dbd.Table(v).Save(&dbdv)
			//	//	//	dbs.Table("synctable").Where("name = ?", v).Update("numid", dbdv.Id)
			//	//	//}
			//	//	//fallthrough
			//	//
			//	//	//WG.Add(1)
			//	//	if flag {
			//	//		MyPhotoAlbum(v, maxnum)
			//	//	}
			//		//WG.Wait()
			//	case v == "MyPhoto":
			//		my := []*Myphoto{}
			//		dbs.Table(v).Where("id >= ?", maxnum).Find(&my)
			//		for _, dbdv := range my {
			//			dbd.Table(v).Save(&dbdv)
			//			dbs.Table("synctable").Where("name = ?", v).Update("numid", dbdv.Id)
			//		}
			//		//fallthrough
			//
			//		//WG.Add(1)
			//		//if flag {
			//		//	MyPhoto(v, maxnum)
			//		//}
			//		//WG.Wait()
			//
			//	//case v == "PhotoRecycleBin":
			//	//	//my := []*Photorecyclebin{}
			//	//	//dbs.Table(v).Where("id >= ?", maxnum).Find(&my)
			//	//	//for _, dbdv := range my {
			//	//	//	dbd.Table(v).Save(&dbdv)
			//	//	//	dbs.Table("synctable").Where("name = ?", v).Update("numid", dbdv.Id)
			//	//	//}
			//	//	//fallthrough
			//	//
			//	//	//WG.Add(1)
			//	//	if flag {
			//	//		PhotoRecycleBin(v, maxnum)
			//	//	}
			//	//	//WG.Wait()
			//	default:
			//		my := []*Myphotoindex{}
			//		dbs.Table(v).Where("id >= ?", maxnum).Find(&my)
			//		for _, dbdv := range my {
			//			dbd.Table(v).Save(&dbdv)
			//			dbs.Table("synctable").Where("name = ?", v).Update("numid", dbdv.Id)
			//		}
			//
			//		//WG.Add(1)
			//		//if flag {
			//		//	MyPhotoIndex(v, maxnum)
			//		//}
			//		//WG.Wait()
			//	}
			//}
			dbs.Table("synctable").Where("name = ?", "MyPhoto").Select("numid").Find(&num)
			for _, v := range num {
				maxnum = v.Numid
			}
			my := []*Myphoto{}
			dbs.Table("MyPhoto").Where("id >= ?", maxnum).Find(&my)
			for _, dbdv := range my {
				dbd.Table("MyPhoto").Save(&dbdv)
				dbs.Table("synctable").Where("name = ?", "MyPhoto").Update("numid", dbdv.Id)
			}
			flag = true
		}
	}

}

//func test() {
//	flag = false
//
//	defer func() {
//		flag = true
//	}()
//
//	fmt.Println(time.Now())
//	time.Sleep(2 * time.Second)
//}

func MyPhotoAlbum(v string, maxnum int) () {
	//defer WG.Done()
	flag = false
	defer func() {
		my := []*Myphotoalbum{}
		dbs.Table(v).Where("id >= ?", maxnum).Find(&my)
		for _, dbdv := range my {
			dbd.Table(v).Save(&dbdv)
			dbs.Table("synctable").Where("name = ?", v).Update("numid", dbdv.Id)
		}
		flag = true
	}()

}

func MyPhoto(v string, maxnum int) () {
	//defer WG.Done()
	flag = false
	defer func() {
		my := []*Myphoto{}
		dbs.Table(v).Where("id >= ?", maxnum).Find(&my)
		for _, dbdv := range my {
			dbd.Table(v).Save(&dbdv)
			dbs.Table("synctable").Where("name = ?", v).Update("numid", dbdv.Id)
		}
		flag = true
	}()

}

func PhotoRecycleBin(v string, maxnum int) () {
	//defer WG.Done()
	flag = false
	defer func() {
		my := []*Photorecyclebin{}
		dbs.Table(v).Where("id >= ?", maxnum).Find(&my)
		for _, dbdv := range my {
			dbd.Table(v).Save(&dbdv)
			dbs.Table("synctable").Where("name = ?", v).Update("numid", dbdv.Id)
		}
		flag = true
	}()

}

func MyPhotoIndex(v string, maxnum int) () {
	//defer WG.Done()
	flag = false
	defer func() {
		my := []*Myphotoindex{}
		dbs.Table(v).Where("id >= ?", maxnum).Find(&my)
		for _, dbdv := range my {
			dbd.Table(v).Save(&dbdv)
			dbs.Table("synctable").Where("name = ?", v).Update("numid", dbdv.Id)
		}
		flag = true
	}()

}
