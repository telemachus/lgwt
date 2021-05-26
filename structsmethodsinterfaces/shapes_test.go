package smi

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	actual := r.Perimeter()
	expected := 40.0

	if actual != expected {
		t.Errorf("expected %f; actual %f", expected, actual)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, s Shape, expected float64) {
		t.Helper()
		actual := s.Area()
		if actual != expected {
			t.Errorf("expected %g; actual %g", expected, actual)
		}
	}

	t.Run("a rectangle", func(t *testing.T) {
		r := Rectangle{12,6}
		checkArea(t, r, 72.0)
	})

	t.Run("a circle", func(t *testing.T) {
		c := Circle{10}
		checkArea(t, c, 314.1592653589793)
	})
}
