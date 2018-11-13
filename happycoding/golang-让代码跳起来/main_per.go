package main

import (
	"io/ioutil"
	"log"
	"os"
	"image"
	"image/gif"
	"time"
)

func drawImg(img image.Image) {
	base := "#8XOHLTI!@%&?/-12345679eW)i=+;:,. "
	bounds := img.Bounds()
	blen := len(base)
	ret := ""
	for j := 0; j < bounds.Dy(); j += 3 {
		for i := 0; i < bounds.Dx(); i += 1 {
			color := img.At(i, j)
			r, g, b, _ := color.RGBA()
			var rr, gg, bb int
			rr = int(r >> 8)
			gg = int(g >> 8)
			bb = int(b >> 8)
			var gray float64
			gray = float64(rr)*299 + float64(gg)*578 + float64(bb)*114
			gray = gray / float64(1000)

			index := int(gray * float64(1+blen) / 255)
			if index >= blen {
				ret += " "
			} else {
				ret += string(base[index])
			}
		}
		ret += "\r\n"
	}
	ioutil.WriteFile("text.txt", []byte(ret), 0777)
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("请输入图片路径")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gs, err := gif.DecodeAll(file)
	if err != nil {
		log.Fatal(err)
	}
	for {
		for i := 0; i < len(gs.Image); i++ {
			time.Sleep(300 * time.Millisecond)
			drawImg(gs.Image[i])
		}
	}
}
