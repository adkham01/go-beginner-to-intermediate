package main

import (
	"fmt"
	"math"
)

func main() {

	f := powerNator(2)

	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 4
	fmt.Println(f()) // 5
}

func powerNator(a float64) func() float64 {
	var c float64 = 0
	return func() float64 {
		c++
		return math.Pow(a, c)
	}

}
