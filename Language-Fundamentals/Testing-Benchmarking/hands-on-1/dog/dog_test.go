package dog

import (
	"fmt"
	"testing"
)

func TestDogYears(t *testing.T) {
	want := 21
	got := Years(3)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
	fmt.Println("Test passed!")
}

func TstDogYearsTwo(t *testing.T) {
	want := 21
	got := YearsTwo(3)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
	fmt.Println("Test passed!")
}

func ExampleYears() {
	fmt.Println(Years(4))
	// Output:
	// 28
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(4))
	// Output:
	// 28
}

func BenchmarkDogYears(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkDogYearsTwo(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}
