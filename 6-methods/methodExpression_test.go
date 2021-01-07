package geometry

import (
	"testing"
)

func TestMethodExpressions(t *testing.T) {
	sAreaFn := Square.area

	if sAreaFn(Square{10}) != 100 {
		t.Errorf("Incorrect square area")
	}
}

func TestTranslateByAdd(t *testing.T) {
	path := Path{{2, 2}, {2, 4}, {4, 4}}

	path.TranslateBy(Point{1, 1}, true)
	expected := Path{{3, 3}, {3, 5}, {5, 5}}

	for i := range path {
		if path[i] != expected[i] {
			t.Errorf("Incorrect translation")
		}
	}
}

func TestTranslateBySub(t *testing.T) {
	path := Path{{2, 2}, {2, 4}, {4, 4}}

	path.TranslateBy(Point{1, 1}, false)
	expected := Path{{1, 1}, {1, 3}, {3, 3}}

	for i := range path {
		if path[i] != expected[i] {
			t.Errorf("Incorrect translation")
		}
	}
}
