package main

import "fmt"

func main() {
	{
		// declaring a variable of type [7]int
		var ni [7]int
		// assign a value to index position zero
		ni[0] = 42
		fmt.Printf("%#v \t\t\t %T\n ", ni, ni)

		// array literal

		ni2 := [4]int{55, 56, 57, 58}

		// array literal
		fmt.Printf("%#v\t\t\t%T\n", ni2, ni2)

		// array literal
		// have compiler count elements
		ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
		fmt.Printf("%#v\t\t\t%T\n", ns, ns)

		// use buildin len function
		fmt.Println(len(ni))
		fmt.Println(len(ni2))
		fmt.Println(len(ns))

	}

}
