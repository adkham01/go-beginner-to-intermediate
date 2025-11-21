package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I am walking")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I am running")
}

func main() {
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()

	d2 := &dog{"Paget"}
	d2.walk()
	d2.run()

}
