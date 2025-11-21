package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["James"] = 42
	fmt.Println(m)
	mapDelta(m)
	fmt.Println(m)
}

func intDelta(n *int) {
	*n = 43
}

func sliceDelta(ii []int) {
	ii[0] = 99
}

func mapDelta(m map[string]int) {
	m["Jane"] = 16
}
