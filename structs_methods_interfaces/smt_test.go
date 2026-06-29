package smi

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	check := func(t testing.TB, shape Shape, want float64) {
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("perimeter of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		check(t, rectangle, 40.0)
	})

	t.Run("circumference of circle", func(t *testing.T) {
		circle := Circle{10.0}
		check(t, circle, 2*math.Pi*10.0)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 2 * math.Pi * 10},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
