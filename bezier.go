package main

import "fmt"

type Point struct {
	X float64
	Y float64
}

func (p *Point) String() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

func (a Point) Plus(b Point) Point {
	return Point{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a Point) Multiply(f float64) Point {
	return Point{
		X: a.X * f,
		Y: a.Y * f,
	}
}

func Bezzier(points []Point, t float64) Point {
	if len(points) == 1 {
		return points[0]
	}
	return Bezzier(points[0:len(points)-1], t).Multiply(1.0 - t).Plus(Bezzier(points[1:], t).Multiply(t))
}
