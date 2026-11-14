package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) diameter() float64 {
	return 2 * c.radius
}

func main() {
	r := Rectangle{width: 3, height: 4}
	c := Circle{radius: 5}

	measure(r)
	measure(c)

	println("Circle diameter:", c.diameter())
	myPrinter("Hello", 42, 3.14, true)
	value(r)
}

func measure(g Geometry) {
	println(g.area())
	println(g.perim())
}

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func value(i interface{}) {
	fmt.Println(i)
}
