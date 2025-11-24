package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initailization for my program occurs")
}

func main() {

	for i := 1; i <= 42; i++ {
		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Printf("Iteration %v \t x is %v\n", i, x)
		case 1:
			fmt.Printf("Iteration %v \t x is %v\n", i, x)
		case 2:
			fmt.Printf("Iteration %v \t x is %v\n", i, x)
		case 3:
			fmt.Printf("Iteration %v \t x is %v\n", i, x)
		case 4:
			fmt.Printf("Iteration %v \t x is %v\n", i, x)
		default:
			fmt.Println("none of the previous were met")
		}
	}

}
