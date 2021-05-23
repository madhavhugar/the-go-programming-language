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