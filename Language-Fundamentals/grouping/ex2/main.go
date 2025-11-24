package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, v := range xi {
		fmt.Printf("value is: %v\n", v)
	}

	fmt.Printf("And the type of slice is %T\n", xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
}
