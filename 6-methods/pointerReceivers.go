package geometry

func (p *Point) ScaleBy(n float64) {
	p.x = p.x * n
	p.y = p.y * n
}

func (p Point) ScaleByNonPointerReceiver(n float64) (q Point) {
	q.x = p.x * n
	q.y = p.y * n
	return q
}
