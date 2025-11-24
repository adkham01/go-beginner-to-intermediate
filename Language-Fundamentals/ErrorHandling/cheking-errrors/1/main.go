package main

import "fmt"

func main() {
	var ans1, ans2, ans3 string

	fmt.Print("Name: ")
	_, err := fmt.Scan(&ans1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&ans2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&ans3)
	if err != nil {
		panic(err)
	}

	fmt.Println(ans1, ans2, ans3)

}
