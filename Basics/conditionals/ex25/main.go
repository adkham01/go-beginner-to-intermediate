package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(5)

	if x == 3 {
		fmt.Println("X is 3")
	}

	fmt.Printf("true && true\t\t%v\n", (true && true))
	fmt.Printf("true && false\t\t%v\n", (true && false))
	fmt.Printf("true || true\t\t%v\n", (true || true))
	fmt.Printf("true || false\t\t%v\n", (true || false))
	fmt.Printf("!true\t\t\t%v\n", (!true))

}
