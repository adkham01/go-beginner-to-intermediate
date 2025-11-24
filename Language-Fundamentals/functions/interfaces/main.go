package main

import (
	"fmt"
	"hash/fnv"
)

type person struct {
	firstName string
}

type secretAgent struct {
	person person
	ltk    bool
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("I am", p.firstName)
}

func (sa secretAgent) speak() {
	fmt.Println("I am the secret agent my id is", hashCode(sa.person.firstName),
		"has lisence to kill", sa.ltk)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa := secretAgent{
		person: person{
			firstName: "James",
		},
		ltk: true,
	}

	p2 := person{
		firstName: "Janny",
	}
	// sa.speak()
	// p2.speak(27)

	saySomething(p2)
	saySomething(sa)

}

// func (r receiver) identifier(p params) (r returns){...}
func hashCode(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}
