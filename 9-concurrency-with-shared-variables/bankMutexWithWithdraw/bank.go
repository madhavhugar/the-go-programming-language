package main

import "sync"

var (
	balance float64
	mu sync.Mutex
)

// Ensure this function is not called directly
// to avoid race condition
func deposit(amount float64) {
	balance += amount
}

func Deposit(amount float64) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() float64 {
	mu.Lock()
	defer mu.Unlock()
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