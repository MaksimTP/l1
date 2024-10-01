package distance

import (
	"math"
)

type Point struct {
	x int
	y int
}

func MakePoint(x, y int) Point {
	return Point{x, y}
}

type Line struct {
	P1, P2 Point
}

func (l *Line) Length() float64 {
	return math.Sqrt(math.Abs((float64(l.P1.x) - float64(l.P2.x))) + math.Abs(float64(l.P1.y)-float64(l.P2.y)))
}
