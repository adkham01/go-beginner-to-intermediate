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

type shape interface {
	Area() float64
}

func info(s shape) float64 {
	return s.Area()
}

func main() {
	c1 := circle{
		radius: 12.3,
	}

	s1 := square{
		length: 3,
		width:  5,
	}
	fmt.Println(info(c1))
	fmt.Println(info(s1))
}

func (s square) Area() float64 {
	return (s.length * s.width)
}

func (c circle) Area() float64 {
	return (math.Pi * math.Pow(2, c.radius))
}
