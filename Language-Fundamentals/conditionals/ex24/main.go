package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James": 42,
		"Paul":  32,
	}

	for k, v := range m {
		fmt.Printf("key : %v\tvalue: %v\n", k, v)
	}

	fmt.Println("----------------")
	age := m["James"]
	fmt.Println("The age of Bond ", age)

	age = m["Q"]
	fmt.Println("The age of Q ", age)

	if v, ok := m["Q"]; ok {
		fmt.Printf("Value of the q: %v and value of ok %v\n", v, ok)
	}

	if v, ok := m["Q"]; !ok {
		fmt.Printf("Value of the q: %v value of ok %v\n", v, ok)
	}

}
