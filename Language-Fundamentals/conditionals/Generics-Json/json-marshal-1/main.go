package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

type Animal struct {
	Name  string
	Order string
}

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "Redis",
		Colors: []string{"Crimson", "Red", "Rubby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error :", err)
	}
	os.Stdout.Write(b)

	var jsonBlob = []byte(`[{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll", "Order": "Dasyuromorphia"}]`)

	var animals []Animal
	err = json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("\n%+v", animals)
}
