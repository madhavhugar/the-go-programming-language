package main

import (
	"testing"
	"time"
)

func TestDeposit(t *testing.T) {
	go Deposit(100)
	go Deposit(200)

	<- time.After(time.Second * 1)

	if b := Balance(); b != 300 {
		t.Errorf("Balance != 300; %f != 300", b)
	}
}

func TestWithdraw(t *testing.T) {
	var success bool
	go Deposit(200)
	go func () {
		success = Withdraw(150)
	}()

	<- time.After(time.Second * 1)

	if b := Balance(); b != 50 {
		t.Errorf("Balance != 50; %f != 50", b)
	}
	if !success {
		t.Errorf("Unsuccessful withdrawal from surplus")
	}
}