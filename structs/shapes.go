package structs

import "math"

type Shape interface {
	Area() float64
}

// A struct is a named collection of fields where you can store data
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface
// This is a method: func (receiverName ReceieverType) MethodName(args)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
