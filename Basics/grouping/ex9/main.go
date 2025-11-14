package main

import "fmt"

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
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("value : %#v \t type : %T\n", p1, p1)
	fmt.Printf("value : %#v \t type : %T\n", p2, p2)

	for i, v := range p1.favIC {
		fmt.Println(i, v)
	}

	for i, v := range p2.favIC {
		fmt.Println(i, v)
	}
}
