package main

import (
	"fmt"
	"math"
)

func main() {
	var i int8 = math.MaxInt8
	var j uint8 = math.MaxUint8

	k := "1 + 1"

	fmt.Println(i)
	fmt.Println(j)
	fmt.Printf("%v \t %T", k, k)

}
