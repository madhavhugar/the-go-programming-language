package channels

import (
	"fmt"
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

func TestParallelGoRoutinesWithBufferedChannel(t *testing.T) {
	// happy case
	files := []string{"kitty", "foo", "cookie", "bar"}
	got, err := parallelGoRoutinesWithBufferedChannel(files)
	expected := []string{"kitty_thumbnail", "foo_thumbnail", "cookie_thumbnail", "bar_thumbnail"}

	assert.ElementsMatch(t, expected, got)
	assert.Equal(t, nil, err)

	// error case
	files = []string{"kitty", "foo", "error", "bar"}
	got, err = parallelGoRoutinesWithBufferedChannel(files)
	expected = nil

	assert.ElementsMatch(t, expected, got)
	assert.Equal(t, fmt.Errorf("invalid file"), err)
}

func TestParallelGoRoutinesWithWaitGroup(t *testing.T) {
	files := []string{"kitty", "foo", "cookie", "bar"}
	got, err := parallelGoRoutinesWithWaitGroup(files)
	expected := []string{"kitty_thumbnail", "foo_thumbnail", "cookie_thumbnail", "bar_thumbnail"}

	assert.ElementsMatch(t, expected, got)
	assert.Equal(t, nil, err)

	// error case
	files = []string{"kitty", "foo", "error", "bar"}
	got, err = parallelGoRoutinesWithWaitGroup(files)
	expected = nil

	assert.ElementsMatch(t, expected, got)
	assert.Equal(t, fmt.Errorf("invalid file"), err)
}
