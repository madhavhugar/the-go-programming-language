package channels

import "time"

var shouldBeSetToTenEventually int

func syncRoutinesWithUnBufferedChannels() {
	shouldBeSetToTenEventually = 5
	done := make(chan bool)
	go func() {
		time.Sleep(time.Second * 5)
		shouldBeSetToTenEventually = 10
		done <- true
	}()
	<-done
}
