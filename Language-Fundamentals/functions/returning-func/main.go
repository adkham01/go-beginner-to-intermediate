package main

import "fmt"

func main() {
	x := foo()
	y := bar()

	fmt.Println(x + y())

	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", y)
}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 44
	}
}

// a named function with identifier
// func (r receive) identifier(p parameter(s)) (r return(s)){...}
// 		an anonymous function
// func(p parameter(s)) (r returns(s)) {...}
