package smi

import "testing"

func TestAreaTableDriven(t *testing.T) {
	areaTests := []struct {
		msg      string
		shape    Shape
		expected float64
	}{
		{msg: "rectangle", shape: Rectangle{12, 6}, expected: 72.0},
		{msg: "circle", shape: Circle{10}, expected: 314.1592653589793},
		{msg: "triangle", shape: Triangle{12.0, 6.0}, expected: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.msg, func(t *testing.T) {
			actual := tt.shape.Area()
			if actual != tt.expected {
				t.Errorf("%#v: expected %g; actual %g", tt.shape, tt.expected, actual)
			}
		})
	}
}
