package main

import (
	"fmt"
	"os"
)

// we want to ensure we fetch them concurrently v2

func crawl2(url string) []string {
	fmt.Println("crawling link", url)
	links, err := Extract(url)
	if err != nil {
		fmt.Println(err)
	}
	return links
}

func linkFetcher() {
	worklist := make(chan []string)

	go func() {
		worklist <- []string{os.Args[1]}
	}()

	for list := range worklist {
		for _, url := range list {
			go func(u string) {
				worklist <- crawl2(u)
			}(url)
		}
	}
}

func main() {
	linkFetcher()
}
