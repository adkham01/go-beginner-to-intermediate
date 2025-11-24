package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("Value of x: %v\n", x)

	if x >= 0 && x <= 100 {
		fmt.Println("Less then 100\t", x)
	} else if x >= 101 && x <= 200 {
		fmt.Println("Between 100 and 200\t", x)
	} else if x >= 201 && x <= 250 {
		fmt.Println("Between 201 and 250\t", x)
	} else {
		fmt.Println("More than 250\t", x)
	}

	for range 10 {
		y := rand.Intn(3)
		fmt.Printf("%v\t", y)
		println()
	}

}
