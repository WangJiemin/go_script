package main

import (
	"flag"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
	//unqrcode "github.com/tuotoo/qrcode"
)

type Config struct {
	text    string
	logo    string
	percent int
	size    int
	out     string
}

func (this *Config) ParseCmdOptions() {
	flag.StringVar(&this.text, "t", "", "二维图内容")
	flag.StringVar(&this.logo, "l", "", "二维图Logo(png)")
	flag.IntVar(&this.percent, "p", 15, "二维码Logo的显示比例(默认15%)")
	flag.IntVar(&this.size, "s", 256, "二维码的大小(默认256)")
	flag.StringVar(&this.out, "o", "", "输出文件")
	flag.Parse()
}

func checkFile(name string) (bool, error) {
	_, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, nil
	}
	return true, nil
}

func resizeLogo(logo string, size uint) (image.Image, error) {
	file, err := os.Open(logo)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	img = resize.Resize(size, size, img, resize.Lanczos3)
	return img, nil
}

func addLogo(srcImage image.Image, logo string, size int) (image.Image, error) {
	logoImage, err := resizeLogo(logo, uint(size))
	if err != nil {
		return nil, err
	}

	offset := image.Pt((srcImage.Bounds().Dx()-logoImage.Bounds().Dx())/2, (srcImage.Bounds().Dy()-logoImage.Bounds().Dy())/2)
	b := srcImage.Bounds()
	m := image.NewNRGBA(b)
	draw.Draw(m, b, srcImage, image.ZP, draw.Src)
	draw.Draw(m, logoImage.Bounds().Add(offset), logoImage, image.ZP, draw.Over)

	return m, nil
}

func main() {
	//flag.Parse()

	var conf *Config = &Config{}
	conf.ParseCmdOptions()

	if conf.text == "" {
		logrus.WithFields(logrus.Fields{
			"qrcode": false,
			"number": 100,
		}).Fatal("请指定二维码的生成内容")
	}

	if conf.out == "" {
		logrus.WithFields(logrus.Fields{
			"qrcode": false,
			"number": 99,
		}).Fatalln("请指定输出文件")
	}

	if exists, err := checkFile(conf.out); err != nil {
		logrus.WithFields(logrus.Fields{
			"qrcode": false,
			"number": 98,
		}).Fatalf("检查输出文件发生错误: %s", err.Error())
	} else if exists {
		logrus.WithFields(logrus.Fields{
			"qrcode": false,
			"number": 97,
		}).Fatalln("输出文件已经存在，请重新指定")
	}

	code, err := qrcode.New(conf.text, qrcode.Highest)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"qrcode": false,
			"number": 90,
		}).Fatalf("创建二维码发生错误: %s", err.Error())
	}

	srImage := code.Image(conf.size)
	if conf.logo != "" {
		logoSize := float64(conf.size) * float64(conf.percent) / 100
		srImage, err = addLogo(srImage, conf.logo, int(logoSize))
		if err != nil {
			logrus.WithFields(logrus.Fields{"qrcode": false, "number": 98}).Fatalf("增加Logo发生错误: %s", err.Error())
		}
	}

	outAbs, err := filepath.Abs(conf.out)
	if err != nil {
		logrus.WithFields(logrus.Fields{"qrcode": false, "number": 98}).Fatalf("获取输出文件绝对路径发生错误: %s", err.Error())
	}

	os.MkdirAll(filepath.Dir(outAbs), 0777)
	outFile, err := os.Create(outAbs)
	if err != nil {
		logrus.WithFields(logrus.Fields{"qrcode": false, "number": 98}).Fatalf("创建输出文件发生错误：%s", err.Error())
	}
	defer outFile.Close()

	jpeg.Encode(outFile, srImage, &jpeg.Options{Quality: 100})

	logrus.WithFields(logrus.Fields{"grcpde": true, "number": 0}).Infof("二维码生成成功，文件路径：%s", outAbs)

	// 读取打印二维码文件
	//f, err := os.Open(outAbs)
	//if err != nil {
	//	logrus.WithFields(logrus.Fields{"grcpde": false, "number": 80}).Fatalf("二维码文件不存在，文件路径：%s", outAbs)
	//}
	//defer f.Close()
	//
	//if qrmatrix, err := unqrcode.Decode(f); err != nil {
	//	logrus.WithFields(logrus.Fields{"qrcode": false, "number": 70}).Fatalf("读取二维码图片错误：%s", err.Error())
	//} else {
	//	fmt.Println(qrmatrix)
	//}

}
