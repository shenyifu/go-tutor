package struct_perimeter

import "math"

type Reactangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Redius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (reactangle Reactangle) Perimeter() float64 {
	return 2 * (reactangle.Height + reactangle.Width)

}

func (reactangle Reactangle) Area() float64 {
	return reactangle.Height * reactangle.Width
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Redius * math.Pi
}

func (c Circle) Area() float64 {
	return math.Pi * c.Redius * c.Redius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
