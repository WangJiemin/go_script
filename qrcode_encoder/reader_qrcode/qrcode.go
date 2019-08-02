package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
)

const (
	QR_CODE_SIZE        = 256
	SHRINK_QR_CODE_SIZE = 35
	MARGIN              = 29
	MULTIPLE		 	= 6
)

//const (
//	QR_CODE_SIZE        = 301
//	SHRINK_QR_CODE_SIZE = 39
//	MARGIN              = 21
//	MULTIPLE            = 7
//)

type QRCode struct {
	text            string
	out             string
	img             image.Image
	genImg          bool
	points          [QR_CODE_SIZE][QR_CODE_SIZE]int
	tmpShrinkPoints [QR_CODE_SIZE][SHRINK_QR_CODE_SIZE]int
	shrinkPoints    [SHRINK_QR_CODE_SIZE][SHRINK_QR_CODE_SIZE]int
}

func (this *QRCode) ParseCmdOptions() {
	flag.StringVar(&this.text, "t", "", "二维图内容")
	flag.StringVar(&this.out, "o", "", "输出文件")
	flag.Parse()
}

//NewQRCode 返回二维码结构
func NewQRCode(text string, out string, genImg bool) *QRCode {
	qr := &QRCode{
		text:   text,
		genImg: genImg,
	}

	qr.genQRCode(text, out)

	return qr
}

//genQRCode 生成二维码
func (qr *QRCode) genQRCode(text string, out string) error {
	code, err := qrcode.Encode(text, qrcode.Medium, QR_CODE_SIZE)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(code)
	img, err := png.Decode(buf)
	if err != nil {
		return err
	}

	qr.img = img

	if qr.genImg {
		newPng, _ := os.Create(out)
		defer newPng.Close()
		png.Encode(newPng, img)
	}

	return nil
}

func (qr *QRCode) SetImage(img image.Image) error {
	qr.img = img

	if qr.genImg {
		newPng, _ := os.Create(qr.out)
		defer newPng.Close()
		png.Encode(newPng, img)
	}

	return nil
}

//binarization 二维码图片二值化 0－1
func (qr *QRCode) binarization() {
	gray := image.NewGray(image.Rect(0, 0, QR_CODE_SIZE, QR_CODE_SIZE))
	for x := 0; x < QR_CODE_SIZE; x++ {
		for y := 0; y < QR_CODE_SIZE; y++ {
			r32, g32, b32, _ := qr.img.At(x, y).RGBA()
			r, g, b := int(r32>>8), int(g32>>8), int(b32>>8)
			if (r+g+b)/3 > 180 {
				qr.points[y][x] = 0
				gray.Set(x, y, color.Gray{uint8(255)})
			} else {
				qr.points[y][x] = 1
				gray.Set(x, y, color.Gray{uint8(0)})
			}
		}
	}

	if qr.genImg {
		newPng, _ := os.Create(qr.out + ".qrcode")
		defer newPng.Close()
		png.Encode(newPng, gray)
	}
}

//shrink 缩小二值化数组
func (qr *QRCode) shrink() {
	for x := 0; x < QR_CODE_SIZE; x++ {
		cal := 1
		for y := MARGIN + 1; y < QR_CODE_SIZE-MARGIN; y += MULTIPLE {
			qr.tmpShrinkPoints[x][cal] = qr.points[x][y]
			cal++
		}
	}

	for y := 1; y < SHRINK_QR_CODE_SIZE-1; y++ {
		row := 1
		for x := MARGIN + 1; x < QR_CODE_SIZE-MARGIN; x += MULTIPLE {
			qr.shrinkPoints[row][y] = qr.tmpShrinkPoints[x][y]
			row++
		}
	}
}

//Output 控制台输出二维码
func (qr *QRCode) Output() {
	qr.binarization()
	qr.shrink()

	for x := 0; x < SHRINK_QR_CODE_SIZE; x++ {
		for y := 0; y < SHRINK_QR_CODE_SIZE; y++ {
			if qr.shrinkPoints[x][y] == 1 {
				//fmt.Print("\033[40;40m  \033[0m")
				randColor()
			} else {
				fmt.Print("\033[47;30m  \033[0m")
			}
		}
		fmt.Println()
	}
}

func randColor() {
	cls := []int{41, 42, 43, 44, 45, 46}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(1)
	fmt.Printf("\033[40;%dm  \033[0m", cls[i])
}

//Debug 调试二维码二值化及缩小过程
func (qr *QRCode) Debug() {
	qr.binarization()
	//qr.shrink()

	src, _ := os.Create("src.txt")
	for i := 0; i < len(qr.points); i++ {
		var line string
		for j := 0; j < len(qr.points[i]); j++ {
			if qr.points[i][j] == 1 {
				line += "1"
			} else {
				line += "0"
			}
		}
		line += "\n"
		src.WriteString(line)
	}
	src.Close()

	tmp, _ := os.Create("tmp.txt")
	for i := 0; i < len(qr.tmpShrinkPoints); i++ {
		var line string
		for j := 0; j < len(qr.tmpShrinkPoints[i]); j++ {
			if qr.tmpShrinkPoints[i][j] == 1 {
				line += "1"
			} else {
				line += "0"
			}
		}
		line += "\n"
		tmp.WriteString(line)
	}
	tmp.Close()

	dst, _ := os.Create("dst.txt")
	for i := 0; i < len(qr.shrinkPoints); i++ {
		var line string
		for j := 0; j < len(qr.shrinkPoints[i]); j++ {
			if qr.shrinkPoints[i][j] == 1 {
				line += "1"
			} else {
				line += "0"
			}
		}
		line += "\n"
		dst.WriteString(line)
	}
	dst.Close()
}

func main() {
	var conf *QRCode = &QRCode{}
	conf.ParseCmdOptions()
	flags := false

	if conf.text == "" {
		logrus.WithFields(logrus.Fields{
			"qrcode": false,
			"number": 100,
		}).Fatal("请指定二维码的生成内容")
	} else if conf.out == "" {
		logrus.WithFields(logrus.Fields{
			"qrcode": false,
			"number": 99,
		}).Fatalln("请指定输出文件")
	} else {
		flags = true
	}

	// if len(os.Args[1]) > 60 {
	// 	fmt.Println("\033[31mERR: max context length is 60.\033[0m")
	// 	return
	// }
	if flags {
		qr := NewQRCode(conf.text, conf.out, flags)
		logrus.WithFields(logrus.Fields{"grcpde": flags, "number": 0}).Infoln("二维码")
		qr.Output()

		os.Remove("src.txt")
		os.Remove("tmp.txt")
		os.Remove("dst.txt")
		os.Remove(qr.out + ".qrcode")
	} else {
		flag.Usage()
	}
}
