package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return (math.Pi * math.Pow(c.radius, 2))
}

type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Println(s.area())
}

func main() {
	mySquare := square{
		length: 20,
		width:  10,
	}

	myCircle := circle{
		radius: 5,
	}

	info(mySquare)
	info(myCircle)
}
