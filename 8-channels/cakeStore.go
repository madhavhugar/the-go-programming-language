package channels

import (
	"fmt"
	"math/rand"
	"time"
)

type CakeStore struct {
	Verbose    bool
	Cakes      int
	BakeTime   time.Duration
	BakeStdDev time.Duration
	BakeBuffer int

	IcingStaff  int
	IcingTime   time.Duration
	IcingStdDev time.Duration
	IcingBuffer int

	InscribeTime   time.Duration
	InscribeStdDev time.Duration
}

func (s *CakeStore) bakeCake(bakeChan chan int) {
	defer close(bakeChan)
	for i := 0; i < s.Cakes; i++ {
		if s.Verbose {
			fmt.Println("Baking cake: ", i)
		}
		work(s.BakeTime, s.BakeStdDev)
		bakeChan <- i
	}
}

func (s *CakeStore) iceCake(bakeChan chan int, inscribeChan chan int) {
	defer close(inscribeChan)
	for c := range bakeChan {
		if s.Verbose {
			fmt.Println("Icing cake: ", c)
		}
		work(s.IcingTime, s.IcingStdDev)
		inscribeChan <- c
	}
}

func (s *CakeStore) inscribeCake(inscribeChan chan int) {
	for c := range inscribeChan {
		if s.Verbose {
			fmt.Println("Inscribing cake: ", c)
		}
		work(s.InscribeTime, s.InscribeStdDev)
	}
}

func work(workTime, stdDev time.Duration) {
	delay := workTime + time.Duration(rand.NormFloat64()*float64(stdDev))
	time.Sleep(delay)
}

func (s *CakeStore) Work(n int) {
	bakeChan := make(chan int, s.BakeBuffer)
	iceChan := make(chan int, s.IcingBuffer)

	go s.bakeCake(bakeChan)
	for i := 0; i < s.IcingStaff; i++ {
		go s.iceCake(bakeChan, iceChan)
	}
	s.inscribeCake(iceChan)
}
