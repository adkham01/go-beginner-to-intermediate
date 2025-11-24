package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("vale of x: %v value of y: %v\n", x, y)

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
