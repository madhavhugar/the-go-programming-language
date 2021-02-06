package channels

import "testing"

func TestHowToPanicWithChannels(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic!")
		}
	}()
	// On send operation to a closed channel, Go panics
	panicOnSendToClosedChannel()
}
