package geometry

import "testing"

func TestMethodValues(t *testing.T) {
	s := Square{10}
	sAreaFn := s.area

	if sAreaFn() != 100 {
		t.Errorf("Incorrect square area")
	}
}
