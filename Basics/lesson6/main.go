package main

import "fmt"

func main() {
	s, i, f := "happy", 42, 42.42

	fmt.Printf("%v is of type %T \n", s, s)
	fmt.Printf("%v is of type %T \n", f, f)
	fmt.Printf("%v is of type %T \n", i, i)
}
