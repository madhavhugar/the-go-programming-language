package main

import (
	geometry "go-prog-lang/6-methods"
	"testing"
)

func TestEncapsulation(t *testing.T) {
	// The "unit of encapsulation" in golang is package level

	// accessing encapsulated fields from Point is not possible
	// below line throws an compilation error
	// p := geometry.Point{x: 1, y: 1}

	p := geometry.NewPoint(1, 1)
	q := geometry.NewPoint(1, 3)

	got := p.Distance(q)
	var expected float64 = 2
	if got != expected {
		t.Errorf("%f != %f", got, expected)
	}

	// encapsulated methods cannot be accessed
	// below line throws an compilation error
	// p.privateMethod()
}
