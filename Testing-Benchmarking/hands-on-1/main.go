package main

import (
	"fmt"
	"hands-on-1/dog"
)

type caine struct {
	nume string
	age  int
}

func main() {
	fido := caine{nume: "Fido", age: 16}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(fido.age))
}
