package main

import (
	"fmt"
	"os"
	"time"
)

func launcher() {
	fmt.Println("Launching SLV")
	time.Sleep(time.Second * 2)
}

var abort = make(chan struct{})

func aborter() {
	os.Stdin.Read(make([]byte, 1))
	abort <- struct{}{}
}

func main() {
	countdown := time.Tick(time.Second)
	go aborter()
	for i := 10; i >= 0; i-- {
		select {
		case <-countdown:
			fmt.Printf("%d...\n", i)
		case <-abort:
			fmt.Println("Halt called. Aborting")
			os.Exit(1)
		}
	}
	launcher()
}
