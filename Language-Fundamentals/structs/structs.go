package main

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	p2 := Person{
		firstName: "Jane",
		lastName:  "Smith",
		age:       25,
	}

	p3 := Person{
		firstName: "Alice",
		age:       28,
	}

	println(p1.firstName, p1.lastName, p1.age)
	println(p2.firstName, p2.lastName, p2.age)
	println(p3.firstName, p3.age)

	// Anonymous struct
	anonymousPerson := struct {
		username string
		email    string
	}{
		username: "Azizbek",
		email:    "azizbek@example.com",
	}

	println(anonymousPerson.username, anonymousPerson.email)

	// Method call
	println(p1.fullName())
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}
