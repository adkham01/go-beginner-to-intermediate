package main

import "fmt"

var i int

func main() {
	fmt.Println(i)

	z := 42
	fmt.Println(z)

	a, b := 43, "Happy"
	fmt.Println(a, "\t", b)

	var c float32 = 42.42
	fmt.Printf("%v is of this type %T\n", c, c)

	e, f, _ := 45, 46, 47
	fmt.Println(e, f)
}
