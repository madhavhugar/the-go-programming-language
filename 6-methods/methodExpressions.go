package geometry

// Add moves a point by a adding a given offset
func (p Point) Add(q Point) Point { return Point{p.x + q.x, p.y + q.y} }

// Sub moves a point by a subracting a given offset
func (p Point) Sub(q Point) Point { return Point{p.x - q.x, p.y - q.y} }

// TranslateBy moves the point (add or sub) by a given offset
func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point

	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}
