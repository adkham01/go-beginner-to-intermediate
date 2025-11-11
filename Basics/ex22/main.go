package main

import (
	"fmt"
)

func main() {
	for i := range 5 {
		for j := range 5 {
			fmt.Printf("Outer loop %v inner loop %v\n", i, j)
		}
	}

}
