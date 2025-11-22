package main

import (
	"encoding/json"
	"fmt"
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

	var people []User

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}

	fmt.Println("All data")
	for _, v := range people {
		fmt.Printf("Person: %v,\n", v)
	}
}
