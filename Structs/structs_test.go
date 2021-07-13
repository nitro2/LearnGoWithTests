package structs

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("Perimeter of non zero rectangle", func(t *testing.T) {
		rect := Rectangle{10, 5}
		got := rect.Perimeter()
		want := 30.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("Perimeter of zero rectangle", func(t *testing.T) {
		rect := Rectangle{0, 0}
		got := rect.Perimeter()
		want := 0.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
