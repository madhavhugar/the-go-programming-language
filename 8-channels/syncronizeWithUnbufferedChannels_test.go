package channels

import "testing"

func TestSyncWithUnbufferedChannels(t *testing.T) {
	syncRoutinesWithUnBufferedChannels()
	wanted := 10
	if shouldBeSetToTenEventually != wanted {
		t.Errorf("channels did not syncronize, and value was not set eventually")
	}
}
