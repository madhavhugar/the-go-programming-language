package geometry

import (
	"testing"
)

func TestDistanceMethod(t *testing.T) {
	p := Point{2, 2}
	q := Point{2, 4}
	got := p.Distance(q)
	var expected float64 = 2
	if got != expected {
		t.Errorf("%f != %f", got, expected)
	}
}

func TestDistanceFunction(t *testing.T) {
	p := Point{2, 2}
	q := Point{2, 4}
	got := Distance(p, q)
	var expected float64 = 2
	if got != expected {
		t.Errorf("%f != %f", got, expected)
	}
}

func TestPathDistance(t *testing.T) {
	p := Path{{2, 2}, {2, 4}, {4, 4}}
	got := p.Distance()
	var expected float64 = 4
	if got != expected {
		t.Errorf("%f != %f", got, expected)
	}

	p = Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	got = p.Distance()
	expected = 12
	if got != expected {
		t.Errorf("%f != %f", got, expected)
	}
}
