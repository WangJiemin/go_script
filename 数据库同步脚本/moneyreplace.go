package main

import (
	//"flag"
	"fmt"
	"regexp"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ConvertNumToCny(num float64) string {
	strnum := strconv.FormatFloat(num*100, 'f', 0, 64)
	sliceUnit := []string{"仟", "佰", "拾", "亿", "仟", "佰", "拾", "万", "仟", "佰", "拾", "元", "角", "分"}

	//	log.Println(sliceUnit[:len(sliceUnit)-2])
	s := sliceUnit[len(sliceUnit)-len(strnum) : len(sliceUnit)]
	//upperDigitUnit := map[int]string{0: "零", 1: "壹", 2: "贰", 3: "叁", 4: "肆", 5: "伍", 6: "陆", 7: "柒", 8: "捌", 9: "玖"}
	upperDigitUnit := map[string]string{"0": "零", "1": "壹", "2": "贰", "3": "叁", "4": "肆", "5": "伍", "6": "陆", "7": "柒", "8": "捌", "9": "玖"}

	str := ""
	for k, v := range strnum[:] {
		str = str + upperDigitUnit[string(v)] + s[k]
	}

	reg, err := regexp.Compile(`零角零分$`)
	checkErr(err)
	str = reg.ReplaceAllString(str, "整")

	reg, err = regexp.Compile(`零角$`)
	checkErr(err)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零分`)
	checkErr(err)
	str = reg.ReplaceAllString(str, "整")

	reg, err = regexp.Compile(`零[仟佰拾]`)
	checkErr(err)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零{2,}`)
	checkErr(err)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零亿`)
	checkErr(err)
	str = reg.ReplaceAllString(str, "亿")

	reg, err = regexp.Compile(`零万`)
	checkErr(err)
	str = reg.ReplaceAllString(str, "万")

	reg, err = regexp.Compile(`亿零{0, 3}万`)
	checkErr(err)
	str = reg.ReplaceAllString(str, "^元")

	reg, err = regexp.Compile(`零*元`)
	checkErr(err)
	str = reg.ReplaceAllString(str, "元")

	reg, err = regexp.Compile(`零元`)
	checkErr(err)
	str = reg.ReplaceAllString(str, "零")

	return str
}

func main() {
	//	// flag.Args方式
	//	flag.Parse()
	//	var inputnum []string = flag.Args()
	//	if inputnum != nil && len(inputnum) > 0 {
	//		fmt.Printf("%s\t的金额转换为:\t", inputnum[0], ConvertNumToCny(inputnum[0])
	//	} else {
	//		fmt.Println("请输入要转换的金额")
	//	}
	// ---------------------------------------------------------------------------------------------
	//	var inputnum float64
	//	var i float64
	//	fmt.Print("Please enter your full money:", "\t")
	//	fmt.Scanln(&inputnum)
	//	fmt.Printf("%s\t的大写金额为:\t%s", inputnum, ConvertNumToCny(inputnum))
	// ---------------------------------------------------------------------------------------------
	fmt.Printf("123456789.45的大写金额为: %s\n", ConvertNumToCny(123456789.45))
}
