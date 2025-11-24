package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	Age       int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByFirstName []Person

func (a ByFirstName) Len() int {
	return len(a)
}

func (a ByFirstName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByFirstName) Less(i, j int) bool {
	return a[i].FirstName < a[j].FirstName
}

func main() {
	p1 := Person{FirstName: "Alice", Age: 30}
	p2 := Person{FirstName: "Bob", Age: 25}
	p3 := Person{FirstName: "Charlie", Age: 35}
	p4 := Person{FirstName: "Diana", Age: 28}

	people := []Person{p1, p2, p3, p4}
	fmt.Println("Before sorting:", people)

	byAge := ByAge(people)
	sort.Sort(byAge)

	fmt.Println("After sorting by age:", people)

	byFirstName := ByFirstName(people)
	sort.Sort(byFirstName)

	fmt.Println("After sorting by first name:", people)
}
