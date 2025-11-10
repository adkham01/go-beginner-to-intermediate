package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initailization for my program occurs")
}

func main() {

	for i := 0; i <= 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("iteration %v\tvale of x: %v value of y: %v\n", i, x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("x and y are both less than 4")
		case x > 6 && y > 6:
			fmt.Println("x and y are both greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("x greater than 4 or less than or equal to 6")
		case y != 5:
			fmt.Println("y is not equal to 5")
		default:
			fmt.Println("none of the previous were met")
		}
	}

}
