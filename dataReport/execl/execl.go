package execl

import (
	"log"
	//"fmt"
	"github.com/xuri/excelize"
)

func Execl_install(excelPath string, dict map[string]string) {
	xlsx := excelize.NewFile()
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
	xlsx_r.SetCellValue("Sheet1", "A1", "任务ID")
	xlsx_r.SetCellValue("Sheet1", "B1", "任务名")
	xlsx_r.SetCellValue("Sheet1", "C1", "链接")
	xlsx_r.SetCellValue("Sheet1", "D1", "创建者")
	xlsx_r.SetCellValue("Sheet1", "E1", "创建时间")
	xlsx_r.SetCellValue("Sheet1", "F1", "页面数")
	xlsx_r.SetCellValue("Sheet1", "G1", "采集数")
	xlsx_r.SetCellValue("Sheet1", "H1", "入库数")
	xlsx_r.SetCellValue("Sheet1", "I1", "累计入库数")

	err = xlsx_r.SaveAs(excelPath)
	if err != nil {
		log.Fatalln(err)
	}
}
