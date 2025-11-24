package main

import "fmt"

func main() {
	sum := sum("Todd", 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum)
}

// Same as Java var arg method
func sum(s string, ii ...int) int {
	sum := 0
	for _, v := range ii {
		sum += v
	}
	fmt.Println("My dog's name is", s)
	return sum
}
