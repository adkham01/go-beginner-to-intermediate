package main

import "fmt"

func main() {
	jb := []string{"James Bond", "Shaken not stirred"}
	mm := []string{"Miss", "Moneypenny", "I'm 008"}
	fmt.Println(jb, mm)

	ss := [][]string{jb, mm}
	for i, v := range ss {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
		fmt.Println("--------------------------")
	}

}
