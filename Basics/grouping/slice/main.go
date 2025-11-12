package main

import "fmt"

func main() {
	xs := []string{"Hello ", "World", "Tom", "Jerry"}
	fmt.Printf("Values : %v, length is: %v \n", xs, (len(xs)))

	for i, v := range xs {
		fmt.Printf("index is : %v value is: %v\n", i, v)
	}
	fmt.Printf("Type is %T\n", xs)
	println("-----------------------------\n")

	for _, v := range xs {
		fmt.Printf("value is: %v\n", v)
	}
	fmt.Printf("Type is %T\n", xs)
	println("-----------------------------\n")

	// acesing the items by index
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
	fmt.Println(xs[3])
}
