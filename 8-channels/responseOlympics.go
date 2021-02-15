package channels

import (
	"math/rand"
	"time"
)

type Response struct {
	payload string
	server  string
}

func serverOne(resChan chan<- Response) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(105)))
	resChan <- Response{
		"payload",
		"one",
	}
}

func serverTwo(resChan chan<- Response) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(103)))
	resChan <- Response{
		"payload",
		"two",
	}
}

func serverThree(resChan chan<- Response) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	resChan <- Response{
		"payload",
		"three",
	}
}

func responseOlympics() Response {
	response := make(chan Response, 3)
	go serverOne(response)
	go serverTwo(response)
	go serverThree(response)
	return <-response
}
