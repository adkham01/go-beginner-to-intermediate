package main

import (
	"fmt"
	"math/rand"
)

func main() {
	y := rand.Intn(3)

	iteration := 1
	for y < 10 {
		fmt.Printf("iteration %v \t y = %v\n", iteration, y)
		y += 1
		iteration += 1
	}

}
