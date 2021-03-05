package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func crawl5(url string) []string {
	fmt.Println("Extracting from", url)
	links, err := Extract(url)
	if err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "error"))
	}
	return links
}

var worklistu = make(chan string)
var linkslist = make(chan []string)

func createWorkers() {
	for i := 0; i < 10; i++ {
		go func() {
			for url := range worklistu {
				links := crawl5(url)
				// implemented Option A here;
				// see crawler4 for Option B implementation;
				go func() {
					linkslist <- links
				}()
			}
		}()
	}
}

func manageWorkers() {
	go func() { worklistu <- os.Args[1] }()

	for links := range linkslist {
		for _, link := range links {
			worklistu <- link
		}
	}
}

func main() {
	createWorkers()
	manageWorkers()
}
