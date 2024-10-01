package distance

import (
	"math"
)

// Implementation of point. Has two private atributes: x and y
type Point struct {
	x int
	y int
}

// Point constructor
func MakePoint(x, y int) Point {
	return Point{x, y}
}

// Implementation of line. Has two private atributes: p1 and p2
type Line struct {
	p1, p2 Point
}

// Method to find an euclidian metric: length
func (l *Line) Length() float64 {
	return math.Sqrt(math.Abs((float64(l.p1.x) - float64(l.p2.x))) + math.Abs(float64(l.p1.y)-float64(l.p2.y)))
}
