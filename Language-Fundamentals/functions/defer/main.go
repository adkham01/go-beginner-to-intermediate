package main

import "fmt"

func main() {
	multipleDefers()
	deferWithArguments()
}

// func (r receiver) identifier(p params) (r returns){...}

func multipleDefers() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
}

func deferWithArguments() {
	x := 1
	defer fmt.Println("Deferred:", x)
	x = 2
	fmt.Println("Current:", x)
}
