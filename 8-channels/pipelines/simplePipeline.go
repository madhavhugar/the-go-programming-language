package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squarer := make(chan int)

	// natural number generator
	go func() {
		defer close(naturals)
		for i := 0; i < 5; i++ {
			naturals <- i
			time.Sleep(time.Second * 1)
		}
	}()

	// number squarer
	go func() {
		defer close(naturals)
		for {
			n := <-naturals
			squarer <- n * n
		}
	}()

	// number printer
	for {
		fmt.Println(<-squarer)
	}
}
