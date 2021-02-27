package main

import (
	"flag"
	"fmt"
)

var root string

func init() {
	flag.StringVar(&root, "root", "", "--root:\troot link to begin crawling")
	flag.Parse()
}

func crawl(url string) []string {
	fmt.Println("Extracting from ", url, len(url))
	list, err := Extract(url)
	// fmt.Println("Got", list, "from", url)
	if err != nil {
		return []string{}
	}
	return list
}

func manager(rootLink string) {
	links := make(chan []string)

	go func() {
		links <- []string{rootLink}
	}()

	seen := make(map[string]bool)
	for childLinks := range links {
		for _, url := range childLinks {
			_, ok := seen[url]
			fmt.Println("url", url, "ok", ok)
			if !ok {
				seen[url] = true
				go func(u string) {
					links <- crawl(url)
				}(url)
			}
		}
	}
}

func main() {
	manager(root)
}
