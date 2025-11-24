package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MaxCost)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(s)
	fmt.Println(bc)

	loginPword := `password123`
	err = bcrypt.CompareHashAndPassword([]byte(loginPword), []byte(s))
	if err != nil {
		fmt.Println("You can not login")
		return
	}

	fmt.Println("You are logged in")

}
