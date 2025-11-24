package main

import (
	"fmt"
	"hands-on-2/quote"
	"hands-on-2/word"
	"testing"
)

func TestUseCount(t *testing.T) {
	m := word.UseCount("a a a b b c")
	if m["a"] != 3 {
		t.Errorf("got %d, want %d for key 'a'", m["a"], 3)
	}
	if m["b"] != 2 {
		t.Errorf("got %d, want %d for key 'b'", m["b"], 2)
	}
	if m["c"] != 1 {
		t.Errorf("got %d, want %d for key 'c'", m["c"], 1)
	}

	fmt.Println("Test passed!")
}

func TestCount(t *testing.T) {
	want := 168
	got := word.Count(quote.Quote)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
	fmt.Println("Test passed!")
}

func BenchmarkUseCount(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		word.UseCount(quote.Quote)
	}
}

func BenchmarkCount(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		word.Count(quote.Quote)
	}
}

func ExampleUseCount() {
	m := word.UseCount("hello world hello")
	fmt.Println(m["hello"])
	fmt.Println(m["world"])
	// Output:
	// 2
	// 1
}

func ExampleCount() {
	fmt.Println(word.Count("hello world hello"))
	// Output:
	// 3
}
