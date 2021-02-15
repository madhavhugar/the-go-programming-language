package channels

import (
	"testing"
)

func TestResponseOlympics(t *testing.T) {
	got := responseOlympics()
	t.Errorf("%+v", got)
}
