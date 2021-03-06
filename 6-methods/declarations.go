package geometry

import (
	"math"
)

type Point struct{ x, y float64 }

// NewPoint is a constructor for Point
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p Point) privateMethod() {
	// do nothing
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

type Path []Point

func (p Path) Distance() float64 {
	if p == nil {
		return 0
	}
	var total float64
	for i := range p {
		if i == 0 {
			continue
		}
		total += p[i-1].Distance(p[i])
	}
	return total
}
