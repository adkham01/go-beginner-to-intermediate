package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func swap(x string, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 8
	y = sum - x
	return
}

var (
	toBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var a, b string = "world", "Hello"
	fmt.Println(swap(a, b))
	fmt.Println(split(17))

	fmt.Printf("type: %T Value %v\n", toBe, toBe)
	fmt.Printf("type: %T Value %v\n", maxInt, maxInt)
	fmt.Printf("type: %T Value %v\n", z, z)

	var i int
	var f float64
	var db bool
	var s string
	fmt.Printf("%v, %v, %v, %q\n", i, f, db, s)

	var x, y int = 3, 4
	var df float64 = math.Sqrt(float64(x*x + y*y))
	var dz uint = uint(df)
	fmt.Println(x, y, dz)
}
