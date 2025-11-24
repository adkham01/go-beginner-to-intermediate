package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func main() {
	p := "Drop by drop, the bucket gets filled"
	q := "Persistently, patiently, I am bound to succeed"
	r := "Meaning of the life is"
	n := 42

	a = &p
	b = &q
	c = &r
	d = &n

	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Printf("%v\t%T\n", c, c)
	fmt.Printf("%v\t%T\n", d, d)

	fmt.Println("Value is:", *a, "address is:", a)
	fmt.Println("Value is:", *b, "address is:", b)
	fmt.Println("Value is:", *c, "address is:", c)
	fmt.Println("Value is:", *d, "address is:", d)
}
