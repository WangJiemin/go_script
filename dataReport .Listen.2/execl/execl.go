package execl

import (
	"fmt"
	"funbird-dataReport/gettime"
	"log"
	"time"

	"github.com/xuri/excelize"
)

func Execl_install(excelPath string, dict map[string]string) {
	now := time.Now().Format("20060102")
	d1, d2 := gettime.TWODAYS()

	xlsx := excelize.NewFile()
	//xlsx.NewSheet()
	for k, v := range dict {
		//fmt.Println(k, v)
		xlsx.SetCellValue("Sheet1", k, v)
	}
	err := xlsx.SaveAs(excelPath)
	if err != nil {
		log.Fatalln(err)

	}

	xlsx_r, err := excelize.OpenFile(excelPath)
	if err != nil {
		log.Fatalln(err)
	}
	xlsx_r.InsertRow("Sheet1", 0)
	xlsx_r.SetCellValue(fmt.Sprintf("%s", "Sheet1"), fmt.Sprintf("%s", "A1"), fmt.Sprintf("%s", "任务ID"))
	xlsx_r.SetCellValue(fmt.Sprintf("%s", "Sheet1"), fmt.Sprintf("%s", "B1"), fmt.Sprintf("%s", "任务名"))
	xlsx_r.SetCellValue(fmt.Sprintf("%s", "Sheet1"), fmt.Sprintf("%s", "C1"), fmt.Sprintf("%s", "链接"))
	xlsx_r.SetCellValue(fmt.Sprintf("%s", "Sheet1"), fmt.Sprintf("%s", "D1"), fmt.Sprintf("%s", "创建者"))
	xlsx_r.SetCellValue(fmt.Sprintf("%s", "Sheet1"), fmt.Sprintf("%s", "E1"), fmt.Sprintf("%s", "创建时间"))
	xlsx_r.SetCellValue(fmt.Sprintf("%s", "Sheet1"), fmt.Sprintf("%s", "F1"), fmt.Sprintf("%s%s", now, "入库数"))
	xlsx_r.SetCellValue(fmt.Sprintf("%s", "Sheet1"), fmt.Sprintf("%s", "G1"), fmt.Sprintf("%s%s", d1, "入库数"))
	xlsx_r.SetCellValue(fmt.Sprintf("%s", "Sheet1"), fmt.Sprintf("%s", "H1"), fmt.Sprintf("%s%s", d2, "入库数"))
	//xlsx_r.SetCellValue(fmt.Sprintf("%s,%s,%d%s", "Sheet1", "I1", "累计入库数"))

	err = xlsx_r.SaveAs(excelPath)
	if err != nil {
		log.Fatalln(err)
	}
}
