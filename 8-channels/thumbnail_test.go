package channels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThumbnailConverter(t *testing.T) {
	got := thumbnailConverter("kitty")
	wanted := "kitty_thumbnail"

	if got != wanted {
		t.Errorf("got != wanted; %s != %s", got, wanted)
	}
}

func TestParallelGoRoutinesWithUnbufferedChannel(t *testing.T) {
	files := []string{"kitty", "foo", "cookie", "bar"}
	got := parallelGoRoutinesWithUnbufferedChannel(files)
	expected := []string{"kitty_thumbnail", "foo_thumbnail", "cookie_thumbnail", "bar_thumbnail"}

	assert.ElementsMatch(t, expected, got)
}
