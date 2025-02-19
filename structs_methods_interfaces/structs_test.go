package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g wanted %g", got, want)
		}

	}
	t.Run("rectange", func(t *testing.T) {
		rectangle := Rectange{10.0, 10.0}
		checkPerimeter(t, rectangle, 40.0)

	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkPerimeter(t, circle, 62.83185307179586)
	})

}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		HasArea float64
	}{
		{name: "Rectangle", shape: Rectange{10.0, 10.0}, HasArea: 100.0},
		{name: "Cirle", shape: Circle{10.0}, HasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{10.0, 10.0}, HasArea: 50.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.HasArea {
				t.Errorf("for %#v got %g wanted %g", tt.shape, got, tt.HasArea)
			}
		})
	}
}
