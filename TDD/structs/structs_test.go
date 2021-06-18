package structs

import "testing"

type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// Table-driven tests for duplicate tests on different inputs
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		// use explicit property assignment so it is clear that this is an assertion of truth
		// rather than sequence of operations
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			// use %#v to print struct with all fields for better test message with structs
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}

}
