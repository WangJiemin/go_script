package main

import (
	"image"
	"image/color"
	"image/gif"
	"io/ioutil"
	"log"
	"os"
	"time"
)

//图片显示的内容
const base = "#8XOHLTI!@%&?/-12345679eW)i=+;:,. "

//话图片
func drawImage(img image.Image) string {
	ret := ""
	blen := len(base)
	bounds := img.Bounds()
	for j := 0; j < bounds.Dy(); j += 2 {
		for i := 0; i < bounds.Dx(); i++ {
			grey := getGrey(img.At(i, j))
			index := int(float64(blen+1) * grey / 255)
			if index >= blen {
				ret += " "
			} else {
				ret += string(base[index])
			}
		}
		ret += "\r\n"
	}
	return ret
}

//灰度算法 -- 可以自行百度
func getGrey(color color.Color) float64 {
	r, g, b, _ := color.RGBA()
	var rr, gg, bb int
	rr = int(r >> 8)
	gg = int(g >> 8)
	bb = int(b >> 8)
	var gray float64
	gray = float64(rr)*299 + float64(gg)*578 + float64(bb)*114
	gray = gray / float64(1000)
	return gray
}

//将gif图片转化为字符动起来
func main() {
	if len(os.Args) <= 1 {
		log.Fatal("请输入图片路径")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//得到gif图片
	gs, err := gif.DecodeAll(file)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for i := 1; i < len(gs.Image); i++ {
			//定时跑，
			time.Sleep(300 * time.Millisecond)
			str := drawImage(gs.Image[i])
			ioutil.WriteFile("text.txt", []byte(str), 0777)
		}
	}
}
