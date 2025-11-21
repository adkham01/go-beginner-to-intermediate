package main

import "fmt"

func main() {
	fmt.Println(doMath(42, 16, add))
	fmt.Println(doMath(42, 16, sub))
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}
