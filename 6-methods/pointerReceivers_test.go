package geometry

import "testing"

func TestScaleBy_SyntacticSugarInvocation(t *testing.T) {
	p := &Point{1, 2}

	p.ScaleBy(2)
	if p.x != 2 && p.y != 4 {
		t.Errorf("Did not scale so well")
	}
}

func TestScaleBy_DeReferenceReceiver(t *testing.T) {
	p := Point{1, 2}
	ptr := &p
	ptr.ScaleBy(2)
	if p.x != 2 && p.y != 4 {
		t.Errorf("Did not scale so well")
	}
}

func TestScaleBy_ReferenceReceiverAddress(t *testing.T) {
	p := Point{1, 2}
	q := (&p).ScaleByNonPointerReceiver(2)
	if q.x != 2 && q.y != 4 {
		t.Errorf("Did not scale so well")
	}
}
