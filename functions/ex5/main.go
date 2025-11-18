package main

import "fmt"

type Person struct {
	firstName string
	age       int
}

func main() {
	p := Person{
		firstName: "Todd",
		age:       45,
	}
	p.speak()
}

func (p Person) speak() {
	fmt.Printf("My name is %v and my age %v\n", p.firstName, p.age)
}
