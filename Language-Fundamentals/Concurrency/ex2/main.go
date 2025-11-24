package main

import "fmt"

type person struct {
	name string
	age  int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("My name is", p.name)
	fmt.Println("My age is", p.age)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "Todd",
		age:  42,
	}
	p2 := person{
		name: "Jane",
		age:  35,
	}

	saySomething(&p1)
	p2.speak()
}
