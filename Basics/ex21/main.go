package main

import (
	"fmt"
	"math/rand"
)

func main() {
	y := rand.Intn(5)

	iteration := 1
	for y > 0 {
		y++
		iteration++
		if y == 50 {
			break
		}
		if y%2 != 0 {
			fmt.Printf("iteration %v \t y is %v\n", iteration, y)
		} else {
			continue
		}
	}

}
