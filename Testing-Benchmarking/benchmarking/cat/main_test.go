package main

import (
	"benchmarking/cat/mystr"
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	s := "Shaken not stirred"

	xs := strings.Split(s, " ")
	s = mystr.Cat(xs)

	if s != "Shaken not stirred " {
		t.Errorf("Cat failed: got %q, want %q", s, "Shaken not stirred ")
	}

	fmt.Println("Cat passed")
}

func TestJoin(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = mystr.Join(xs)
	if s != "Shaken not stirred" {
		t.Errorf("Join failed: got %q, want %q", s, "Shaken not stirred")
	}
	fmt.Println("Join passed")
}

func ExampleCat() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(mystr.Cat(xs))
	// Output:
	// Shaken not stirred
}

func ExampleJoin() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(mystr.Join(xs))
	// Output:
	// Shaken not stirred
}

var xs []string

func BenchmarkCat(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mystr.Cat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mystr.Join(xs)
	}
}
