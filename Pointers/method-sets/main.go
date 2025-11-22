package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I am walking")
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I am running")
}

type youngin interface {
	walk()
	run()
}

func younginFunc(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	younginFunc(d1)

	d2 := dog{"Paget"}
	younginFunc(d2)
}
