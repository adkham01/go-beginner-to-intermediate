package main

import "fmt"

func main() {
	foo()
	bar("Todd")
	fmt.Println(aloha("Tom"))
	name, age := dogYears("Todd ", 32)

	fmt.Println(name, age)

}

func foo() {
	fmt.Println("I am frim function foo")
}

func bar(s string) {
	fmt.Println("My name is", s, "I am from function bar")
}

func aloha(s string) string {
	return fmt.Sprint("My name is", s, "I am from function aloha")
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, "is years old in dog years"), age
}
