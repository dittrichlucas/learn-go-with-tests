package geometry

import "math"

// Form is
type Form interface {
	Area() float64
}

// Circle is
type Circle struct {
	Radius float64
}

// Rectangle is
type Rectangle struct {
	Width  float64
	Height float64
}

// Triangle is
type Triangle struct {
	Width  float64
	Height float64
}

// Area is
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Area is
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area is
func (t Triangle) Area() float64 {
	return (t.Width * t.Height) / 2
}

// Perimeter is
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area is
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
