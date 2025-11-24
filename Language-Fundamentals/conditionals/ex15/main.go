package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initailization for my program occurs")
}

func main() {
	x := rand.Intn(300)
	fmt.Printf("Value of x: %v\n", x)

	switch {
	case x >= 0 && x <= 100:
		fmt.Println("Less then 100\t", x)
	case x >= 101 && x <= 200:
		fmt.Println("Between 100 and 200\t", x)
	case x >= 201 && x <= 250:
		fmt.Println("Between 201 and 250\t", x)
	default:
		fmt.Println("More than 250\t", x)
	}

}
