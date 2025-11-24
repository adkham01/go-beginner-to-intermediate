package main

import "fmt"

type person struct {
	firstName string
}

func (p person) speak(age int) {
	fmt.Println("I am", p.firstName, "and my age is", age)
}

func main() {
	p1 := person{
		firstName: "James",
	}

	p2 := person{
		firstName: "Janny",
	}
	p1.speak(32)
	p2.speak(27)

}

// func (r receiver) identifier(p params) (r returns){...}
