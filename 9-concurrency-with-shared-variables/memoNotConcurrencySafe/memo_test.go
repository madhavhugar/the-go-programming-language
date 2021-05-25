package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMemo(t *testing.T) {
	urls := []string{
		"https://google.com",
		"https://golang.org",
		"https://godoc.org",
		"https://google.com",
		"https://golang.org",
		"https://godoc.org",
	}

	t.Run("non concurrent version", func(t *testing.T) {
		memo := New()
		for _, u := range urls {
			start := time.Now()
			memo.Get(u)
			fmt.Println("Time since: ", time.Since(start))
		}
	})

	t.Run("concurrent version", func(t *testing.T) {
		memo := New()
		for _, u := range urls {
			go func (url string) {
				start := time.Now()
				memo.Get(url)
				fmt.Println("URL", url, "Time taken", time.Since(start))
			}(u)
		}

		<- time.After(time.Second * 2)
	})
}
