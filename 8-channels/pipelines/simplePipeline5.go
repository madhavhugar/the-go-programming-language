package main

import (
	"fmt"
)

func first(ch chan int) {
	defer func() {
		close(ch)
	}()
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func second(input chan int, output chan int) {
	defer func() {
		close(output)
	}()
	for num := range input {
		output <- num * num
	}
}

func third(ch chan int) {
	for n := range ch {
		fmt.Println(n)
	}
}

func main() {
	nums := make(chan int)
	sqs := make(chan int)
	go first(nums)
	go second(nums, sqs)
	third(sqs)
}
