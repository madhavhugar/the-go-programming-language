package main

import (
	"fmt"
	"os"
)

// we avoid unbounded parallelism
// by creating a predefined number
// of goroutines who would be designated workers

func crawl4(url string) []string {
	fmt.Println("extracting from ", url)
	links, err := Extract(url)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return links
}

var worklist = make(chan string, 20)
var resultlist = make(chan []string)

func workers() {
	go func() {
		worklist <- os.Args[1]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for url := range worklist {
				resultlist <- crawl4(url)
			}
		}()
	}

	for list := range resultlist {
		for _, url := range list {
			worklist <- url
		}
	}
}

func main() {
	workers()
}
