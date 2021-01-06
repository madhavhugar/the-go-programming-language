// method promotion with one embedded field
package geometry

import (
	"image/color"
)

type ColoredPoint struct {
	Point
	Color color.RGBA
}

// struct with multiple anonymous fields

// unnamed structs types methods
