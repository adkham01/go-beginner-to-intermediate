package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastName  string
	favIC     []string
}

func main() {

	p1 := person{
		firstname: "Tom",
		lastName:  "Smith",
		favIC:     []string{"Chocolate", "Banana"},
	}

	p2 := person{
		firstname: "John",
		lastName:  "Simon",
		favIC:     []string{"Apple", "Malina"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
