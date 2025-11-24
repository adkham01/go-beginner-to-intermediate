package intro

import (
	"fmt"
	"testing"
)

func TestMySum(t *testing.T) {
	result := MySum(2, 3)
	expected := 5
	if result != expected {
		fmt.Printf("Test failed: expected %d, got %d\n", expected, result)
	}
	fmt.Println("Test passed result is:", result)

	result = MySum(4, 7, 5, 7)
	expected = 23
	if result != expected {
		fmt.Printf("Test failed: expected %d, got %d\n", expected, result)
	}
	fmt.Println("Test passed result is:", result)

	result = MySum(5, 9, 4, 0, 2)
	expected = 20
	if result != expected {
		fmt.Printf("Test failed: expected %d, got %d\n", expected, result)
	}
	fmt.Println("Test passed result is:", result)
}
