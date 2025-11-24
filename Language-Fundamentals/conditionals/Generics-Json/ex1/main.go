package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	p1 := User{Name: "Alice", Age: 30}
	p2 := User{Name: "Bob", Age: 25}
	people := []User{p1, p2}

	p, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(p)
	fmt.Println(string(p))
}
