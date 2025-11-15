package main

import "fmt"

func main() {
	f := incrementor()
	xi := []int{}

	for range 10 {
		xi = append(xi, f())
	}
	fmt.Println(xi)
}

func incrementor() func() int {
	x := 0
	return func() int {
		x += 5
		return x
	}
}
