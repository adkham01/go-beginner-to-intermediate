package main

import "fmt"

func main() {
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			fmt.Printf("Element at [%d][%d]: %d\n", i, j, matrix[i][j])
		}

	}

	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("Fruits:", fruits)

	fruits = append(fruits, "limon", "orange", "grape")
	fmt.Println("Updated Fruits:", fruits)

	for i, v := range fruits {
		fmt.Printf("Fruit %d: %s\n", i+1, v)
	}

	capitalCities := map[string]string{
		"France":      "Paris",
		"Italy":       "Rome",
		"Germany":     "Berlin",
		"Portugal":    "Lisbon",
		"Netherlands": "Amsterdam",
	}

	capital, exists := capitalCities["Italy"]
	if exists {
		fmt.Println("The capital of Italy is", capital)
	} else {
		fmt.Println("Capital not found")
	}

	delete(capitalCities, "Germany")
}
