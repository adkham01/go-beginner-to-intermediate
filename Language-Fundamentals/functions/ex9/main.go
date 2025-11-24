package main

import "fmt"

func main() {
	f := doMath()
	fmt.Printf("Value is %v\n", f())
}

func doMath() func() int {
	return func() int {
		return 42
	}
}
