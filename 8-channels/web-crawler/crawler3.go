package main

import (
	"fmt"
	"os"
)

// how to avoid unbounded parallelism
// we create a buffered channel, that behaves
// like a bouncer for only certain number of
// goroutines to execute at a given time

var bouncer = make(chan struct{}, 20)

func crawl3(url string) []string {
	fmt.Println("extracting url", url)
	bouncer <- struct{}{}
	links, err := Extract(url)
	<-bouncer
	if err != nil {
		fmt.Println(err)
	}
	return links
}

func worker() {
	worklist := make(chan []string)

	go func() {
		worklist <- []string{os.Args[1]}
	}()

	for list := range worklist {
		for _, url := range list {
			go func(u string) {
				worklist <- crawl3(u)
			}(url)
		}
	}
}

func main() {
	worker()
}
