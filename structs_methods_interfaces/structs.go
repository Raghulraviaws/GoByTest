package structs

import "math"

type Rectange struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height float64
	Base   float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectange) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectange) Area() float64 {
	return float64(r.Height) * float64(r.Width)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return 0.0
}
