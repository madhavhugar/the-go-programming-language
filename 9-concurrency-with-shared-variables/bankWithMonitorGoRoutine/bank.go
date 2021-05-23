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