package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := r.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 10.0, Height: 10.0}, want: 100},
		{shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{shape: Triangle{Width: 10.0, Height: 5.0}, want: 25.0},
	}
	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}
}
