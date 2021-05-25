package main

var balances = make(chan float64)
var deposit = make(chan float64)

func Deposit(amount float64) {
	deposit <- amount
}

func Balance() float64 {
	b := <- balances
	return b
}

func Withdraw(amount float64) bool {
	deposit <- -amount
	if b := <- balances; b < 0 {
		deposit <- amount
		return false
	}
	return true
}

func teller() {
	var balance float64
	for {
		select {
		case a := <- deposit:
			balance += a
		case balances <- balance:
		}
	}
}