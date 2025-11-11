package main

import (
	"fmt"
	"math/rand"
)

func main() {
	y := rand.Intn(50)

	for y > 0 {
		fmt.Println("y ", y)
		y--
	}

}
