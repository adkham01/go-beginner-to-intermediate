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

	bs, err := toJSON(p1)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))
}

func toJSON(a any) ([]byte, error) {

	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("There is an error during Marshaling %v", err)
	}

	return bs, nil
}
