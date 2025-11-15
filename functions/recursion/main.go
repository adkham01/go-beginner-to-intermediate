package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factLoop(5))
}

func factorial(n int) int {
	fmt.Println("n is equal to:", n)
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func factLoop(n int) int {
	fac := 1
	copyN := n
	for i := 1; i < n; i++ {
		fac *= copyN
		copyN--
	}
	return fac
}
