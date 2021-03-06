package main

import (
	"fmt"
	"time"
)

func launchVehicle() {
	fmt.Println("Launching SLV")
	time.Sleep(time.Second * 2)
}

func main() {
	countdown := time.Tick(time.Second * 2)
	count := 10
	for i := count; i >= 0; i-- {
		<-countdown
		fmt.Printf("%d...\n", i)
	}
	launchVehicle()
}
