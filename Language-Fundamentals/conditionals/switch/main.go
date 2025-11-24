package main

import "fmt"

func main() {

	x := 40

	switch {
	case x < 42:
		fmt.Println("x is LESS than 42")
	case x == 42:
		fmt.Println("x is EQUAL to 42")
	case x > 42:
		fmt.Println("x is GRATER than 42")
	default:
		fmt.Println("This is the default case for 42")
	}

	println("First switch block ended\n\n\n")

	// no default  fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("Printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("Printed because of fallthrough: x is 42")
	case 43:
		fmt.Println("Printed because of fallthrough: x is 43")
	default:
		fmt.Println("Printed because of fallthrough: this is default case for x")
	}
	println("Second switch block ended\n\n\n")

	// no default  fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("Printed because of fallthrough: x is 41")
		fallthrough
	case 42:
		fmt.Println("Printed because of fallthrough: x is 42")
		fallthrough
	case 43:
		fmt.Println("Printed because of fallthrough: x is 43")
		fallthrough
	default:
		fmt.Println("Printed because of fallthrough: this is default case for x")
	}

}
