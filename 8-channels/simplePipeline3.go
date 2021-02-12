package main

import "fmt"

func main() {
	naturals := make(chan int)
	squarer := make(chan int)

	go func() {
		defer close(naturals)
		for i := 0; i < 100; i++ {
			naturals <- i
		}
	}()

	go func() {
		defer close(squarer)
		for n := range naturals {
			squarer <- n * n
		}
	}()

	for p := range squarer {
		fmt.Println(p)
	}
}
