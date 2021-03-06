package main

import (
	"testing"
	"time"
)

func TestMonitorGoRoutineBank(t *testing.T) {
	go teller()
	<- time.After(time.Second * 1)

	go Deposit(100)
	go Deposit(200)


	if Balance() != 300 {
		t.Errorf("Balance != 300; %f != 300", Balance())
	}
}

func TestWithdraw(t *testing.T) {
	var success bool
	go teller()

	<- time.After(time.Second * 1)

	go Deposit(100)
	go func () {
		success = Withdraw(50)
	}()

	if Balance() != 50 {
		t.Errorf("Balance != 50; %f != 50", Balance())
	}
	if !success {
		t.Error("Unsuccessful withdrawal")
	}
}