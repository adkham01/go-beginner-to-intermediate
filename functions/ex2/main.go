package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	i, s := bar()
	fmt.Println(i, s)

}

func foo() int {
	age := 24
	return age
}

func bar() (int, string) {
	age := 23
	name := "Tom"
	return age, name
}
