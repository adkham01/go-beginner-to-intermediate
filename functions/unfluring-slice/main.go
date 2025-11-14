package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result2 := sum(xi...)

	result1 := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("First result:", result1)
	fmt.Println("Second result:", result2)
}

// Same as Java var arg method
func sum(ii ...int) int {
	sum := 0
	for _, v := range ii {
		sum += v
	}
	return sum
}
