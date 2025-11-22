package main

import "fmt"

type Person struct {
	first string
}

func main() {
	p := Person{
		first: "Jenny",
	}
	fmt.Println(p)

	p2 := p.changeName("Jennyfer")
	fmt.Println(p2)

	changeNamePtr(&p2, "Tom")
	fmt.Println(p2)

}

func (p Person) changeName(s string) Person {
	p.first = s
	return p
}

func changeNamePtr(p *Person, s string) {
	p.first = s
}
