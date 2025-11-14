package main

import (
	"fmt"
)

func main() {
	as := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Tom",
		friends: map[string]int{
			"Jenny":   5,
			"Simon":   4,
			"Olivier": 6,
		},
		favDrinks: []string{"RedBull", "Coke", "Water"},
	}

	for _, v := range as.friends {
		fmt.Println(as.first, "-- friends --", v)
	}

	for _, v := range as.favDrinks {
		fmt.Println(as.first, "-- Faviorit Drinks --", v)
	}

}
