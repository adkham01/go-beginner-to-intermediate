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

}
