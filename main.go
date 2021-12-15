package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Position
	x2, y2 float64
}

type Circle struct {
	Position
	r float64
}

type Position struct {
	x, y float64
}

type Shape interface {
	Area() float64
}

func fullArea(s ...Shape) float64 {
	fullArea := 0.0
	for _, shape := range s {
		fullArea += shape.Area()
	}
	return fullArea
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) Area() float64 {
	l := distance(r.x, r.y, r.x, r.y2)
	w := distance(r.x, r.y, r.x2, r.y)
	return l * w
}

func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	rect := &Rectangle{
		Position: Position{
			x: 0,
			y: 0,
		},
		x2: 10,
		y2: 10,
	}
	circ := &Circle{
		Position: Position{
			x: 0,
			y: 0,
		},
		r: 5,
	}

	//if nil == nil {
	//	println("here")
	//}
	fmt.Println(fullArea(rect, circ))
}
