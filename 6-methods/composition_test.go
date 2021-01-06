package geometry

import (
	"image/color"
	"testing"
)

func TestColoredPoint_PromotedMethod_Distance(t *testing.T) {
	c := ColoredPoint{Point{1, 2}, color.RGBA{1, 1, 1, 1}}

	q := Point{3, 2}
	got := c.Distance(q)
	var expected float64 = 2
	if got != expected {
		t.Errorf("%f != %f", got, expected)
	}
}
