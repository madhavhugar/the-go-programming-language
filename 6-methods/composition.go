package geometry

import (
	"image/color"
)

// ColoredPoint objects can access the methods of Point
// method promotion with one embedded field
type ColoredPoint struct {
	Point
	Color color.RGBA
}

// Shape represents 2D or 3D geometric shape
type Shape interface {
	area() float64
}

// Square represents a square shape
type Square struct {
	length float64
}

func (s Square) area() float64 {
	return s.length * s.length
}

// Rectangle represents a rectangle shape
type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

// ColoredShapes objects can access methods of Square and Rectangle
// struct with multiple anonymous fields
type ColoredShapes struct {
	Square
	Rectangle
	Color color.RGBA
}

// ShapeArea computes the area of a given shape object
func ShapeArea(s Shape) float64 {
	return s.area()
}
