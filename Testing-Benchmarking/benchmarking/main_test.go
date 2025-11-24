package main

import (
	"benchmarking/saying"
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	expected := "Welcome Gophers from Tood!"
	result := saying.Greet("Tood")
	if result != expected {
		fmt.Printf("Test failed: expected '%s', got '%s'", expected, result)
	}
	println("Test passed")
}

func ExampleGreet() {
	fmt.Println(saying.Greet("James"))
	// Output: Welcome Gophers from James!
}

func BenchmarkGreet(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		saying.Greet("Tood")
	}
}
