package geometry

import (
	"testing"
)

func TestDeclaration(t *testing.T) {
	p := Point{2, 2}
	q := Point{4, 4}
	got := p.Distance(q)
	var expected float64 = 2
	if got != expected {
		t.Errorf("%f != %f", got, expected)
	}
}
