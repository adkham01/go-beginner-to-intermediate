package main

import "fmt"

func main() {
	xi := []int{21, 43, 64, 54, 93, 82, 70}
	fmt.Println(xi)
	fmt.Println("----------------------------")
	xi = append(xi, 42, 58, 33, 49, 27)
	fmt.Println(xi)
	fmt.Println("----------------------------")

	xi2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf(" - %#v\n", xi2)
	fmt.Println("----------------------------")
}
