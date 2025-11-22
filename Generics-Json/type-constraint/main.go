package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

type myNumbers interface {
	~int | ~float64
}

func addT[T myNumbers](a, b T) T {
	return a + b
}

type myPrimitive int

func main() {
	var age myPrimitive = 42
	fmt.Printf("%T\t%v\n", age, age)
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))
	fmt.Println(addT(age, 2))
	fmt.Println(addT(1.2, 2.2))
}
