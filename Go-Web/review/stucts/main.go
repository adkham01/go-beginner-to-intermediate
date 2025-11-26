package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	person := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}
	fmt.Printf("Name: %s, Age: %d, City: %s\n", person.Name, person.Age, person.City)

	emoloyee := struct {
		Name string
		ID   int
	}{
		Name: "Bob",
		ID:   101,
	}
	fmt.Printf("Employee Name: %s, ID: %d\n", emoloyee.Name, emoloyee.ID)

	type Address struct {
		Street string
		City   string
	}

	type Contact struct {
		Name    string
		Phone   string
		Address Address
	}

	contact := Contact{
		Name:  "Charlie",
		Phone: "123-456-7890",
		Address: Address{
			Street: "123 Main St",
			City:   "Los Angeles",
		},
	}
	fmt.Printf("Contact Name: %s, Phone: %s, Address: %s, %s\n",
		contact.Name, contact.Phone, contact.Address.Street, contact.Address.City)

	fmt.Println("Before update:", person)
	// time.Sleep(1 * time.Second)
	person.updatePersonAge("Alex")
	fmt.Println("After update:", person)
}

func (person *Person) updatePersonAge(name string) {
	person.Name = name
}
