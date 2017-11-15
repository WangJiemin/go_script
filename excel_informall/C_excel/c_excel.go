package C_excel

import (
	"excel_informall/compress"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xuri/excelize"
)

func CreateExcel(excelPath string, dict map[string]string) {
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
	xlsx_r.SetCellValue("Sheet1", "A1", "链接")
	xlsx_r.SetCellValue("Sheet1", "B1", "标题")
	xlsx_r.SetCellValue("Sheet1", "C1", "作者")
	xlsx_r.SetCellValue("Sheet1", "D1", "来源")
	xlsx_r.SetCellValue("Sheet1", "E1", "发布时间")
	xlsx_r.SetCellValue("Sheet1", "F1", "正文")
	xlsx_r.SetCellValue("Sheet1", "G1", "媒体类型")
	xlsx_r.SetCellValue("Sheet1", "H1", "原标题")
	xlsx_r.SetCellValue("Sheet1", "I1", "责任编辑")
	xlsx_r.SetCellValue("Sheet1", "J1", "记者")
	xlsx_r.SetCellValue("Sheet1", "K1", "文章评论数")
	xlsx_r.SetCellValue("Sheet1", "L1", "文章阅读次数")
	xlsx_r.SetCellValue("Sheet1", "M1", "文章摘要")
	xlsx_r.SetCellValue("Sheet1", "N1", "媒体名称")
	xlsx_r.SetCellValue("Sheet1", "O1", "分类（媒体频道）")
	xlsx_r.SetCellValue("Sheet1", "P1", "页面位置")
	xlsx_r.SetCellValue("Sheet1", "Q1", "位置路径")
	xlsx_r.SetCellValue("Sheet1", "R1", "采集工具")
	xlsx_r.SetCellValue("Sheet1", "S1", "项目客户")
	xlsx_r.SetCellValue("Sheet1", "T1", "采集站点网址")
	xlsx_r.SetCellValue("Sheet1", "U1", "获取时间")

	err = xlsx_r.SaveAs(excelPath)
	if err != nil {
		log.Fatalln(err)
	}
	ZIP(excelPath)
}

func ZIP(excelPath string) {
	f1, err := os.Open(fmt.Sprintf("%s", excelPath))
	if err != nil {
		log.Fatalln(err)
	}
	defer f1.Close()
	var files = []*os.File{f1}
	dest := strings.Replace(excelPath, "xlsx", "zip", -1)
	compress.ComZIP(files, dest)
	os.Remove(excelPath)
}
