package structs

import "testing"

func TestArea(t *testing.T) {
	// creating an anonymous struct (areaTests) with the fields 'shape' and 'want'
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	// iterate over them like any other slice using the struct fields
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
