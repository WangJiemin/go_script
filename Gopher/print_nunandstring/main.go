package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	chan_n := make(chan bool)
	chan_c := make(chan bool, 1)
	done := make(chan struct{})

	go func() {
		for i := 0; i < 11; i += 2 {
			<-chan_c
			fmt.Print(i)
			fmt.Print(i + 1)
			chan_n <- true
		}
	}()

	go func() {
		char_req := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
		for i := 0; i < 10; i += 2 {
			<-chan_n
			fmt.Print(char_req[i])
			fmt.Print(char_req[i+1])
			chan_c <- true
		}

		done <- struct{}{}
	}()

	chan_c <- true
	<-done
}
