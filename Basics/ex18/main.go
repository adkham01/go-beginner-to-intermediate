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

		fmt.Printf("iteration %v\tvale of x: %v value of y: %v\n", i, x)

		switch x {
		case 0:
			fmt.Println("X is 0")
		case 1:
			fmt.Println("X is 1")
		case 2:
			fmt.Println("X is 2")
		case 3:
			fmt.Println("X is 3")
		case 4:
			fmt.Println("X is 4")
		default:
			fmt.Println("none of the previous were met")
		}
	}

}
