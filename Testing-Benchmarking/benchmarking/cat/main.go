package main

import (
	"benchmarking/cat/mystr"
	"fmt"
	"strings"
)

const s = "We ask not for wealth, but for freedom. We ask not for a rise in standard of living but for a life of dignity."

func main() {
	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n\n", mystr.Join(xs))
}
