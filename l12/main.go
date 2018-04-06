package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

func info (s Shape) {
	fmt.Println(s.area())
}

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	square := Square{10}
	circle := Circle{10}

	info(square)
	info(circle)
}
