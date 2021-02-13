package main

import "fmt"

func main() {
	naturals := make(chan int)
	squarer := make(chan int)

	go func() {
		defer close(naturals)
		for i := 0; i < 10; i++ {
			naturals <- i
		}
	}()

	go func() {
		defer close(squarer)
		for {
			n, ok := <-naturals
			if !ok {
				break
			}
			squarer <- n * n
		}
	}()

	for {
		p, ok := <-squarer
		if !ok {
			break
		}
		fmt.Println(p)
	}
}
