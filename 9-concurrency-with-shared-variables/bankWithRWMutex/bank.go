package main

import "sync"

var (
	balance float64
	mu sync.RWMutex
)

func deposit(amount float64) {
	balance += amount
}

func Deposit(amount float64) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() float64 {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

func Withdraw(amount float64) bool {
	mu.Lock()
	defer mu.Unlock()

	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}

	return true
}