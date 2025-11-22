package main

import (
	"fmt"
	"math"
)

type User struct {
	ID     int
	Name   string
	Email  string
	Active bool
}

func main() {
	s := `[
		{"id": 1, "name": "John Doe", "email": "john@example.com", "active": true},
		{"id": 2, "name": "Jane Smith", "email": "jane@example.com", "active": false},
		{"id": 3, "name": "Bob Johnson", "email": "bob@example.com", "active": true}
	]`

	bs := []byte(s)

	fmt.Printf("%T\n", bs)
	fmt.Printf("%T\n", s)

	fmt.Println("JSON:", s)
	fmt.Println(math.Sqrt(17))
}
