package main

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	Phone
}

type Phone struct {
	number string
}

type Address struct {
	city    string
	country string
}

func main() {

	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "New York",
			country: "USA",
		},
		Phone: Phone{
			number: "123-456-7890",
		},
	}

	p2 := Person{
		firstName: "Jane",
		lastName:  "Smith",
		age:       25,
		address: Address{
			city:    "Los Angeles",
			country: "USA",
		},
		Phone: Phone{
			number: "987-654-3210",
		},
	}

	p3 := Person{
		firstName: "Alice",
		age:       28,
	}

	println(p1.firstName, p1.lastName, p1.age, "\nAddress: ", p1.address.city, p1.address.country, ", phone: ", p1.number)
	println(p2.firstName, p2.lastName, p2.age, "\nAddress: ", p2.address.city, p2.address.country, ", phone: ", p2.number)
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
	println("Before incrementing age:", p1.firstName, p1.lastName, p1.age)
	p1.incrementAge()
	println("After incrementing age:", p1.firstName, p1.lastName, p1.age)

	println("Comparing p1 and p2:", p1 == p2)
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAge() {
	p.age = p.age + 1
}
