package smi

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (t Triangle) Perimeter() float64 {
	// This is a dummy implementation for the moment.
	return 0
}
