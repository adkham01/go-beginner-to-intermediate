package main

import "fmt"

func main() {
	add := x(42, 16)
	fmt.Println(add)

	fmt.Println(x(34, 75))
}

var x = func(a int, b int) int {
	return a + b
}
