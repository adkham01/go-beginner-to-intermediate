package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m[`bond_james`] = []string{"shaken", "not stirred", "martines", "fast cars"}
	m[`moneypenny_jenny`] = []string{"intelligence", "literature", "computer science"}
	m[`no_dr`] = []string{"cats", "ice cream", "sunsets"}
	m[`fleming_ian`] = []string{"steaks", "cigars", "girls"}
	delete(m, `fleming_ian`)

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
		fmt.Println("----------------------------------------")
	}

	delete(m, `fleming_ian`)
	fmt.Println("---- After the delete -----")

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
		fmt.Println("----------------------------------------")
	}

}
