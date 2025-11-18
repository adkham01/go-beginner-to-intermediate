package main

import "fmt"

func main() {
	check := []int{34, 27, 43, 25, 43, 64}
	fmt.Println(foo(check...))
	fmt.Println(bar(check))
}

func foo(ii ...int) int {
	return unfurl(ii...)
}
func unfurl(ii ...int) int {
	total := 0
	for _, v := range ii {
		total += v
	}

	return total
}

func bar(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}

	return
}
