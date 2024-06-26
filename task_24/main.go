package main

import (
	"fmt"
	"math"
)

type PointAction interface {
	GetDist(p1 Point, p2 Point) float64
}

type Point struct {
	x float64
	y float64
}

func (p *Point) GetDist(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)
	fmt.Println(p1.GetDist(*p1, *p2))
}
