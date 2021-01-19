package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 2*time.Second, "sleep duration")

func main() {
	flag.Parse()

	fmt.Printf("Sleeping for a given duration %v...\n", *period)
	time.Sleep(*period)
	fmt.Println("Waking up now...")
}
