package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	FirstName string
	LastName  string
	Sayings   []string
}

func main() {
	p1 := person{
		FirstName: "James",
		LastName:  "Bond",
		Sayings:   []string{"Sahken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println("There is an error Marshaling", err)
	}
	fmt.Println(string(bs))
}
