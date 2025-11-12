package main

// the imported package "fmt"
// is in the "file block"
// on this file
import (
	"fmt"
)

var x = 42

func main() {
	fmt.Println(x)

	// y does not exist here yet
	// so this doesn't work

	// y is "block scope"
	y := 32
	fmt.Println(y)

	p1 := person{
		"James",
	}

	p1.sayHello()

	// variable "shadowing"
	x := 32
	fmt.Println(x)

	// furtherExplored.Fascinating()

	// also good to know

	if z := 82; z > 50 {
		fmt.Println(z)
	}
	printMe()

	// you cannot access z here

}
func printMe() {
	// x can be accessed here
	fmt.Println(x)
}

// type person is sin "package block" scope
type person struct {
	first string
}

// the method sayHello
// which is attached to VALUES of TYPE person
func (p person) sayHello() {
	fmt.Println("Hi my name is ", p.first)
}
