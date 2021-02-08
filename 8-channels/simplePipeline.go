package main

import "fmt"

func main() {
	naturals := make(chan int)
	squarer := make(chan int)

	// natural number generator
	go func() {
		for i := 0; i < 100; i++ {
			naturals <- i
		}
	}()

	// number squarer
	go func() {
		n := <-naturals
		squarer <- n * n
	}()

	// number printer
	s := <-squarer
	fmt.Println(s)
}
