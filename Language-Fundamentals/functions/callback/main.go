package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	x := doMath(65, 21, add)
	fmt.Println(x)

	y := doMath(89, 18, subtract)
	fmt.Println(y)

	fmt.Println(doMath(x, y, subtract))
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}
