package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu sync.Mutex
	balance float64
)

func Deposit(amount float64) {
	mu.Lock()
	balance += amount
	defer mu.Unlock()
	return
}

func Balance() float64 {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func main() {
	go Deposit(100)
	go fmt.Println(Balance())
	go Deposit(250)
	go fmt.Println(Balance())

	<- time.After(time.Second * 5)
	fmt.Println(Balance())
}

