package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
	Age       int
}

type secretAgernt struct {
	person        person
	licenseToKill bool
}

func main() {
	p1 := person{
		FirstName: "James",
		LastName:  "Bond",
		Age:       32,
	}

	p2 := person{
		FirstName: "Miss",
		LastName:  "Moneypenny",
		Age:       27,
	}

	people := []person{p1, p2}

	fmt.Println("people:", people)

	bc, err := json.Marshal(people)

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("JSON:", string(bc))
}
