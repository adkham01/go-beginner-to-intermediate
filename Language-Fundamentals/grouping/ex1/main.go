package main

import "fmt"

func main() {
	ai := [5]int{}

	for i := 0; i < len(ai); i++ {
		ai[i] = i + 1
	}

	for i, v := range ai {
		fmt.Printf("Index is : %v, value is : %v\n", i, v)
	}
	fmt.Printf("Type of array is %T", ai)
}
