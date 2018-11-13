package main

import (
	"github.com/hajimehoshi/oto"
	"github.com/nsf/termbox-go"
	"github.com/tosone/minimp3"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"log"
)

func initialize() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
}

func Sonud(filename string) {
	if len(filename) == 0 {
		return
	}

	file, err := ioutil.ReadFile("mp3/" + filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	dec, data1, _ := minimp3.DecodeFull(file)
	player, _ := oto.NewPlayer(
		dec.SampleRate,
		dec.Channels,
		2,
		1024,
	)
	player.Write(data1)
	player.Close()
}

func main_bak() {
	files, _ := ioutil.ReadDir("mp3")
	for _, file := range files {
		if err := os.Rename(
			"mp3/"+file.Name(),
			"mp3/"+strings.Replace(file.Name(), "", "", -1),
		); err != nil {
			log.Fatal(err)
		}
		fmt.Println(file.Name())
	}
}

func main() {
	files, _ := ioutil.ReadDir("mp3")
	lenf := len(files)
	fmt.Println(lenf)

LOOP:
	for {
		switch ev := termbox.PollEvent(); ev.Key {
		case termbox.KeyEsc:
			break LOOP
		case termbox.KeyEnter:
			fmt.Println()
		case termbox.KeySpace:
			fmt.Print(" ")
		case termbox.KeyTab:
			fmt.Print("     ")
		default:
			fmt.Printf("%s", string(ev.Ch))
			i := int(ev.Ch) % lenf
			go Sonud(files[i].Name())
		}
	}
}
