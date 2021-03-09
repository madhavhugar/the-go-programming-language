package main

import (
	"fmt"
	"os"
	"time"
)

var aborting = make(chan struct{})

func pressAbort() {
	os.Stdin.Read(make([]byte, 2))
	aborting <- struct{}{}
}

func main() {
	ticker := time.NewTicker(time.Second)
	go pressAbort()
	for i := 10; i >= 0; i-- {
		select {
		case <-ticker.C:
			fmt.Printf("%d...\n", i)
		case <-aborting:
			fmt.Println("Aborting launch")
			os.Exit(1)
		}
	}
	ticker.Stop()
}
