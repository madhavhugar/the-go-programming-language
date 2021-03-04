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

var worklist = make(chan string)
var resultlist = make(chan []string)

func workers() {
	go func() {
		worklist <- os.Args[1]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for url := range worklist {
				unseenLinks := crawl4(url)
				// Option A: ensure that sending to resultlist is not blocking
				// by performing the send operation inside a goroutine
				// which is commented here
				// go func() {
				resultlist <- unseenLinks
				// }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range resultlist {
		for _, url := range list {
			if !seen[url] {
				seen[url] = true
				// Option B: ensure that sending to worklist is not blocking
				// by performing the send operation inside a goroutine
				go func() {
					worklist <- url
				}()
			}
		}
	}
}

func main() {
	workers()
}
