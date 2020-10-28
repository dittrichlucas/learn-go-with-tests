package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("\ngot: %.2f\nwant: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	testArea := []struct {
		name string
		form Form
		want float64
	}{
		{name: "Circle", form: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Rectangle", form: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Triangle", form: Triangle{Width: 10, Height: 5}, want: 25.0},
	}

	for _, tt := range testArea {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.form.Area()

			if got != tt.want {
				t.Errorf("%#v\ngot: %.2f\nwant: %.2f", tt.form, got, tt.want)
			}
		})
	}
	// checkArea := func(t *testing.T, form Form, want float64) {
	// 	t.Helper()
	// 	got := form.Area()

	// 	if got != want {
	// 		t.Errorf("\ngot: %.2f\nwant: %.2f", got, want)
	// 	}
	// }
	// t.Run("should return the area of a rectangle", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	checkArea(t, rectangle, 72.0)

	// })

	// t.Run("should return the area of a circle", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })
}
