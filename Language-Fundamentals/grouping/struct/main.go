package main

import "fmt"

type person struct {
	firstname string
	lastName  string
	age       int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	car := struct {
		model string
		mark  string
		vin   int
	}{
		model: "Chevrolet",
		mark:  "Spark",
		vin:   42345,
	}

	sa1 := secretAgent{
		person: person{
			firstname: "James",
			lastName:  "Bond",
			age:       32,
		},
		ltk: true,
	}

	fmt.Println(sa1)
	p2 := person{
		firstname: "Jane",
		lastName:  "Smith",
		age:       27,
	}

	fmt.Println(p2)
	fmt.Println(car)
	fmt.Printf("%T \t %#v\n", car, car)

	fmt.Println(p2.firstname, p2.lastName, p2.age)

}
