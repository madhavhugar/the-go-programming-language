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

func TestColoredShape_PromotedMethods(t *testing.T) {
	cs := ColoredShapes{Square{10}, Rectangle{10, 20}, color.RGBA{1, 1, 1, 1}}

	// cs.area()
	// will not work since go will complain about ambigious selection of promoted methods
	// both Rectange and Square have area method. So we have to explicitly specify which methods
	// needs to be called
	sqArea := cs.Square.area()
	reArea := cs.Rectangle.area()

	if sqArea != 100 {
		t.Errorf("Incorrect square area")
	}

	if reArea != 200 {
		t.Errorf("Incorrect rectangle area")
	}
}

func TestShapeArea(t *testing.T) {
	s := Square{10}
	r := Rectangle{10, 20}

	sqArea := ShapeArea(s)
	reArea := ShapeArea(r)

	if sqArea != 100 {
		t.Errorf("Incorrect square area")
	}

	if reArea != 200 {
		t.Errorf("Incorrect rectangle area")
	}
}

func TestAnonymousStruct(t *testing.T) {
	// only via embedding an anonymous struct can have methods
	twoDimShape := struct {
		Rectangle
	}{
		Rectangle: Rectangle{10, 20},
	}

	if twoDimShape.area() != 200 {
		t.Errorf("Incorrect rectangle area")
	}
}
