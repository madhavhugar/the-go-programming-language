package channels

import "testing"

func TestThumbnailConverter(t *testing.T) {
	got := thumbnailConverter("kitty")
	wanted := "kitty_thumbnail"

	if got != wanted {
		t.Errorf("got != wanted; %s != %s", got, wanted)
	}
}
