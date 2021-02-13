package main

import "fmt"

func naturals(nums chan<- int) {
	defer close(nums)
	for i := 0; i < 100; i++ {
		nums <- i
	}
}

func squarer(nums <-chan int, squares chan<- int) {
	defer close(squares)
	for n := range nums {
		squares <- n * n
	}
}

func printer(squares <-chan int) {
	for p := range squares {
		fmt.Println(p)
	}
}

func main() {
	nums := make(chan int)
	squares := make(chan int)
	go naturals(nums)
	go squarer(nums, squares)
	printer(squares)
}
