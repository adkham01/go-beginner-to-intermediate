package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("This is anonymous func-1 ran")
	}

	y := func(s string) {
		fmt.Println("My name is", s)
	}
	y("Adkham")
	x()
}

// a named function with identifier
// func (r receive) identifier(p parameter(s)) (r return(s)){...}

// an anonymous function
// func(p parameter(s)) (r returns(s)) {...}
