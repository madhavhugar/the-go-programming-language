package geometry

import (
	"math"
)

type Point struct{ x, y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-q.y)
}
