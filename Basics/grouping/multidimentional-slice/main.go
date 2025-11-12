package main

import "fmt"

func main() {

	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Trump", "Guiness", "Milk"}
	fmt.Println(jb)
	fmt.Println(jm)

	xs := [][]string{jb, jm}
	fmt.Println(xs)
}
