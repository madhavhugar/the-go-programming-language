package main

import (
	"testing"
	"time"
)

func TestDeposit(t *testing.T) {
	go Deposit(100)
	go Deposit(200)

	<- time.After(time.Second * 3)

	if Balance() != 300 {
		t.Errorf("Balance != 300; %f != 300", Balance())
	}
}