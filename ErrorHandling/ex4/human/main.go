package main

import (
	"dog"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(42),
	}
	fmt.Println(fido)
}
