package main

import "fmt"

func main() {
	defer foo()
	func() {
		fmt.Println("This is anonymous func-1 ran")
	}()

	func(s string) {
		fmt.Println("My name is", s)
	}("Adkham")
}

func foo() {
	fmt.Println("Foo ran")
}

// a named function with identifier
// func (r receive) identifier(p parameter(s)) (r return(s)){...}

// an anonymous function
// func(p parameter(s)) (r returns(s)) {...}
