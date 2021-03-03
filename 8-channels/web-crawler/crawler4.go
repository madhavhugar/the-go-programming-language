package main

import "fmt"

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

func workers() {
	for i := 0; i < 20; i++ {
		go func() {
			url := <-worklist
			<-crawl4(url)
		}()
	}
}
