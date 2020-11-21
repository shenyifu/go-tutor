package struct_perimeter

import "testing"

func TestPerimeter(t *testing.T) {

	testPerimeter := func(t *testing.T, want, got float64) {
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Reactangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0
		testPerimeter(t, want, got)
	})
	t.Run("circle", func(t *testing.T) {
		rectangle := Circle{10.0}
		got := rectangle.Perimeter()
		want := 62.83
		testPerimeter(t, want, got)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Reactangle{12, 6}, 72.0},
		//{Circle{10}, 314.1592653589793},

		{Triangle{2, 7}, 7.0},
	}

	for _, td := range areaTests {
		got := td.shape.Area()
		want := td.want
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

}