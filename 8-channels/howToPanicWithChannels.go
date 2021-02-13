package channels

import "fmt"

func panicOnClosingClosedChannel() {
	c := make(chan int)
	close(c)
	close(c)
}

func panicOnSendToClosedChannel() {
	// create an unbuffered channel of type bool
	c := make(chan bool)

	// conduit some bool through channel c
	go receiveFromChannel(c)
	go sendToChannel(c)
	// close channel c
	close(c)

	// send on a closed channel to cause panic
	sendToChannel(c)
}

func receiveFromChannel(c chan bool) {
	r := <-c
	fmt.Println("received: ", r)
}

func sendToChannel(c chan bool) {
	c <- true
	fmt.Println("sent: ", true)
}
