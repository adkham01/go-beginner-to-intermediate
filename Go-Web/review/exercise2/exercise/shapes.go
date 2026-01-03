package exercise2

import "fmt"

const PI = 3.14

type Circle struct {
	Radius float64
}

func (c Circle) Area() {
	area := PI * 2 * c.Radius
	fmt.Println("Circle's area is", area)
}
